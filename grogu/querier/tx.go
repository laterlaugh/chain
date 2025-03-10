package querier

import (
	rpcclient "github.com/cometbft/cometbft/rpc/client"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type TxQuerier struct {
	clientCtxs []client.Context
}

func NewTxQuerier(clientCtx client.Context, clients []rpcclient.RemoteClient) *TxQuerier {
	clientCtxs := make([]client.Context, 0, len(clients))
	for _, cl := range clients {
		clientCtxs = append(clientCtxs, clientCtx.WithClient(cl))
	}
	return &TxQuerier{clientCtxs}
}

func (q *TxQuerier) QueryTx(hash string) (*types.TxResponse, error) {
	resultCh := make(chan *types.TxResponse, len(q.clientCtxs))
	failureCh := make(chan error, len(q.clientCtxs))

	for _, ctx := range q.clientCtxs {
		go func(ctx client.Context) {
			resp, err := tx.QueryTx(ctx, hash)
			if err != nil {
				failureCh <- err
				return
			}

			resultCh <- resp
		}(ctx)
	}

	var err error
	for range q.clientCtxs {
		select {
		case res := <-resultCh:
			return res, nil
		case err = <-failureCh:
			continue
		}
	}

	return nil, err
}
