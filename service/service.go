package service

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Service interface {
	Name() string
	Init(ctx context.Context) error
	Ready() <-chan struct{}
	MarkReady()
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Err() <-chan error
}

type BaseService struct {
	name   string
	errCh  chan error
	ready  chan struct{}
	ctx    context.Context
	cancel context.CancelFunc

	mu      sync.Mutex
	started bool
	stopped bool
	wg      errgroup.Group
}

func NewBaseService(name string) *BaseService {
	return &BaseService{
		name:  name,
		errCh: make(chan error, 1),
		ready: make(chan struct{}),
	}
}

func (b *BaseService) Name() string           { return b.name }
func (b *BaseService) Ready() <-chan struct{} { return b.ready }
func (b *BaseService) Err() <-chan error      { return b.errCh }

func (b *BaseService) Init(ctx context.Context) error { return nil }

func (b *BaseService) Start(ctx context.Context) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.started {
		return fmt.Errorf("%s already started", b.name)
	}
	b.ctx, b.cancel = context.WithCancel(ctx)
	b.started = true
	b.stopped = false

	return nil
}

func (b *BaseService) Stop(ctx context.Context) error {
	b.mu.Lock()
	if !b.started || b.stopped {
		b.mu.Unlock()
		return nil
	}
	b.stopped = true
	b.cancel()
	b.mu.Unlock()

	done := make(chan struct{})
	go func() {
		b.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (b *BaseService) MarkReady() {
	select {
	case <-b.ready:
	default:
		close(b.ready)
	}
}

// Run starts a goroutine managed by the BaseService using WaitGroup.Go (Go 1.25+).
// Any panic or returned error is sent to Err().
func (b *BaseService) Run(fn func(ctx context.Context) error) {
	b.wg.Go(func() error {
		defer func() {
			if r := recover(); r != nil {
				select {
				case b.errCh <- fmt.Errorf("panic in %s: %v", b.name, r):
				default:
				}
			}
		}()
		if err := fn(b.ctx); err != nil {
			select {
			case b.errCh <- err:
			default:
			}
			return err
		}
		return nil
	})
}
