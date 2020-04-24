package app

import (
	"encoding/json"
	"io/ioutil"
	"time"

	codecstd "github.com/cosmos/cosmos-sdk/codec/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
)

// GenesisState defines a type alias for the Band genesis application state.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() GenesisState {
	cdc := codecstd.MakeCodec(ModuleBasics)
	denom := "uband"

	stakingGenesis := staking.DefaultGenesisState()
	mintGenesis := mint.DefaultGenesisState()
	govGenesis := gov.DefaultGenesisState()
	crisisGenesis := crisis.DefaultGenesisState()
	slashingGenesis := slashing.DefaultGenesisState()

	stakingGenesis.Params.BondDenom = denom
	stakingGenesis.Params.HistoricalEntries = 1000
	mintGenesis.Params.BlocksPerYear = 10519200 // target 3-second block time
	mintGenesis.Params.MintDenom = denom
	govGenesis.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(denom, sdk.TokensFromConsensusPower(1000)))
	crisisGenesis.ConstantFee = sdk.NewCoin(denom, sdk.TokensFromConsensusPower(10000))
	slashingGenesis.Params.SignedBlocksWindow = 30000                         // approximately 1 day
	slashingGenesis.Params.MinSignedPerWindow = sdk.NewDecWithPrec(5, 2)      // 5%
	slashingGenesis.Params.DowntimeJailDuration = 60 * 10 * time.Second       // 10 minutes
	slashingGenesis.Params.SlashFractionDoubleSign = sdk.NewDecWithPrec(5, 2) // 5%
	slashingGenesis.Params.SlashFractionDowntime = sdk.NewDecWithPrec(1, 4)   // 0.01%

	return GenesisState{
		genutil.ModuleName:    genutil.AppModuleBasic{}.DefaultGenesis(cdc),
		auth.ModuleName:       auth.AppModuleBasic{}.DefaultGenesis(cdc),
		bank.ModuleName:       bank.AppModuleBasic{}.DefaultGenesis(cdc),
		staking.ModuleName:    cdc.MustMarshalJSON(stakingGenesis),
		mint.ModuleName:       cdc.MustMarshalJSON(mintGenesis),
		distr.ModuleName:      distr.AppModuleBasic{}.DefaultGenesis(cdc),
		gov.ModuleName:        cdc.MustMarshalJSON(govGenesis),
		crisis.ModuleName:     cdc.MustMarshalJSON(crisisGenesis),
		slashing.ModuleName:   cdc.MustMarshalJSON(slashingGenesis),
		ibc.ModuleName:        ibc.AppModuleBasic{}.DefaultGenesis(cdc),
		capability.ModuleName: capability.AppModuleBasic{}.DefaultGenesis(cdc),
		upgrade.ModuleName:    upgrade.AppModuleBasic{}.DefaultGenesis(cdc),
		evidence.ModuleName:   evidence.AppModuleBasic{}.DefaultGenesis(cdc),
		transfer.ModuleName:   transfer.AppModuleBasic{}.DefaultGenesis(cdc),
		oracle.ModuleName:     oracle.AppModuleBasic{}.DefaultGenesis(cdc),
	}
}

func GetDefaultDataSourcesAndOracleScripts(owner sdk.AccAddress) json.RawMessage {
	state := oracle.DefaultGenesisState()
	dataSources := []struct {
		name        string
		description string
		path        string
	}{
		{
			"Coingecko script",
			"The Script that queries crypto price from https://coingecko.com",
			"./datasources/coingecko_price.py",
		},
		{
			"Crypto compare script",
			"The Script that queries crypto price from https://cryptocompare.com",
			"./datasources/crypto_compare_price.sh",
		},
		{
			"Binance price",
			"The Script that queries crypto price from https://www.binance.com/en",
			"./datasources/binance_price.sh",
		},
		{
			"Open weather map",
			"The Script that queries weather information from https://api.openweathermap.org",
			"./datasources/open_weather_map.sh",
		},
		{
			"Gold price",
			"The script that queries current gold price",
			"./datasources/gold_price.sh",
		},
		{
			"Alphavantage",
			"The script that queries stock price from Alphavantage",
			"./datasources/alphavantage.sh",
		},
		{
			"Bitcoin block count",
			"The script that queries latest block height of Bitcoin",
			"./datasources/bitcoin_count.sh",
		},
		{
			"Bitcoin block hash",
			"The script that queries block hash of Bitcoin",
			"./datasources/bitcoin_hash.sh",
		},
		{
			"Coingecko volume script",
			"The Script that queries crypto volume from Coingecko",
			"./datasources/coingecko_volume.sh",
		},
		{
			"Crypto compare volume script",
			"The Script that queries crypto volume from Crypto compare",
			"./datasources/crypto_compare_volume.sh",
		},
		{
			"ETH gas station",
			"The Script that queries current Ethereum gas price https://ethgasstation.info",
			"./datasources/ethgasstation.sh",
		},
		{
			"Open sky network",
			"The Script that queries flight information from https://opensky-network.org",
			"./datasources/open_sky_network.sh",
		},
		{
			"Quantum random numbers",
			"The Script that queries array of random number from https://qrng.anu.edu.au",
			"./datasources/qrng_anu.sh",
		},
		{
			"Yahoo finance",
			"The Script that queries stock price from https://finance.yahoo.com",
			"./datasources/yahoo_finance.py",
		},
	}

	// TODO: Find a better way to specify path to data sources
	state.DataSources = make([]oracle.DataSource, len(dataSources))
	for i, dataSource := range dataSources {
		script, err := ioutil.ReadFile(dataSource.path)
		if err != nil {
			panic(err)
		}
		state.DataSources[i] = oracle.NewDataSource(
			owner,
			dataSource.name,
			dataSource.description,
			sdk.Coins{},
			script,
		)
	}

	// TODO: Find a better way to specify path to oracle scripts
	oracleScripts := []struct {
		name          string
		description   string
		path          string
		schema        string
		sourceCodeURL string
	}{
		{
			"Crypto price script (Borsh version)",
			"Oracle script for getting an average crypto price from many sources encoding parameter by borsh.",
			"./owasm/res/crypto_price_borsh.wasm",
			`{"Input": "{ \"kind\": \"struct\", \"fields\": [ [\"symbol\", \"string\"], [\"multiplier\", \"u64\"] ] }", "Output": "{ \"kind\": \"struct\", \"fields\": [ [\"px\", \"u64\"] ] }"}`,
			`https://ipfs.io/ipfs/QmUrYgDKXT8V8DPdCYMEwPM6n82r6zxbvBf6p4gb4m1RA5`,
		},
		{
			"Gold price script",
			"Oracle script for getting an average gold price in ATOM",
			"./owasm/res/gold_price.wasm",
			`{"Input": "{ \"kind\": \"struct\", \"fields\": [ [\"multiplier\", \"u64\"] ] }","Output": "{ \"kind\": \"struct\", \"fields\": [ [\"px\", \"u64\"] ] }"}`,
			`https://ipfs.io/ipfs/QmPheBfYjM4fZ6ngSHYrnDgmapZi9r1i4x5hGFUUyZiP5y`,
		},
		{
			"Alphavantage stock price script",
			"Oracle script for getting an average stock price from Alphavantage",
			"./owasm/res/alphavantage.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"symbol\\", \\"string\\"],[\\"api_key\\", \\"string\\"], [\\"multiplier\\", \\"u64\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"px\\", \\"u64\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
		{
			"Bitcoin block count",
			"Oracle script for getting Bitcoin latest block height",
			"./owasm/res/bitcoin_block_count.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"_unused\\", \\"u8\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"block_count\\", \\"u64\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
		{
			"Bitcoin block hash",
			"Oracle script for getting Bitcoin latest block hash",
			"./owasm/res/bitcoin_block_hash.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"block_height\\", \\"u64\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"block_hash\\", \\"string\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
		{
			"Coingecko crypto volume",
			"Oracle script for getting an average crypto price from Coingecko",
			"./owasm/res/coingecko_volume.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"symbol\\", \\"string\\"], [\\"multiplier\\", \\"u64\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"volume\\", \\"u64\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
		{
			"Crypto compare crypto volume",
			"Oracle script for getting an average crypto price from Crypto compare",
			"./owasm/res/crypto_compare_volume.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"symbol\\", \\"string\\"], [\\"multiplier\\", \\"u64\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"volume\\", \\"u64\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
		{
			"Ethereum gas price",
			"Oracle script for getting gas price from ETH gas station",
			"./owasm/res/eth_gas_station.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"gas_option\\", \\"string\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"gweix10\\", \\"u64\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
		{
			"Open sky network",
			"Oracle script for getting the verification of a flight",
			"./owasm/res/open_sky_network.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"flight_op\\", \\"string\\"],[\\"airport\\", \\"string\\"],[\\"icao24\\", \\"string\\"],[\\"begin\\", \\"string\\"], [\\"end\\", \\"string\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"flight_existence\\", \\"u8\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
		{
			"Open weather map",
			"Oracle script for getting weather information",
			"./owasm/res/open_weather_map.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"country\\", \\"string\\"],[\\"main_field\\", \\"string\\"],[\\"sub_field\\", \\"string\\"],[\\"multiplier\\", \\"u64\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"value\\", \\"u64\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
		{
			"Quantum random number generator",
			"Oracle script for getting a big random number from quantum computer",
			"./owasm/res/qrng.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"size\\", \\"u64\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"random_bytes\\", \\"string\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
		{
			"Yahoo stock price",
			"Oracle script for getting stock price from Yahoo",
			"./owasm/res/yahoo_price.wasm",
			`{"Input":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"symbol\\", \\"string\\"], [\\"multiplier\\", \\"u64\\"] ] }","Output":"{ \\"kind\\": \\"struct\\", \\"fields\\": [ [\\"px\\", \\"u64\\"] ] }"}`,
			`https://bandprotocol.com`,
		},
	}
	state.OracleScripts = make([]oracle.OracleScript, len(oracleScripts))
	for i, oracleScript := range oracleScripts {
		code, err := ioutil.ReadFile(oracleScript.path)
		if err != nil {
			panic(err)
		}
		state.OracleScripts[i] = oracle.NewOracleScript(
			owner,
			oracleScript.name,
			oracleScript.description,
			code,
			oracleScript.schema,
			oracleScript.sourceCodeURL,
		)
	}
	return oracle.ModuleCdc.MustMarshalJSON(state)
}
