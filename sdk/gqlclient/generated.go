// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package gqlclient

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// CreateTrackCreateTrack includes the requested fields of the GraphQL type Track.
type CreateTrackCreateTrack struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StreamUrl   string `json:"streamUrl"`
	UserId      string `json:"userId"`
}

// GetId returns CreateTrackCreateTrack.Id, and is useful for accessing the field via an interface.
func (v *CreateTrackCreateTrack) GetId() string { return v.Id }

// GetTitle returns CreateTrackCreateTrack.Title, and is useful for accessing the field via an interface.
func (v *CreateTrackCreateTrack) GetTitle() string { return v.Title }

// GetDescription returns CreateTrackCreateTrack.Description, and is useful for accessing the field via an interface.
func (v *CreateTrackCreateTrack) GetDescription() string { return v.Description }

// GetStreamUrl returns CreateTrackCreateTrack.StreamUrl, and is useful for accessing the field via an interface.
func (v *CreateTrackCreateTrack) GetStreamUrl() string { return v.StreamUrl }

// GetUserId returns CreateTrackCreateTrack.UserId, and is useful for accessing the field via an interface.
func (v *CreateTrackCreateTrack) GetUserId() string { return v.UserId }

// CreateTrackResponse is returned by CreateTrack on success.
type CreateTrackResponse struct {
	CreateTrack CreateTrackCreateTrack `json:"createTrack"`
}

// GetCreateTrack returns CreateTrackResponse.CreateTrack, and is useful for accessing the field via an interface.
func (v *CreateTrackResponse) GetCreateTrack() CreateTrackCreateTrack { return v.CreateTrack }

// CreateUserCreateUser includes the requested fields of the GraphQL type User.
type CreateUserCreateUser struct {
	Handle  string `json:"handle"`
	Address string `json:"address"`
	Bio     string `json:"bio"`
	Txhash  string `json:"txhash"`
}

// GetHandle returns CreateUserCreateUser.Handle, and is useful for accessing the field via an interface.
func (v *CreateUserCreateUser) GetHandle() string { return v.Handle }

// GetAddress returns CreateUserCreateUser.Address, and is useful for accessing the field via an interface.
func (v *CreateUserCreateUser) GetAddress() string { return v.Address }

// GetBio returns CreateUserCreateUser.Bio, and is useful for accessing the field via an interface.
func (v *CreateUserCreateUser) GetBio() string { return v.Bio }

// GetTxhash returns CreateUserCreateUser.Txhash, and is useful for accessing the field via an interface.
func (v *CreateUserCreateUser) GetTxhash() string { return v.Txhash }

// CreateUserResponse is returned by CreateUser on success.
type CreateUserResponse struct {
	CreateUser CreateUserCreateUser `json:"createUser"`
}

// GetCreateUser returns CreateUserResponse.CreateUser, and is useful for accessing the field via an interface.
func (v *CreateUserResponse) GetCreateUser() CreateUserCreateUser { return v.CreateUser }

// FollowFollowUserFollow includes the requested fields of the GraphQL type Follow.
type FollowFollowUserFollow struct {
	FollowerId string `json:"followerId"`
	FolloweeId string `json:"followeeId"`
}

// GetFollowerId returns FollowFollowUserFollow.FollowerId, and is useful for accessing the field via an interface.
func (v *FollowFollowUserFollow) GetFollowerId() string { return v.FollowerId }

// GetFolloweeId returns FollowFollowUserFollow.FolloweeId, and is useful for accessing the field via an interface.
func (v *FollowFollowUserFollow) GetFolloweeId() string { return v.FolloweeId }

// FollowResponse is returned by Follow on success.
type FollowResponse struct {
	FollowUser FollowFollowUserFollow `json:"followUser"`
}

// GetFollowUser returns FollowResponse.FollowUser, and is useful for accessing the field via an interface.
func (v *FollowResponse) GetFollowUser() FollowFollowUserFollow { return v.FollowUser }

// RepostRepostTrackRepost includes the requested fields of the GraphQL type Repost.
type RepostRepostTrackRepost struct {
	TrackId    string `json:"trackId"`
	ReposterId string `json:"reposterId"`
}

// GetTrackId returns RepostRepostTrackRepost.TrackId, and is useful for accessing the field via an interface.
func (v *RepostRepostTrackRepost) GetTrackId() string { return v.TrackId }

// GetReposterId returns RepostRepostTrackRepost.ReposterId, and is useful for accessing the field via an interface.
func (v *RepostRepostTrackRepost) GetReposterId() string { return v.ReposterId }

// RepostResponse is returned by Repost on success.
type RepostResponse struct {
	RepostTrack RepostRepostTrackRepost `json:"repostTrack"`
}

// GetRepostTrack returns RepostResponse.RepostTrack, and is useful for accessing the field via an interface.
func (v *RepostResponse) GetRepostTrack() RepostRepostTrackRepost { return v.RepostTrack }

// __CreateTrackInput is used internally by genqlient
type __CreateTrackInput struct {
	Title       string `json:"title"`
	StreamUrl   string `json:"streamUrl"`
	UserId      string `json:"userId"`
	Description string `json:"description"`
}

// GetTitle returns __CreateTrackInput.Title, and is useful for accessing the field via an interface.
func (v *__CreateTrackInput) GetTitle() string { return v.Title }

// GetStreamUrl returns __CreateTrackInput.StreamUrl, and is useful for accessing the field via an interface.
func (v *__CreateTrackInput) GetStreamUrl() string { return v.StreamUrl }

// GetUserId returns __CreateTrackInput.UserId, and is useful for accessing the field via an interface.
func (v *__CreateTrackInput) GetUserId() string { return v.UserId }

// GetDescription returns __CreateTrackInput.Description, and is useful for accessing the field via an interface.
func (v *__CreateTrackInput) GetDescription() string { return v.Description }

// __CreateUserInput is used internally by genqlient
type __CreateUserInput struct {
	Handle  string `json:"handle"`
	Address string `json:"address"`
	Bio     string `json:"bio"`
}

// GetHandle returns __CreateUserInput.Handle, and is useful for accessing the field via an interface.
func (v *__CreateUserInput) GetHandle() string { return v.Handle }

// GetAddress returns __CreateUserInput.Address, and is useful for accessing the field via an interface.
func (v *__CreateUserInput) GetAddress() string { return v.Address }

// GetBio returns __CreateUserInput.Bio, and is useful for accessing the field via an interface.
func (v *__CreateUserInput) GetBio() string { return v.Bio }

// __FollowInput is used internally by genqlient
type __FollowInput struct {
	FollowerId string `json:"followerId"`
	FolloweeId string `json:"followeeId"`
}

// GetFollowerId returns __FollowInput.FollowerId, and is useful for accessing the field via an interface.
func (v *__FollowInput) GetFollowerId() string { return v.FollowerId }

// GetFolloweeId returns __FollowInput.FolloweeId, and is useful for accessing the field via an interface.
func (v *__FollowInput) GetFolloweeId() string { return v.FolloweeId }

// __RepostInput is used internally by genqlient
type __RepostInput struct {
	TrackId    string `json:"trackId"`
	ReposterId string `json:"reposterId"`
}

// GetTrackId returns __RepostInput.TrackId, and is useful for accessing the field via an interface.
func (v *__RepostInput) GetTrackId() string { return v.TrackId }

// GetReposterId returns __RepostInput.ReposterId, and is useful for accessing the field via an interface.
func (v *__RepostInput) GetReposterId() string { return v.ReposterId }

// The query or mutation executed by CreateTrack.
const CreateTrack_Operation = `
mutation CreateTrack ($title: String!, $streamUrl: String!, $userId: String!, $description: String!) {
	createTrack(input: {title:$title,streamUrl:$streamUrl,userId:$userId,description:$description}) {
		id
		title
		description
		streamUrl
		userId
	}
}
`

func CreateTrack(
	ctx_ context.Context,
	client_ graphql.Client,
	title string,
	streamUrl string,
	userId string,
	description string,
) (*CreateTrackResponse, error) {
	req_ := &graphql.Request{
		OpName: "CreateTrack",
		Query:  CreateTrack_Operation,
		Variables: &__CreateTrackInput{
			Title:       title,
			StreamUrl:   streamUrl,
			UserId:      userId,
			Description: description,
		},
	}
	var err_ error

	var data_ CreateTrackResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by CreateUser.
const CreateUser_Operation = `
mutation CreateUser ($handle: String!, $address: String!, $bio: String!) {
	createUser(input: {handle:$handle,address:$address,bio:$bio}) {
		handle
		address
		bio
		txhash
	}
}
`

func CreateUser(
	ctx_ context.Context,
	client_ graphql.Client,
	handle string,
	address string,
	bio string,
) (*CreateUserResponse, error) {
	req_ := &graphql.Request{
		OpName: "CreateUser",
		Query:  CreateUser_Operation,
		Variables: &__CreateUserInput{
			Handle:  handle,
			Address: address,
			Bio:     bio,
		},
	}
	var err_ error

	var data_ CreateUserResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by Follow.
const Follow_Operation = `
mutation Follow ($followerId: String!, $followeeId: String!) {
	followUser(input: {followerId:$followerId,followeeId:$followeeId}) {
		followerId
		followeeId
	}
}
`

func Follow(
	ctx_ context.Context,
	client_ graphql.Client,
	followerId string,
	followeeId string,
) (*FollowResponse, error) {
	req_ := &graphql.Request{
		OpName: "Follow",
		Query:  Follow_Operation,
		Variables: &__FollowInput{
			FollowerId: followerId,
			FolloweeId: followeeId,
		},
	}
	var err_ error

	var data_ FollowResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by Repost.
const Repost_Operation = `
mutation Repost ($trackId: String!, $reposterId: String!) {
	repostTrack(input: {trackId:$trackId,reposterId:$reposterId}) {
		trackId
		reposterId
	}
}
`

func Repost(
	ctx_ context.Context,
	client_ graphql.Client,
	trackId string,
	reposterId string,
) (*RepostResponse, error) {
	req_ := &graphql.Request{
		OpName: "Repost",
		Query:  Repost_Operation,
		Variables: &__RepostInput{
			TrackId:    trackId,
			ReposterId: reposterId,
		},
	}
	var err_ error

	var data_ RepostResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
