package position

import (
	"fmt"
	"sync"

	"code.vegaprotocol.io/vega/internal/logging"
	types "code.vegaprotocol.io/vega/proto"
)

type MarketPosition struct {
	size    int64
	margins map[string]uint64
	partyID string
}

// Margins returns a copy of the current margins map
func (m *MarketPosition) Margins() map[string]uint64 {
	return m.margins
}

// UpdateMargin updates the margin value for a single asset
func (m *MarketPosition) UpdateMargin(assetID string, margin uint64) {
	m.margins[assetID] = margin
}

func (m MarketPosition) Size() int64 {
	return m.size
}

func (m MarketPosition) Party() string {
	return m.partyID
}

type Engine struct {
	*Config

	mu *sync.RWMutex
	// partyID -> MarketPosition
	positions map[string]*MarketPosition
}

func New(config *Config) *Engine {
	return &Engine{
		mu:        &sync.RWMutex{},
		Config:    config,
		positions: map[string]*MarketPosition{},
	}
}

func (e *Engine) Update(trade *types.Trade) {

	e.mu.Lock()
	// todo(cdm): overflow should be managed at the trade/order creation point. We shouldn't accept an order onto
	// your book that would overflow your position. Order validation requires position store/state lookup.

	buyer, ok := e.positions[trade.Buyer]
	if !ok {
		e.positions[trade.Buyer] = &MarketPosition{
			margins: map[string]uint64{},
			partyID: trade.Buyer,
		}
		buyer = e.positions[trade.Buyer]
	}

	seller, ok := e.positions[trade.Seller]
	if !ok {
		e.positions[trade.Seller] = &MarketPosition{
			margins: map[string]uint64{},
			partyID: trade.Seller,
		}
		seller = e.positions[trade.Seller]
	}

	// Buyer INCREASED position size buy trade.Size
	buyer.size += int64(trade.Size)

	// Seller DECREASED position size buy trade.Size
	seller.size -= int64(trade.Size)

	if e.LogPositionUpdate {
		e.log.Info("Positions Updated for trade",
			logging.Trade(*trade),
			logging.String("buyer-position", fmt.Sprintf("%+v", buyer)),
			logging.String("seller-position", fmt.Sprintf("%+v", seller)))
	}
	// we've set all the values now, unlock after logging
	// because we're working on MarketPosition pointers
	e.mu.Unlock()
}

func (e *Engine) Positions() []MarketPosition {
	e.mu.RLock()
	out := make([]MarketPosition, 0, len(e.positions))
	for _, value := range e.positions {
		out = append(out, *value)
	}
	e.mu.RUnlock()
	return out
}