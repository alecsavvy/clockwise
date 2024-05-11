package core

type UserRepository interface {
	CreateUser(*UserEntity) (*UserEntity, error)
	GetUser(handle string) (*UserEntity, error)
	GetUsers() ([]*UserEntity, error)
}

type TrackRepository interface {
	CreateTrack(*TrackEntity) (*TrackEntity, error)
	GetTrack(title string) (*TrackEntity, error)
	GetTracks() ([]*TrackEntity, error)
}
