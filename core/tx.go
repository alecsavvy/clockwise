/* utilities specific to the core, not intended for external use */
package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/rpc/client/local"
	"google.golang.org/protobuf/proto"
)

// returns the original
func SendTx[T proto.Message](logger *utils.Logger, rpc *local.Local, msg T) error {
	ctx := context.Background()

	tx, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	result, err := rpc.BroadcastTxSync(ctx, tx)
	if err != nil {
		return err
	}

	if result.Code != abcitypes.CodeTypeOK {
		return errors.New(result.Log)
	}

	txChan, err := rpc.Subscribe(ctx, "tx-subscriber", fmt.Sprintf("tm.event = 'Tx' AND tx.hash = '%X'", result.Hash))
	if err != nil {
		return err
	}

	defer func() {
		if err := rpc.Unsubscribe(ctx, "tx-subscriber", fmt.Sprintf("tm.event = 'Tx' AND tx.hash = '%X'", result.Hash)); err != nil {
			// Handle the unsubscribe error if necessary
			fmt.Println("Failed to unsubscribe:", err)
		}
	}()

	select {
	case _ = <-txChan:
		return nil
	case <-time.After(30 * time.Second):
		return errors.New("tx waiting timeout")
	}
}
