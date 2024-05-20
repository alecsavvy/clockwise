/*
services.go

Put all services interfaces in one file unless they need to be broken out.
*/
package services

import (
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/pubsub"
)

type UserPubsub = pubsub.Pubsub[*entities.UserEntity]
type TrackPubsub = pubsub.Pubsub[*entities.TrackEntity]
type FollowPubsub = pubsub.Pubsub[*entities.FollowEntity]
type RepostPubsub = pubsub.Pubsub[*entities.RepostEntity]
