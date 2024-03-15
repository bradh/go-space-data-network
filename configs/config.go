package config

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"golang.org/x/crypto/argon2"
)

var once sync.Once

type datastoreConfig struct {
	Directory string
	Password  string
}

type webserverConfig struct {
	Port int
}

type keyConfig struct {
	EntropyLengthBits int
}

type keys struct {
	EncryptionAccountDerivationPath string
	SigningAccountDerivationPath    string
}

// AppConfig holds the entire application configuration with namespaces
type AppConfig struct {
	Datastore datastoreConfig
	Webserver webserverConfig
	KeyConfig keyConfig
	Keys      keys
}

// Conf is the exported variable that will hold all the configuration settings
var Conf AppConfig

// Init initializes the global configuration
func Init() {
	once.Do(func() {
		// Webserver settings
		var webserverPortStr string
		flag.StringVar(&webserverPortStr, "webserver.port", "4000", "Port for the webserver to listen on")

		// Datastore settings
		flag.StringVar(&Conf.Datastore.Directory, "datastore.directory", "", "Directory for the datastore")
		flag.StringVar(&Conf.Datastore.Password, "datastore.password", "", "Password for the datastore encryption")

		// Parse command-line flags
		flag.Parse()

		// Override webserver port with environment variable if exists
		if portStr, exists := os.LookupEnv("SPACE_DATA_NETWORK_WEBSERVER_PORT"); exists {
			webserverPortStr = portStr
		}
		if port, err := strconv.Atoi(webserverPortStr); err == nil {
			Conf.Webserver.Port = port
		} else {
			Conf.Webserver.Port = 1969 // Default port if conversion fails
		}

		// Override datastore settings with environment variables if they exist
		if dir, exists := os.LookupEnv("SPACE_DATA_NETWORK_DATASTORE_DIRECTORY"); exists {
			Conf.Datastore.Directory = dir
		} else {
			// Use the user's home directory if no environment variable or custom path is provided
			homeDir, err := os.UserHomeDir()
			if err != nil {
				log.Printf("Error getting user home directory: %v", err)
				return // Early return on error
			}

			// Set the default directory to a specific folder in the user's home directory
			defaultDir := filepath.Join(homeDir, ".spacedatanetwork")
			Conf.Datastore.Directory = defaultDir

			// Ensure the base directory exists
			if _, err := os.Stat(Conf.Datastore.Directory); os.IsNotExist(err) {
				if err := os.MkdirAll(Conf.Datastore.Directory, 0700); err != nil {
					log.Printf("Error creating directory '%s': %v", Conf.Datastore.Directory, err)
					return // Early return on error
				}
			}
		}

		if password, exists := os.LookupEnv("SPACE_DATA_NETWORK_DATASTORE_PASSWORD"); exists {
			Conf.Datastore.Password = password
		} else {

			hostname, _ := os.Hostname()
			homeDir, _ := os.UserHomeDir()

			input := fmt.Sprintf("%s:%s", homeDir, hostname)
			Conf.Datastore.Password = hex.EncodeToString(argon2.IDKey([]byte(input), []byte("some_salt"), 1, 64*1024, 4, 32))

		}

		Conf.Keys.SigningAccountDerivationPath = "m/44'/60'/0'/0/0'"
		Conf.Keys.EncryptionAccountDerivationPath = "m/44'/60'/0'/1/0'"

		Conf.KeyConfig.EntropyLengthBits = 256

	})
}
