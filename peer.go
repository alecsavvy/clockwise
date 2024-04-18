package main

type Peer struct {
	host string
}

func NewPeer(host string) Peer {
	return Peer{host: host}
}
