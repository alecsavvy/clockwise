package common

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	/** Network Config */
	NetworkID                 string
	NetworkTransactionCeiling uint32
	NetworkBlockMineMaximumMS uint32

	/** Node Config */
	NodeEndpoint       string
	NodeBootstrapPeers []string
	NodeType           string

	/** Database Config */
	DatabaseBlockPersistenceThreshold uint32
	DatabasePruningIntervalMS         uint32
}

func ReadConfig() (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile("path/to/your/config.toml", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
