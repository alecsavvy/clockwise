package core

import (
	"context"
	"fmt"

	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/protocol"
	"github.com/alecsavvy/clockwise/protocol/gen"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/proxy"
	"github.com/cometbft/cometbft/rpc/client/local"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type Core struct {
	logger *utils.Logger
	rpc    *local.Local
	node   *node.Node
	pubsub *Pubsub

	validationRoutes protocol.MessageRouterMap
	indexingRoutes   protocol.MessageRouterMap

	queries *db.Queries

	pool      *pgxpool.Pool
	currentTx pgx.Tx

	// app config
	RetainBlocks int64
}

var _ abcitypes.Application = (*Core)(nil)

func NewCore(logger *utils.Logger, pool *pgxpool.Pool, retainBlocks int64) *Core {
	core := &Core{
		logger:       logger,
		pubsub:       NewPubsub(),
		queries:      db.New(pool),
		pool:         pool,
		RetainBlocks: retainBlocks,
	}

	// register validation routes
	validateRoutes := make(protocol.MessageRouterMap, 0)
	validateRoutes[gen.MessageType_MESSAGE_TYPE_CREATE_USER] = core.validateCreateUser
	validateRoutes[gen.MessageType_MESSAGE_TYPE_CREATE_TRACK] = core.validateCreateTrack
	validateRoutes[gen.MessageType_MESSAGE_TYPE_REPOST_TRACK] = core.validateRepostTrack
	validateRoutes[gen.MessageType_MESSAGE_TYPE_UNREPOST_TRACK] = core.validateUnRepostTrack
	validateRoutes[gen.MessageType_MESSAGE_TYPE_FOLLOW_USER] = core.validateFollowUser
	validateRoutes[gen.MessageType_MESSAGE_TYPE_UNFOLLOW_USER] = core.validateUnfollowUser

	// register indexing routes
	indexRoutes := make(protocol.MessageRouterMap, 0)
	indexRoutes[gen.MessageType_MESSAGE_TYPE_CREATE_USER] = core.indexCreateUser
	indexRoutes[gen.MessageType_MESSAGE_TYPE_CREATE_TRACK] = core.indexCreateTrack
	indexRoutes[gen.MessageType_MESSAGE_TYPE_REPOST_TRACK] = core.indexRepostTrack
	indexRoutes[gen.MessageType_MESSAGE_TYPE_UNREPOST_TRACK] = core.indexUnrepostTrack
	indexRoutes[gen.MessageType_MESSAGE_TYPE_FOLLOW_USER] = core.indexFollowUser
	indexRoutes[gen.MessageType_MESSAGE_TYPE_UNFOLLOW_USER] = core.indexUnfollowUser

	core.validationRoutes = validateRoutes
	core.indexingRoutes = indexRoutes

	return core
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
		context.Background(),
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
	c.node = node
	err := c.RunPubsub()
	if err != nil {
		c.logger.Error("pubsub error", err)
	}
	return nil
}
