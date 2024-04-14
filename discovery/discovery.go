package discovery

import (
	"fmt"
	"time"
)

type NodeSet map[string]bool

type Discovery struct {
	self  string
	nodes NodeSet
}

func New(endpoint string) (*Discovery, error) {
	return &Discovery{
		self:  endpoint,
		nodes: make(NodeSet),
	}, nil
}

func (discovery *Discovery) DiscoverNodes(bootstrap []string) error {
	fmt.Println("bootstrapping nodes")
	for true {
		time.Sleep(time.Duration(time.Second * 3))
		fmt.Println("discovered new nodes")
	}
	return nil
}
