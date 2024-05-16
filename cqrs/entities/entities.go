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
	Genre       string
	Description string
	UserID      string
}

type FollowEntity struct {
	ID          string
	FollowerID  string
	FollowingID string
}

type RepostEntity struct {
	ID         string
	ReposterID string
	TrackID    string
}
