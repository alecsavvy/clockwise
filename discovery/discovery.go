package discovery

import "fmt"

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
	fmt.Println("discovery nodes...")
	return nil
}
