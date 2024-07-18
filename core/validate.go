package core

import (
	"context"
	"errors"

	"github.com/alecsavvy/clockwise/protocol"
	"github.com/alecsavvy/clockwise/protocol/gen"
	"google.golang.org/protobuf/proto"
)

func (c *Core) validateTx(ctx context.Context, msg []byte) error {
	c.logger.Info("validating tx", "routes", len(c.validationRoutes))
	return protocol.MessageRouter(ctx, c.validationRoutes, msg)
}

func (c *Core) validateCreateUser(ctx context.Context, msg proto.Message) error {
	message, ok := msg.(*gen.CreateUser)
	c.logger.Info("validating create user", "message", message)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}

	if len(message.GetData().GetHandle()) > 8 {
		return errors.New("handle is too long")
	}

	return nil
}

func (c *Core) validateCreateTrack(ctx context.Context, msg proto.Message) error {
	msg, ok := msg.(*gen.CreateTrack)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}

func (c *Core) validateRepostTrack(ctx context.Context, msg proto.Message) error {
	msg, ok := msg.(*gen.RepostTrack)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}

func (c *Core) validateUnRepostTrack(ctx context.Context, msg proto.Message) error {
	msg, ok := msg.(*gen.UnrepostTrack)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}

func (c *Core) validateFollowUser(ctx context.Context, msg proto.Message) error {
	msg, ok := msg.(*gen.FollowUser)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}

func (c *Core) validateUnfollowUser(ctx context.Context, msg proto.Message) error {
	msg, ok := msg.(*gen.UnfollowUser)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}
