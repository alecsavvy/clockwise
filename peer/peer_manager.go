package peer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/alecsavvy/clockwise/common"
)

type PeerSet map[string]Peer

type PeerManager struct {
	self   string
	nodes  PeerSet
	config *common.Config
}

func New(config *common.Config) (*PeerManager, error) {
	return &PeerManager{
		self:   config.NodeEndpoint,
		nodes:  make(PeerSet),
		config: config,
	}, nil
}

func (pm *PeerManager) ConnectPeers() error {
	fmt.Println("connecting to peers")
	for {
		bootstrapNodes := pm.config.NodeBootstrapPeers

		var wg sync.WaitGroup
		var mutex sync.Mutex

		for _, endpoint := range bootstrapNodes {
			wg.Add(1)
			go func(endpoint string) {
				defer wg.Done()
				mutex.Lock()
				_, exists := pm.nodes[endpoint]
				mutex.Unlock()
				if exists {
					return
				}
				fmt.Printf("connecting to %s\n", endpoint)
				peer, err := NewPeer(endpoint)
				if err != nil {
					fmt.Printf("could not peer with node %s due to: %s\n", endpoint, err)
					return
				}
				mutex.Lock()
				pm.nodes[endpoint] = *peer
				mutex.Unlock()
				fmt.Printf("connected to peer %s\n", endpoint)
			}(endpoint)
		}
		wg.Wait()
		time.Sleep(time.Duration(time.Second * 5))
	}
}

func (pm *PeerManager) PollPeerHealth() error {
	fmt.Println("bootstrapping nodes")
	for {
		for endpoint, node := range pm.nodes {
			_, err := node.rpc.GetNodeHealth(context.Background(), nil)
			if err != nil {
				fmt.Printf("could not get health from %s\n", endpoint)
				continue
			}
			fmt.Printf("received health from node %s\n", endpoint)
		}
		time.Sleep(time.Duration(time.Second * 3))
	}
}
