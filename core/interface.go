package core

type ManageEntity struct {
	RequestID  string
	UserID     int
	Signer     string
	EntityType string
	EntityID   int
	Action     string
	Metadata   string
}
