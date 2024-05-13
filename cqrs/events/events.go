package events

import "github.com/alecsavvy/clockwise/cqrs/entities"

type EventMetadata struct {
	BlockHeight     uint64
	TransactionHash string
}

type UserCreatedEvent struct {
	EventMetadata
	User entities.UserEntity
}

type TrackCreatedEvent struct {
	EventMetadata
	Track entities.TrackEntity
}

type FollowCreatedEvent struct {
	EventMetadata
	Follow entities.FollowEntity
}

type RepostCreatedEvent struct {
	EventMetadata
	Repost entities.RepostEntity
}
