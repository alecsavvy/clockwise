package pubsub

import "sync"

type Pubsub[Message any] struct {
	subscribers map[chan Message]struct{}
	mu          sync.RWMutex
}

func NewPubsub[Message any]() *Pubsub[Message] {
	return &Pubsub[Message]{
		subscribers: make(map[chan Message]struct{}),
	}
}

func (ps *Pubsub[Message]) Subscribe() chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan Message)
	ps.subscribers[ch] = struct{}{}
	return ch
}

func (ps *Pubsub[Message]) Unsubscribe(ch chan Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	delete(ps.subscribers, ch)
	close(ch)
}

func (ps *Pubsub[Message]) Publish(msg Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for ch := range ps.subscribers {
		ch := ch
		// send in goroutine so slow receivers don't block
		go func() {
			select {
			case ch <- msg:
				// message sent successfully
			default:
				// subscriber is not ready, drop the message
			}
		}()
	}
}
