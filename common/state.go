package common

type PeerState struct {
	IsHealthy bool
	NodeType  int
}

// internal state of the application held in memory
type AppState struct {
	Peers map[string]PeerState
}
