package node

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"runtime/debug"
	"time"

	configs "github.com/DigitalArsenal/space-data-network/configs"
	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ipfs/kubo/core"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/p2p/host/autorelay"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	tls "github.com/libp2p/go-libp2p/p2p/security/tls"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/libp2p/go-libp2p/p2p/transport/websocket"
	webtransport "github.com/libp2p/go-libp2p/p2p/transport/webtransport"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"golang.org/x/crypto/argon2"

	dht "github.com/libp2p/go-libp2p-kad-dht"
)

type Node struct {
	Host              host.Host
	DHT               *dht.IpfsDHT
	KeyStore          *KeyStore
	wallet            *hdwallet.Wallet
	signingAccount    accounts.Account
	encryptionAccount accounts.Account
	IPFS              *core.IpfsNode
	peerChan          chan peer.AddrInfo
}

// autoRelayPeerSource returns a function that provides peers for auto-relay.
func autoRelayFeeder(ctx context.Context, h host.Host, dht *dht.IpfsDHT, peerChan chan<- peer.AddrInfo) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from unexpected error in AutoRelayFeeder:", r)
			debug.PrintStack()
		}
	}()
	go func() {
		// Feed peers more often right after the bootstrap, then backoff
		bo := backoff.NewExponentialBackOff()
		bo.InitialInterval = 15 * time.Second
		bo.Multiplier = 3
		bo.MaxInterval = 1 * time.Hour
		bo.MaxElapsedTime = 0 // never stop
		t := backoff.NewTicker(bo)
		defer t.Stop()
		for {
			select {
			case <-t.C:
			case <-ctx.Done():
				return
			}

			closestPeers, err := dht.GetClosestPeers(ctx, h.ID().String())
			if err != nil {
				// no-op: usually 'failed to find any peer in table' during startup
				continue
			}
			for _, p := range closestPeers {
				addrs := h.Peerstore().Addrs(p)
				if len(addrs) == 0 {
					continue
				}
				dhtPeer := peer.AddrInfo{ID: p, Addrs: addrs}
				select {
				case peerChan <- dhtPeer:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
}

func NewNode(ctx context.Context) (*Node, error) {
	configs.Init()

	node := &Node{}

	var err error

	repo, privKey, encPrivKey, err := LoadOrCreateIPFSRepo(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to load or create IPFS repo: %w", err)
	}

	node.Host, err = libp2p.New(
		libp2p.Identity(privKey),
		libp2p.EnableNATService(),
		libp2p.EnableRelay(),
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0", "/ip6/::/tcp/0"),
		libp2p.Security(noise.ID, noise.New),
		libp2p.Security(tls.ID, tls.New),
		libp2p.Transport(tcp.NewTCPTransport, tcp.WithMetrics()),
		libp2p.Transport(websocket.New),
		libp2p.Transport(quic.NewTransport),
		libp2p.Transport(webtransport.New),
		libp2p.NATPortMap(),
		libp2p.EnableRelay(),
		libp2p.EnableHolePunching(),
		libp2p.EnableAutoRelayWithPeerSource(
			func(ctx context.Context, _ int) <-chan peer.AddrInfo {
				return node.peerChan
			},
			autorelay.WithMinInterval(0)),
	)

	if err != nil {
		return node, fmt.Errorf("failed to create libp2p host: %w", err)
	}

	peerID := node.Host.ID()

	// Update the repository configuration with the peer ID
	cfg, err := repo.Config()
	if err != nil {
		return nil, fmt.Errorf("failed to read IPFS config: %w", err)
	}

	if cfg.Identity.PeerID != peerID.String() {
		cfg.Identity.PeerID = peerID.String()
		cfg.Identity.PrivKey = encPrivKey
		if err := repo.SetConfig(cfg); err != nil {
			return nil, fmt.Errorf("failed to update IPFS repo config: %w", err)
		}
	}

	cPK, _ := crypto.MarshalPrivateKey(privKey)
	cfg.Identity.PrivKey = base64.StdEncoding.EncodeToString(cPK)

	customHostOption := func(id peer.ID, ps peerstore.Peerstore, options ...libp2p.Option) (host.Host, error) {
		return node.Host, nil
	}

	node.IPFS, err = core.NewNode(ctx, &core.BuildCfg{
		Online:    true,
		Permanent: true,
		Host:      customHostOption,
		Repo:      repo,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create IPFS node: %w", err)
	}

	//node.wallet = wallet
	//node.signingAccount = signingAccount
	//node.encryptionAccount = encryptionAccount
	repoConfig, err := node.IPFS.Repo.Config()
	if err != nil {
		fmt.Printf("Failed to get IPFS repo config: %s\n", err)
	} else {
		// The Repo's path is typically not directly exposed in the config,
		// but the Repo itself knows its path which can be accessed this way
		repoPath := repoConfig.Datastore.Path
		fmt.Println("IPFS repo path:", repoPath)
	}
	fmt.Println("")
	fmt.Println("Node PeerID: ", node.Host.ID())
	fmt.Println("Node PeerID: ", node.IPFS.PeerHost.ID())
	fmt.Println("Node Signing Ethereum Address: ", node.signingAccount.Address)
	fmt.Println("Node Encryption Ethereum Address: ", node.encryptionAccount.Address)
	fmt.Println("")

	CreateDefaultServerEPM(node)

	return node, nil
}

func (n *Node) Start(ctx context.Context) error {
	var err error
	n.peerChan = make(chan peer.AddrInfo, 100) // Buffer to avoid blocking

	//vepm, _ := n.KeyStore.LoadEPM()

	//newCID, _ := n.AddFileFromBytes(ctx, vepm)

	//fmt.Println("ADDED CID FOR EPM: ")
	//fmt.Println(newCID)

	SetupPNMExchange(n)

	n.DHT, err = initDHT(ctx, n.Host)
	if err != nil {
		return fmt.Errorf("failed to initialize DHT: %w", err)
	}

	//Start auto relay
	autoRelayFeeder(ctx, n.Host, n.DHT, n.peerChan)
	go discoverPeers(ctx, n, "space-data-network", 30*time.Second)
	//Find others with the same version
	versionHex := []byte(configs.Conf.Info.Version)
	discoveryHex := hex.EncodeToString(argon2.IDKey(versionHex, versionHex, 1, 64*1024, 4, 32))
	go discoverPeers(ctx, n, discoveryHex, 30*time.Second)

	//_, err = n.PublishIPNSRecord(ctx, newCID.String())
	//if err != nil {
	//	return fmt.Errorf("failed to publish CID to IPNS: %w", err)
	//}
	//fmt.Println("NEW IPNS CID:", ipnsCID)

	return nil
}

func (n *Node) Stop() {
	fmt.Println("Shutting down node...")

	if n.KeyStore != nil {
		if err := n.KeyStore.Close(); err != nil {
			fmt.Println("Failed to close Keystore:", err)
		}
	}
	if n.Host != nil {
		if err := n.Host.Close(); err != nil {
			fmt.Println("Failed to close libp2p host:", err)
		}
	}
	if n.DHT != nil {
		if err := n.DHT.Close(); err != nil {
			fmt.Println("Failed to close DHT:", err)
		}
	}
	fmt.Println("Node stopped successfully.")
}
