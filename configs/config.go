package config

import (
	"flag"
	"os"
	"strconv"
)

type datastoreConfig struct {
	Directory string
	Password  string
}

type webserverConfig struct {
	Port int
}

// AppConfig holds the entire application configuration with namespaces
type AppConfig struct {
	Datastore datastoreConfig
	Webserver webserverConfig
}

// Conf is the exported variable that will hold all the configuration settings
var Conf AppConfig

// Init initializes the global configuration
func Init() {
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
	}
	if password, exists := os.LookupEnv("SPACE_DATA_NETWORK_DATASTORE_PASSWORD"); exists {
		Conf.Datastore.Password = password
	}
}