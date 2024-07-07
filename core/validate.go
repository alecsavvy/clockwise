package core

import (
	"errors"

	"github.com/alecsavvy/clockwise/protocol"
	"github.com/alecsavvy/clockwise/protocol/gen"
	"google.golang.org/protobuf/proto"
)

func (c *Core) validateTx(msg []byte) error {
	c.logger.Info("validating tx", "routes", len(c.validationRoutes))
	return protocol.MessageRouter(c.validationRoutes, msg)
}

func (c *Core) getSender(*gen.Envelope) (string, error) {
	// get sender via signature
	return "", nil
}

func (c *Core) validateSignature() {}

func (c *Core) validateCreateUser(msg proto.Message) error {

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

func (c *Core) validateCreateTrack(msg proto.Message) error {
	msg, ok := msg.(*gen.CreateTrack)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}

func (c *Core) validateRepostTrack(msg proto.Message) error {
	msg, ok := msg.(*gen.RepostTrack)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}

func (c *Core) validateUnRepostTrack(msg proto.Message) error {
	msg, ok := msg.(*gen.UnrepostTrack)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}

func (c *Core) validateFollowUser(msg proto.Message) error {
	msg, ok := msg.(*gen.FollowUser)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}

func (c *Core) validateUnfollowUser(msg proto.Message) error {
	msg, ok := msg.(*gen.UnfollowUser)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}
	return nil
}
