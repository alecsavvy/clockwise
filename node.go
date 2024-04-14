package main

const (
	Observer = iota
	Sealer
	Archiver
)

// node instance config, network wide config goes in network.go e.g. pruning depth (how many blocks to keep in storage)
type Node struct {
	NodeConfig
	DatabaseConfig
}

// config about this nodes runtime
type NodeConfig struct {
}

// config about this nodes storage mechanism
type DatabaseConfig struct {
}
