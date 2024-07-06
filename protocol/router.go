package protocol

import (
	"errors"
	"fmt"

	"github.com/alecsavvy/clockwise/protocol/gen"
	"google.golang.org/protobuf/proto"
)

type MessageRouterFunc func(proto.Message) error
type MessageRouterMap map[gen.MessageType]MessageRouterFunc

func MessageRouter(router MessageRouterMap, message []byte) error {
	envelope, err := deserializeEnvelope(message)
	if err != nil {
		return err
	}

	route, ok := router[envelope.MessageType]
	if !ok {
		return errors.New(fmt.Sprintf("route for message %s not registered", envelope.MessageType))
	}

	switch envelope.MessageType {
	case gen.MessageType_MESSAGE_TYPE_CREATE_USER:
		msg, err := deserializeCreateUser(message)
		if err != nil {
			return err
		}
		return route(msg)
	case gen.MessageType_MESSAGE_TYPE_FOLLOW_USER:
		msg, err := deserializeFollowUser(message)
		if err != nil {
			return err
		}
		return route(msg)
	case gen.MessageType_MESSAGE_TYPE_UNFOLLOW_USER:
		msg, err := deserializeUnfollowUser(message)
		if err != nil {
			return err
		}
		return route(msg)
	case gen.MessageType_MESSAGE_TYPE_CREATE_TRACK:
		msg, err := deserializeCreateTrack(message)
		if err != nil {
			return err
		}
		return route(msg)
	case gen.MessageType_MESSAGE_TYPE_REPOST_TRACK:
		msg, err := deserializeRepostTrack(message)
		if err != nil {
			return err
		}
		return route(msg)
	case gen.MessageType_MESSAGE_TYPE_UNREPOST_TRACK:
		msg, err := deserializeUnrepostTrack(message)
		if err != nil {
			return err
		}
		return route(msg)
	default:
		return errors.New(fmt.Sprintf("route for message %s not handled", envelope.MessageType))
	}
}
