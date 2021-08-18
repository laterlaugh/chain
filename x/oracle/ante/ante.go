package ante

import (
	"fmt"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	lru "github.com/hashicorp/golang-lru"

	"github.com/bandprotocol/chain/v2/x/oracle/keeper"
	"github.com/bandprotocol/chain/v2/x/oracle/types"
)

var (
	repTxCount       *lru.Cache
	mu               sync.Mutex
	nextRepOnlyBlock int64
)

func init() {
	var err error
	repTxCount, err = lru.New(20000)
	if err != nil {
		panic(err)
	}
}

// TODO: Add test on this ante before merge

func checkValidReportMsg(ctx sdk.Context, oracleKeeper keeper.Keeper, r *types.MsgReportData) error {
	validator, err := sdk.ValAddressFromBech32(r.Validator)
	if err != nil {
		return err
	}
	report := types.NewReport(validator, false, r.RawReports)
	return oracleKeeper.CheckValidReport(ctx, r.RequestID, report)
}

func updateCache(val string, rid types.RequestID) (trigger bool) {
	key := fmt.Sprintf("%s:%d", val, rid)
	mu.Lock()
	defer mu.Unlock()
	value, ok := repTxCount.Get(key)
	nextVal := 1
	if ok {
		nextVal = value.(int) + 1
	}
	repTxCount.Add(key, nextVal)
	return nextVal > 20
}

// NewFeelessReportsAnteHandler returns a new ante handler that waives minimum gas price
// requirement if the incoming tx is a valid report transaction.
func NewFeelessReportsAnteHandler(ante sdk.AnteHandler, oracleKeeper keeper.Keeper) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
		if ctx.IsCheckTx() && !simulate {
			isRepOnlyBlock := ctx.BlockHeight() == nextRepOnlyBlock
			isValidReportTx := true
			for _, msg := range tx.GetMsgs() {
				// Check direct report msg
				if dr, ok := msg.(*types.MsgReportData); ok {
					// Check if it's not valid report msg, discard this transaction
					if err := checkValidReportMsg(ctx, oracleKeeper, dr); err != nil {
						return ctx, err
					}
					if !isRepOnlyBlock {
						if updateCache(dr.Validator, dr.RequestID) {
							nextRepOnlyBlock = ctx.BlockHeight() + 1
						}
					}
				} else {
					// Check is the MsgExec from reporter
					me, ok := msg.(*authz.MsgExec)
					if !ok {
						isValidReportTx = false
						break
					}

					// If cannot get message, then pretend as non-free transaction
					msgs, err := me.GetMessages()
					if err != nil {
						isValidReportTx = false
						break
					}

					grantee, err := sdk.AccAddressFromBech32(me.Grantee)
					if err != nil {
						isValidReportTx = false
						break
					}

					allValidReportMsg := true
					for _, m := range msgs {
						r, ok := m.(*types.MsgReportData)
						// If this is not report msg, skip other msgs on this exec msg
						if !ok {
							allValidReportMsg = false
							break
						}

						// Fail to parse validator, then discard this transaction
						validator, err := sdk.ValAddressFromBech32(r.Validator)
						if err != nil {
							return ctx, err
						}

						// If this grantee is not a reporter of validator, then discard this transaction
						if !oracleKeeper.IsReporter(ctx, validator, grantee) {
							return ctx, sdkerrors.ErrUnauthorized.Wrap("authorization not found")
						}

						// Check if it's not valid report msg, discard this transaction
						if err := checkValidReportMsg(ctx, oracleKeeper, r); err != nil {
							return ctx, err
						}

						// Update cache in case it's a valid report
						if !isRepOnlyBlock {
							if updateCache(r.Validator, r.RequestID) {
								nextRepOnlyBlock = ctx.BlockHeight() + 1
							}
						}
					}

					// If this exec msg has other non-report msg, disable feeless and skip other msgs in tx
					if !allValidReportMsg {
						isValidReportTx = false
						break
					}

				}
			}
			if isRepOnlyBlock && !isValidReportTx {
				return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Block reserved for report txs")
			}
			if isValidReportTx {
				minGas := ctx.MinGasPrices()
				newCtx, err := ante(ctx.WithMinGasPrices(sdk.DecCoins{}), tx, simulate)
				// Set minimum gas price context and return context to caller.
				return newCtx.WithMinGasPrices(minGas), err
			}
		}
		return ante(ctx, tx, simulate)
	}
}
