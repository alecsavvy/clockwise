package main

// things regarding the network e.g. network id, block sealing time, etc
type Network struct {
	NetworkConfig
	BlockConfig
	SealerConfig
}

// network wide network config
type NetworkConfig struct {
	NetworkID string
}

// network wide block config
type BlockConfig struct {
	// amount of transactions that can be housed in one block
	TransactionCeiling uint32
	// amount of time a signer will wait until signing new block should transaction ceiling not be hit
	BlockMineMaximumMS uint32
}

type SealerConfig struct {
}
