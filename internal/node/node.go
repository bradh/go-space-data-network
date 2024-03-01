// node/node.go
package node

import (
	"bufio"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto"

	"os"
	"sync"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/libp2p/go-libp2p/p2p/host/autorelay"
)

type Node struct {
	Host host.Host
	DHT  *dht.IpfsDHT
}

// NewNode initializes a new libp2p node with the given configuration.
func NewNode(ctx context.Context) (*Node, error) {
	keyStore, err := NewKeyStore()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize key store: %w", err)
	}
	pass := os.Getenv("KEYSTORE_PASSWORD")
	if pass == "" {
		pass = generatePassword() // From keystore.go
	}
	privKey, err := keyStore.GetPrivateKey(pass)
	if err != nil {
		return nil, fmt.Errorf("failed to get private key: %w", err)
	}

	h, err := libp2p.New(
		libp2p.Identity(privKey),
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
		libp2p.EnableNATService(),
		libp2p.EnableRelay(),
		libp2p.EnableHolePunching(),
		libp2p.EnableAutoRelayWithPeerSource(
			autoRelayPeerSource,
			autorelay.WithMinInterval(0)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create libp2p host: %w", err)
	}

	d, err := initDHT(ctx, h)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize DHT: %w", err)
	}

	go discoverPeers(ctx, h, d)

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}

	topic, err := ps.Join("space-data-network")
	if err != nil {
		panic(err)
	}

	go streamConsoleTo(ctx, topic)

	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}

	printMessagesFrom(ctx, sub)

	return &Node{Host: h, DHT: d}, nil
}

// autoRelayPeerSource returns a function that provides peers for auto-relay.
func autoRelayPeerSource(ctx context.Context, numPeers int) <-chan peer.AddrInfo {

	peerChan := make(chan peer.AddrInfo)

	r := make(chan peer.AddrInfo)

	go func() {
		defer close(r)
		for ; numPeers != 0; numPeers-- {
			select {
			case v, ok := <-peerChan:
				if !ok {
					return
				}
				select {
				case r <- v:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return r

}

func initDHT(ctx context.Context, h host.Host) (*dht.IpfsDHT, error) {
	kademliaDHT, err := dht.New(ctx, h)
	if err != nil {
		return nil, err
	}
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	for _, peerAddr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Connect(ctx, *peerinfo); err != nil {
				fmt.Println("Bootstrap warning:", err)
			}
		}()
	}
	wg.Wait()

	return kademliaDHT, nil
}

func discoverPeers(ctx context.Context, h host.Host, d *dht.IpfsDHT) {
	routingDiscovery := drouting.NewRoutingDiscovery(d)
	dutil.Advertise(ctx, routingDiscovery, "space-data-network")

	anyConnected := false
	for !anyConnected {
		fmt.Println("Searching for peers...")
		peerChan, err := routingDiscovery.FindPeers(ctx, "space-data-network")
		if err != nil {
			panic(err)
		}
		for peer := range peerChan {
			if peer.ID == h.ID() {
				continue
			}
			err := h.Connect(ctx, peer)
			if err != nil {
				// Commented out the error printing for connection failure
				// fmt.Printf("Failed connecting to %s, error: %s\n", peer.ID, err)
			} else {
				fmt.Printf("Connected to: %s\n", peer.ID)
				for _, addr := range peer.Addrs {
					fmt.Printf("\t%s/p2p/%s\n", addr, peer.ID)
				}
				anyConnected = true
			}
		}
	}
	fmt.Println("Peer discovery complete")
}

// Start begins the node operation, such as connecting to peers and handling connections.
func (n *Node) Start(ctx context.Context) {

}

// Stop gracefully shuts down the libp2p node.
func (n *Node) Stop() {
	if err := n.Host.Close(); err != nil {
		fmt.Println("Failed to stop node:", err)
	} else {
		fmt.Println("Node stopped successfully.")
	}
}

func streamConsoleTo(ctx context.Context, topic *pubsub.Topic) {
	reader := bufio.NewReader(os.Stdin)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if err := topic.Publish(ctx, []byte(s)); err != nil {
			fmt.Println("### Publish error:", err)
		}
	}
}

func printMessagesFrom(ctx context.Context, sub *pubsub.Subscription) {
	for {
		m, err := sub.Next(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(m.ReceivedFrom, ": ", string(m.Message.Data))
	}
}

// Add a new method to Node to extract the public key
func (n *Node) PublicKey() (string, error) {
	pubKey, err := n.Host.ID().ExtractPublicKey()
	if err != nil {
		return "", fmt.Errorf("failed to extract public key: %w", err)
	}
	if pubKey != nil {
		return "", fmt.Errorf("public key is nil")
	}
	pubKeyBytes, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return "", fmt.Errorf("failed to marshal public key: %w", err)
	}

	return hex.EncodeToString(pubKeyBytes), nil

}
