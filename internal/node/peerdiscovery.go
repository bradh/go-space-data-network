package node

import (
	"context"
	"fmt"
	"sync"
	"time"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
)

func discoverPeers(ctx context.Context, n *Node, channelName string, discoveryInterval time.Duration) {

	h := n.Host
	d := n.DHT

	contactedPeers := make(map[peer.ID]struct{})
	connectedPeers := make(map[peer.ID]struct{})
	mutex := sync.Mutex{}

	// Create a NotifyBundle and assign event handlers
	notifiee := &NotifyBundle{
		ConnectedF: func(_ network.Network, conn network.Conn) {
			//TODO connect to any peer
		},
		DisconnectedF: func(_ network.Network, conn network.Conn) {
			mutex.Lock()
			defer mutex.Unlock()
			peerID := conn.RemotePeer()
			_, exists := connectedPeers[peerID]
			if exists {
				delete(connectedPeers, peerID)
			}
		},
	}

	// Register notifiee with the host's network
	h.Network().Notify(notifiee)
	defer h.Network().StopNotify(notifiee)

	ticker := time.NewTicker(discoveryInterval)
	defer ticker.Stop()

	routingDiscovery := drouting.NewRoutingDiscovery(d)
	dutil.Advertise(ctx, routingDiscovery, channelName)
	discoveredPeersChan := make(chan peer.AddrInfo)

	// Initialize mDNS service
	notifee := &discoveryNotifee{h: h, contactedPeers: make(map[peer.ID]struct{}), mutex: &sync.Mutex{}, discoveredPeersChan: discoveredPeersChan}
	mdnsService := mdns.NewMdnsService(h, "space-data-network-mdns", notifee)
	go func() {
		if err := mdnsService.Start(); err != nil {
			fmt.Println("Failed to start mDNS service:", err)
		}
	}()
	defer mdnsService.Close()

	ticker = time.NewTicker(discoveryInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping peer discovery due to context cancellation")
			return
		case <-ticker.C:
			fmt.Println("Searching for peers...")
			peerChan, err := routingDiscovery.FindPeers(ctx, channelName)
			if err != nil {
				panic(err)
			}
			for peer := range peerChan {
				if peer.ID == h.ID() {
					continue
				}

				if alreadyContacted(peer.ID, contactedPeers, &mutex) {
					continue
				}
				markAsContacted(peer.ID, contactedPeers, &mutex)

				err := h.Connect(ctx, peer)
				if err != nil {
					continue
				}

				for _, addr := range peer.Addrs {
					fmt.Printf("\t%s/p2p/%s\n", addr, peer.ID)
				}
				fmt.Printf("Connected to: %s\n", peer.ID)
				// Request PNM from the connected DHT peer
				if err := RequestPNM(ctx, h, peer.ID); err != nil {
					fmt.Printf("Failed to request PNM from %s: %v\n", peer.ID, err)
					continue
				}

				// Add connected peer to connectedPeers map
				connectedPeers[peer.ID] = struct{}{}
			}

		case pi := <-discoveredPeersChan: // Handle peers discovered via mDNS
			fmt.Printf("mDNS discovered peer: %s\n", pi.ID)

			if alreadyContacted(pi.ID, contactedPeers, &mutex) {
				continue
			}
			markAsContacted(pi.ID, contactedPeers, &mutex)

			if err := h.Connect(ctx, pi); err != nil {
				continue
			}
			fmt.Printf("Connected to mDNS peer: %s\n", pi.ID)

			// Request PNM from the connected mDNS peer
			if err := RequestPNM(ctx, h, pi.ID); err != nil {
				fmt.Printf("Failed to request PNM from %s: %v\n", pi.ID, err)
			}

			// Add connected mDNS peer to connectedPeers map
			connectedPeers[pi.ID] = struct{}{}
		}
	}
}

func alreadyContacted(peerID peer.ID, contactedPeers map[peer.ID]struct{}, mutex *sync.Mutex) bool {
	mutex.Lock()
	defer mutex.Unlock()
	_, contacted := contactedPeers[peerID]
	return contacted
}

func markAsContacted(peerID peer.ID, contactedPeers map[peer.ID]struct{}, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	contactedPeers[peerID] = struct{}{}
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
