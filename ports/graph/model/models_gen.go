// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type TrackEvents interface {
	IsTrackEvents()
}

type UserEvents interface {
	IsUserEvents()
}

type Follow struct {
	ID          string `json:"id"`
	FollowerID  string `json:"followerId"`
	FollowingID string `json:"followingId"`
}

func (Follow) IsUserEvents() {}

type Mutation struct {
}

type NewFollow struct {
	FollowerID  string `json:"followerId"`
	FollowingID string `json:"followingId"`
}

type NewRepost struct {
	ReposterID string `json:"reposterId"`
	TrackID    string `json:"trackId"`
}

type NewTrack struct {
	Title       string `json:"title"`
	StreamURL   string `json:"streamUrl"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	UserID      string `json:"userId"`
}

type NewUser struct {
	Handle  string `json:"handle"`
	Bio     string `json:"bio"`
	Address string `json:"address"`
}

type Query struct {
}

type Repost struct {
	ID         string `json:"id"`
	ReposterID string `json:"reposterId"`
	TrackID    string `json:"trackId"`
}

func (Repost) IsUserEvents() {}

func (Repost) IsTrackEvents() {}

type Subscription struct {
}

type Track struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	StreamURL   string `json:"streamUrl"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
}

func (Track) IsUserEvents() {}

func (Track) IsTrackEvents() {}

type UpdateTrack struct {
	StreamURL   *string `json:"streamUrl,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateUser struct {
	Bio     *string `json:"bio,omitempty"`
	Address *string `json:"address,omitempty"`
}

type User struct {
	ID      string `json:"id"`
	Handle  string `json:"handle"`
	Bio     string `json:"bio"`
	Address string `json:"address"`
}

func (User) IsUserEvents() {}
