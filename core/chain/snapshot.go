/*
snapshot.go

Contains logic for the snapshot steps of consensus.
*/
package chain

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// ListSnapshots implements types.Application.
func (a *Application) ListSnapshots(context.Context, *abcitypes.RequestListSnapshots) (*abcitypes.ResponseListSnapshots, error) {
	return &abcitypes.ResponseListSnapshots{}, nil
}

// LoadSnapshotChunk implements types.Application.
func (a *Application) LoadSnapshotChunk(context.Context, *abcitypes.RequestLoadSnapshotChunk) (*abcitypes.ResponseLoadSnapshotChunk, error) {
	return &abcitypes.ResponseLoadSnapshotChunk{}, nil
}

// OfferSnapshot implements types.Application.
func (a *Application) OfferSnapshot(context.Context, *abcitypes.RequestOfferSnapshot) (*abcitypes.ResponseOfferSnapshot, error) {
	return &abcitypes.ResponseOfferSnapshot{}, nil
}

// ApplySnapshotChunk implements types.Application.
func (a *Application) ApplySnapshotChunk(context.Context, *abcitypes.RequestApplySnapshotChunk) (*abcitypes.ResponseApplySnapshotChunk, error) {
	return &abcitypes.ResponseApplySnapshotChunk{}, nil
}
