package core

import (
	corev1 "github.com/alecsavvy/clockwise/api/oap/core/v1"
	v1 "github.com/cometbft/cometbft/api/cometbft/types/v1"
)

type CoreService interface {
	SendTransaction(*corev1.SignedTransaction) (*corev1.SignedTransaction, error)
	GetTransaction(txHash string) (*corev1.SignedTransaction, error)
	GetBlock(height int64) (*v1.Block, error)
	StreamBlocks() <-chan *v1.Block
	StreamTransactions() <-chan *corev1.SignedTransaction
}

var _ CoreService = (*Core)(nil)

// GetBlock implements CoreService.
func (c *Core) GetBlock(height int64) (*v1.Block, error) {
	panic("unimplemented")
}

// GetTransaction implements CoreService.
func (c *Core) GetTransaction(txHash string) (*corev1.SignedTransaction, error) {
	panic("unimplemented")
}

// SendTransaction implements CoreService.
func (c *Core) SendTransaction(*corev1.SignedTransaction) (*corev1.SignedTransaction, error) {
	panic("unimplemented")
}

// StreamBlocks implements CoreService.
func (c *Core) StreamBlocks() <-chan *v1.Block {
	panic("unimplemented")
}

// StreamTransactions implements CoreService.
func (c *Core) StreamTransactions() <-chan *corev1.SignedTransaction {
	panic("unimplemented")
}
