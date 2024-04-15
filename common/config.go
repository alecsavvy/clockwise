package common

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

type NodeType string

const (
	Observer NodeType = "observer"
	Sealer   NodeType = "sealer"
	Archiver NodeType = "archiver"
)

type Config struct {
	/** Network Config */
	NetworkID                 string
	NetworkTransactionCeiling uint32
	NetworkBlockMineMaximumMS uint32

	/** Node Config */
	NodeEndpoint       string
	NodeBootstrapPeers []string
	NodeType           NodeType

	/** Database Config */
	DatabaseBlockPersistenceThreshold uint32
	DatabasePruningIntervalMS         uint32
}

func ReadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	if !IsValidNodeType(config.NodeType) {
		return nil, errors.New(fmt.Sprintf("invalid node type %s", config.NodeType))
	}
	return &config, nil
}

func IsValidNodeType(nodeType NodeType) bool {
	switch NodeType(nodeType) {
	case Observer, Sealer, Archiver:
		return true
	}
	return false
}
