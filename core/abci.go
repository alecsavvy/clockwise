package core

import (
	"context"
	"errors"
	"fmt"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

const (
	CodeTypeNotOK = 1
)

// Collects info about the chain node and returns it to other nodes.
func (c *Core) Info(context.Context, *abcitypes.RequestInfo) (*abcitypes.ResponseInfo, error) {
	return &abcitypes.ResponseInfo{}, nil
}

// Instructions for how the chain should initialize.
func (c *Core) InitChain(context.Context, *abcitypes.RequestInitChain) (*abcitypes.ResponseInitChain, error) {
	return &abcitypes.ResponseInitChain{}, nil
}

// Performs validation on a proposed transaction, should be very performant as this check
// gets called a lot (per the cometbft docs)
func (c *Core) CheckTx(ctx context.Context, req *abcitypes.RequestCheckTx) (*abcitypes.ResponseCheckTx, error) {
	if err := CheckTx(c, ctx, req); err != nil {
		return &abcitypes.ResponseCheckTx{Code: CodeTypeNotOK, Log: err.Error()}, nil
	}
	return &abcitypes.ResponseCheckTx{Code: abcitypes.CodeTypeOK}, nil
}

// Prepares a new block proposal for the network
func (c *Core) PrepareProposal(_ context.Context, proposal *abcitypes.RequestPrepareProposal) (*abcitypes.ResponsePrepareProposal, error) {
	// TODO: reorder transactions in here
	return &abcitypes.ResponsePrepareProposal{Txs: proposal.Txs}, nil
}

// Processes block proposal from the network created by PrepareProposal
func (c *Core) ProcessProposal(context.Context, *abcitypes.RequestProcessProposal) (*abcitypes.ResponseProcessProposal, error) {
	return &abcitypes.ResponseProcessProposal{Status: abcitypes.ResponseProcessProposal_ACCEPT}, nil
}

// Prepares a block for commitment and provides final validation
func (c *Core) FinalizeBlock(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) (*abcitypes.ResponseFinalizeBlock, error) {
	var txResults = make([]*abcitypes.ExecTxResult, len(rfb.Txs))
	for i, tx := range rfb.Txs {
		var me ManageEntity
		c.fromTxBytes(tx, &me)
		txResults[i] = &abcitypes.ExecTxResult{
			Code: abcitypes.CodeTypeOK,
			Events: []abcitypes.Event{
				{
					Type: fmt.Sprintf("%s%s", me.EntityType, me.Action),
					Attributes: []abcitypes.EventAttribute{
						{Key: "requestId", Value: me.RequestID},
					},
				},
			},
		}
	}
	return &abcitypes.ResponseFinalizeBlock{TxResults: txResults}, nil
}

// Writes the state changes to the database after checking and finalizing a block
func (c *Core) Commit(ctx context.Context, req *abcitypes.RequestCommit) (*abcitypes.ResponseCommit, error) {
	resp := &abcitypes.ResponseCommit{}
	/**
	// TODO: check if indexer is up to date here, only prune once indexer is up to date.
	// i.e. a node can seed postgres after indexing
	if app.RetainBlocks > 0 && app.state.Height >= app.RetainBlocks {
		resp.RetainHeight = app.state.Height - app.RetainBlocks + 1
	}
	*/
	latestBlock, err := c.rpc.Block(ctx, nil)
	if err != nil {
		return nil, err
	}
	if c.RetainBlocks > 0 && latestBlock.Block.Height >= c.RetainBlocks {
		resp.RetainHeight = latestBlock.Block.Height - c.RetainBlocks + 1
	}
	return resp, nil
}

// Reads from the chain state, may be irrelevant with the higher level GQL interface
func (c *Core) Query(_ context.Context, req *abcitypes.RequestQuery) (*abcitypes.ResponseQuery, error) {
	return nil, errors.New("internal query not supported")
}

// ListSnapshots implements types.Application.
func (c *Core) ListSnapshots(context.Context, *abcitypes.RequestListSnapshots) (*abcitypes.ResponseListSnapshots, error) {
	return &abcitypes.ResponseListSnapshots{}, nil
}

// LoadSnapshotChunk implements types.Application.
func (c *Core) LoadSnapshotChunk(context.Context, *abcitypes.RequestLoadSnapshotChunk) (*abcitypes.ResponseLoadSnapshotChunk, error) {
	return &abcitypes.ResponseLoadSnapshotChunk{}, nil
}

// OfferSnapshot implements types.Application.
func (c *Core) OfferSnapshot(context.Context, *abcitypes.RequestOfferSnapshot) (*abcitypes.ResponseOfferSnapshot, error) {
	return &abcitypes.ResponseOfferSnapshot{}, nil
}

// ApplySnapshotChunk implements types.Application.
func (c *Core) ApplySnapshotChunk(context.Context, *abcitypes.RequestApplySnapshotChunk) (*abcitypes.ResponseApplySnapshotChunk, error) {
	return &abcitypes.ResponseApplySnapshotChunk{}, nil
}
