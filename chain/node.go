package chain

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/alecsavvy/clockwise/utils"
	cfg "github.com/cometbft/cometbft/config"
	nm "github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/proxy"
	"github.com/dgraph-io/badger/v4"
	"github.com/spf13/viper"
)

type Node struct {
	node *nm.Node
	abci *KVStoreApplication
}

func New(logger *utils.Logger, homeDir string) (*Node, error) {
	config := cfg.DefaultConfig()
	config.SetRoot(homeDir)
	viper.SetConfigFile(fmt.Sprintf("%s/%s", homeDir, "config/config.toml"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, utils.AppError("Reading config", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, utils.AppError("Decoding config", err)
	}
	if err := config.ValidateBasic(); err != nil {
		return nil, utils.AppError("Chain config validation", err)
	}

	dbPath := filepath.Join(homeDir, "badger")
	dbOpts := badger.DefaultOptions(dbPath)
	dbOpts.Logger = logger
	db, err := badger.Open(dbOpts)

	if err != nil {
		return nil, utils.AppError("Opening database", err)
	}

	app := NewKVStoreApplication(logger, db)

	pv := privval.LoadFilePV(
		config.PrivValidatorKeyFile(),
		config.PrivValidatorStateFile(),
	)

	nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, utils.AppError("Error loading p2p key", err)
	}

	// logger := cmtlog.NewTMLogger(cmtlog.NewSyncWriter(os.Stdout))

	node, err := nm.NewNode(
		config,
		pv,
		nodeKey,
		proxy.NewLocalClientCreator(app),
		nm.DefaultGenesisDocProviderFunc(config),
		cfg.DefaultDBProvider,
		nm.DefaultMetricsProvider(config.Instrumentation),
		logger,
	)

	if err != nil {
		return nil, utils.AppError("Error initializing node", err)
	}

	return &Node{
		node: node,
		abci: app,
	}, nil
}

func (n *Node) RPC() string {
	return n.node.Config().RPC.ListenAddress
}

func (n *Node) Run() {
	node := n.node
	abci := n.abci

	node.Start()

	defer func() {
		if err := abci.db.Close(); err != nil {
			log.Printf("Closing database: %v", err)
		}
		node.Stop()
		node.Wait()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
