package protocol

import (
	"github.com/alecsavvy/clockwise/protocol/gen"
	"google.golang.org/protobuf/proto"
)

func deserializeEnvelope(data []byte) (*gen.Headers, error) {
	envelope := &gen.Envelope{}
	err := proto.Unmarshal(data, envelope)
	return envelope.GetHeaders(), err
}

func deserializeCreateUser(data []byte) (*gen.CreateUser, error) {
	createUser := &gen.CreateUser{}
	err := proto.Unmarshal(data, createUser)
	return createUser, err
}

func deserializeFollowUser(data []byte) (*gen.FollowUser, error) {
	followUser := &gen.FollowUser{}
	err := proto.Unmarshal(data, followUser)
	return followUser, err
}

func deserializeUnfollowUser(data []byte) (*gen.UnfollowUser, error) {
	unfollowUser := &gen.UnfollowUser{}
	err := proto.Unmarshal(data, unfollowUser)
	return unfollowUser, err
}

func deserializeCreateTrack(data []byte) (*gen.CreateTrack, error) {
	createTrack := &gen.CreateTrack{}
	err := proto.Unmarshal(data, createTrack)
	return createTrack, err
}

func deserializeRepostTrack(data []byte) (*gen.RepostTrack, error) {
	repostTrack := &gen.RepostTrack{}
	err := proto.Unmarshal(data, repostTrack)
	return repostTrack, err
}

func deserializeUnrepostTrack(data []byte) (*gen.UnrepostTrack, error) {
	unrepostTrack := &gen.UnrepostTrack{}
	err := proto.Unmarshal(data, unrepostTrack)
	return unrepostTrack, err
}
