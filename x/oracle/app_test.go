package oracle_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/kv"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	"github.com/bandprotocol/bandchain/chain/x/oracle/testapp"
	"github.com/bandprotocol/bandchain/chain/x/oracle/types"
)

func parseEventAttribute(attr interface{}) []byte {
	return []byte(fmt.Sprint(attr))
}

func TestSuccessRequestOracleData(t *testing.T) {
	app, ctx, k := testapp.CreateTestInput()

	ctx = ctx.WithBlockHeight(4).WithBlockTime(time.Unix(int64(1581589790), 0))
	handler := oracle.NewHandler(k)
	requestMsg := types.NewMsgRequestData(types.OracleScriptID(1), []byte("calldata"), 3, 2, "app_test", testapp.Alice.Address)
	res, err := handler(ctx, requestMsg)
	require.NotNil(t, res)
	require.NoError(t, err)

	expectRequest := types.NewRequest(
		types.OracleScriptID(1), []byte("calldata"),
		[]sdk.ValAddress{testapp.Validator3.ValAddress, testapp.Validator1.ValAddress, testapp.Validator2.ValAddress},
		2, 4, 1581589790, "app_test", []types.RawRequest{
			types.NewRawRequest(1, 1, []byte("beeb")),
			types.NewRawRequest(2, 2, []byte("beeb")),
			types.NewRawRequest(3, 3, []byte("beeb")),
		},
	)
	app.EndBlocker(ctx, abci.RequestEndBlock{Height: 4})
	request, err := k.GetRequest(ctx, types.RequestID(1))
	require.Equal(t, expectRequest, request)

	reportMsg1 := types.NewMsgReportData(
		types.RequestID(1), []types.RawReport{
			types.NewRawReport(1, 0, []byte("answer1")),
			types.NewRawReport(2, 0, []byte("answer2")),
			types.NewRawReport(3, 0, []byte("answer3")),
		},
		testapp.Validator1.ValAddress, testapp.Validator1.Address,
	)
	res, err = handler(ctx, reportMsg1)
	require.NotNil(t, res)
	require.NoError(t, err)

	ids := k.GetPendingResolveList(ctx)
	require.Equal(t, []types.RequestID{}, ids)
	_, err = k.GetResult(ctx, types.RequestID(1))
	require.Error(t, err)

	result := app.EndBlocker(ctx, abci.RequestEndBlock{Height: 6})
	expectEvents := []abci.Event{}

	require.Equal(t, expectEvents, result.GetEvents())

	ctx = ctx.WithBlockTime(time.Unix(int64(1581589795), 0))
	reportMsg2 := types.NewMsgReportData(
		types.RequestID(1), []types.RawReport{
			types.NewRawReport(1, 0, []byte("answer1")),
			types.NewRawReport(2, 0, []byte("answer2")),
			types.NewRawReport(3, 0, []byte("answer3")),
		},
		testapp.Validator2.ValAddress, testapp.Validator2.Address,
	)
	res, err = handler(ctx, reportMsg2)
	require.NotNil(t, res)
	require.NoError(t, err)

	ids = k.GetPendingResolveList(ctx)
	require.Equal(t, []types.RequestID{1}, ids)
	_, err = k.GetResult(ctx, types.RequestID(1))
	require.Error(t, err)

	result = app.EndBlocker(ctx, abci.RequestEndBlock{Height: 8})
	reqPacket := types.NewOracleRequestPacketData(
		expectRequest.ClientID, types.OracleScriptID(1), expectRequest.Calldata,
		uint64(len(expectRequest.RequestedValidators)), expectRequest.MinCount,
	)
	resPacket := types.NewOracleResponsePacketData(
		expectRequest.ClientID, types.RequestID(1), 2, expectRequest.RequestTime, 1581589795,
		types.ResolveStatus_Success, []byte("beeb"),
	)
	expectEvents = []abci.Event{{Type: types.EventTypeRequestExecute, Attributes: []kv.Pair{
		{Key: []byte(types.AttributeKeyClientID), Value: parseEventAttribute(reqPacket.ClientID)},
		{Key: []byte(types.AttributeKeyOracleScriptID), Value: parseEventAttribute(reqPacket.OracleScriptID)},
		{Key: []byte(types.AttributeKeyCalldata), Value: expectRequest.Calldata},
		{Key: []byte(types.AttributeKeyAskCount), Value: parseEventAttribute(reqPacket.AskCount)},
		{Key: []byte(types.AttributeKeyMinCount), Value: parseEventAttribute(reqPacket.MinCount)},
		{Key: []byte(types.AttributeKeyRequestID), Value: parseEventAttribute(resPacket.RequestID)},
		{Key: []byte(types.AttributeKeyResolveStatus), Value: parseEventAttribute(uint32(resPacket.ResolveStatus))},
		{Key: []byte(types.AttributeKeyAnsCount), Value: parseEventAttribute(resPacket.AnsCount)},
		{Key: []byte(types.AttributeKeyRequestTime), Value: parseEventAttribute(resPacket.RequestTime)},
		{Key: []byte(types.AttributeKeyResolveTime), Value: parseEventAttribute(resPacket.ResolveTime)},
		{Key: []byte(types.AttributeKeyResult), Value: resPacket.Result},
	}}}

	require.Equal(t, expectEvents, result.GetEvents())

	ids = k.GetPendingResolveList(ctx)
	require.Equal(t, []types.RequestID{}, ids)

	req, err := k.GetRequest(ctx, types.RequestID(1))
	require.NotEqual(t, types.Request{}, req)
	require.NoError(t, err)

	ctx = ctx.WithBlockHeight(32).WithBlockTime(ctx.BlockTime().Add(time.Minute))
	app.EndBlocker(ctx, abci.RequestEndBlock{Height: 32})
}

func TestExpiredRequestOracleData(t *testing.T) {
	app, ctx, k := testapp.CreateTestInput()

	ctx = ctx.WithBlockHeight(4).WithBlockTime(time.Unix(int64(1581589790), 0))
	handler := oracle.NewHandler(k)
	requestMsg := types.NewMsgRequestData(types.OracleScriptID(1), []byte("calldata"), 3, 2, "app_test", testapp.Alice.Address)
	res, err := handler(ctx, requestMsg)
	require.NotNil(t, res)
	require.NoError(t, err)

	expectRequest := types.NewRequest(
		types.OracleScriptID(1), []byte("calldata"),
		[]sdk.ValAddress{testapp.Validator3.ValAddress, testapp.Validator1.ValAddress, testapp.Validator2.ValAddress},
		2, 4, 1581589790, "app_test", []types.RawRequest{
			types.NewRawRequest(1, 1, []byte("beeb")),
			types.NewRawRequest(2, 2, []byte("beeb")),
			types.NewRawRequest(3, 3, []byte("beeb")),
		},
	)
	app.EndBlocker(ctx, abci.RequestEndBlock{Height: 4})
	request, err := k.GetRequest(ctx, types.RequestID(1))
	require.Equal(t, expectRequest, request)

	ctx = ctx.WithBlockHeight(132).WithBlockTime(ctx.BlockTime().Add(time.Minute))
	result := app.EndBlocker(ctx, abci.RequestEndBlock{Height: 132})

	reqPacket := types.NewOracleRequestPacketData(
		expectRequest.ClientID, types.OracleScriptID(1), expectRequest.Calldata,
		uint64(len(expectRequest.RequestedValidators)), expectRequest.MinCount,
	)
	resPacket := types.NewOracleResponsePacketData(
		expectRequest.ClientID, types.RequestID(1), 0, expectRequest.RequestTime, ctx.BlockTime().Unix(),
		types.ResolveStatus_Expired, []byte{},
	)
	expectEvents := []abci.Event{{
		Type: types.EventTypeRequestExecute,
		Attributes: []kv.Pair{
			{Key: []byte(types.AttributeKeyClientID), Value: parseEventAttribute(reqPacket.ClientID)},
			{Key: []byte(types.AttributeKeyOracleScriptID), Value: parseEventAttribute(reqPacket.OracleScriptID)},
			{Key: []byte(types.AttributeKeyCalldata), Value: expectRequest.Calldata},
			{Key: []byte(types.AttributeKeyAskCount), Value: parseEventAttribute(reqPacket.AskCount)},
			{Key: []byte(types.AttributeKeyMinCount), Value: parseEventAttribute(reqPacket.MinCount)},
			{Key: []byte(types.AttributeKeyRequestID), Value: parseEventAttribute(resPacket.RequestID)},
			{Key: []byte(types.AttributeKeyResolveStatus), Value: parseEventAttribute(uint32(resPacket.ResolveStatus))},
			{Key: []byte(types.AttributeKeyAnsCount), Value: parseEventAttribute(resPacket.AnsCount)},
			{Key: []byte(types.AttributeKeyRequestTime), Value: parseEventAttribute(resPacket.RequestTime)},
			{Key: []byte(types.AttributeKeyResolveTime), Value: parseEventAttribute(resPacket.ResolveTime)},
			{Key: []byte(types.AttributeKeyResult), Value: resPacket.Result},
		},
	}, {
		Type: types.EventTypeDeactivate,
		Attributes: []kv.Pair{
			{Key: []byte(types.AttributeKeyValidator), Value: parseEventAttribute(testapp.Validator3.ValAddress.String())},
		},
	}, {
		Type: types.EventTypeDeactivate,
		Attributes: []kv.Pair{
			{Key: []byte(types.AttributeKeyValidator), Value: parseEventAttribute(testapp.Validator1.ValAddress.String())},
		},
	}, {
		Type: types.EventTypeDeactivate,
		Attributes: []kv.Pair{
			{Key: []byte(types.AttributeKeyValidator), Value: parseEventAttribute(testapp.Validator2.ValAddress.String())},
		},
	}}

	require.Equal(t, expectEvents, result.GetEvents())
}
