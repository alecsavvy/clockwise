/* the client that external modules use like grpc and clis */
package core

import (
	"fmt"

	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/proxy"
	"github.com/cometbft/cometbft/rpc/client/local"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type Core struct {
	logger *utils.Logger
	rpc    *local.Local
	pool   *pgxpool.Pool
	pubsub *Pubsub
}

var _ abcitypes.Application = (*Core)(nil)

func NewCore(logger *utils.Logger, pool *pgxpool.Pool) *Core {
	return &Core{
		logger: logger,
		pool:   pool,
		pubsub: NewPubsub(),
	}
}

func NewNode(logger *utils.Logger, homeDir string, app abcitypes.Application) (*node.Node, error) {
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

	pv := privval.LoadFilePV(
		config.PrivValidatorKeyFile(),
		config.PrivValidatorStateFile(),
	)

	nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, utils.AppError("Error loading p2p key", err)
	}

	node, err := node.NewNode(
		config,
		pv,
		nodeKey,
		proxy.NewLocalClientCreator(app),
		node.DefaultGenesisDocProviderFunc(config),
		cfg.DefaultDBProvider,
		node.DefaultMetricsProvider(config.Instrumentation),
		logger,
	)

	if err != nil {
		return nil, utils.AppError("Error initializing node", err)
	}

	return node, nil
}

func (c *Core) Pubsub() *Pubsub {
	return c.pubsub
}

func (c *Core) Rpc() *local.Local {
	return c.rpc
}

func (c *Core) Run(node *node.Node) error {
	c.rpc = local.New(node)
	err := c.RunPubsub()
	if err != nil {
		c.logger.Error("pubsub error", err)
	}
	return nil
}
