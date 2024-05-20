package chain

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/utils"
	cfg "github.com/cometbft/cometbft/config"
	nm "github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/proxy"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type Node struct {
	node *nm.Node
}

func New(logger *utils.Logger, homeDir string, pool *pgxpool.Pool) (*Node, error) {
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

	db := db.New(pool)

	app := NewApplication(logger, db, pool)

	pv := privval.LoadFilePV(
		config.PrivValidatorKeyFile(),
		config.PrivValidatorStateFile(),
	)

	nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, utils.AppError("Error loading p2p key", err)
	}

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
	}, nil
}

func (n *Node) Node() *nm.Node {
	return n.node
}

func (n *Node) Run() {
	node := n.node

	node.Start()

	defer func() {
		node.Stop()
		node.Wait()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
