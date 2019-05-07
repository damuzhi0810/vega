package config

import (
	"code.vegaprotocol.io/vega/internal/api"
	"code.vegaprotocol.io/vega/internal/blockchain"
	"code.vegaprotocol.io/vega/internal/candles"
	"code.vegaprotocol.io/vega/internal/execution"
	"code.vegaprotocol.io/vega/internal/gateway"
	"code.vegaprotocol.io/vega/internal/logging"
	"code.vegaprotocol.io/vega/internal/markets"
	"code.vegaprotocol.io/vega/internal/monitoring"
	"code.vegaprotocol.io/vega/internal/orders"
	"code.vegaprotocol.io/vega/internal/parties"
	"code.vegaprotocol.io/vega/internal/pprof"
	"code.vegaprotocol.io/vega/internal/storage"
	"code.vegaprotocol.io/vega/internal/trades"
	"code.vegaprotocol.io/vega/internal/vegatime"
)

// Config ties together all other application configuration types.
type Config struct {
	API        api.Config
	Blockchain blockchain.Config
	Candles    candles.Config
	Execution  execution.Config
	Logging    logging.Config
	Markets    markets.Config
	Orders     orders.Config
	Parties    parties.Config
	Storage    storage.Config
	Trades     trades.Config
	Time       vegatime.Config
	Monitoring monitoring.Config
	Gateway    gateway.Config

	Pprof          pprof.Config
	GatewayEnabled bool
}

// NewDefaultConfig returns a set of default configs for all vega packages, as specified at the per package
// config level, if there is an error initialising any of the configs then this is returned.
func NewDefaultConfig(defaultStoreDirPath string) Config {
	return Config{
		Trades:         trades.NewDefaultConfig(),
		Blockchain:     blockchain.NewDefaultConfig(),
		Execution:      execution.NewDefaultConfig(defaultStoreDirPath),
		API:            api.NewDefaultConfig(),
		Orders:         orders.NewDefaultConfig(),
		Time:           vegatime.NewDefaultConfig(),
		Markets:        markets.NewDefaultConfig(),
		Parties:        parties.NewDefaultConfig(),
		Candles:        candles.NewDefaultConfig(),
		Storage:        storage.NewDefaultConfig(defaultStoreDirPath),
		Pprof:          pprof.NewDefaultConfig(),
		Monitoring:     monitoring.NewDefaultConfig(),
		Logging:        logging.NewDefaultConfig(),
		Gateway:        gateway.NewDefaultConfig(),
		GatewayEnabled: true,
	}
}