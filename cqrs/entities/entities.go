package entities

type UserEntity struct {
	ID      string
	Handle  string
	Address string
	Bio     string
}

type TrackEntity struct {
	ID          string
	Title       string
	StreamURL   string
	Description string
	UserID      string
}

type FollowEntity struct {
	FollowerID  string
	FollowingID string
}

type RepostEntity struct {
	ReposterID string
	TrackID    string
}
