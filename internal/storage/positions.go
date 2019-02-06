package storage

import (
	"vega/msg"
)

type MarketBucket struct {
	Buys                []*msg.Trade
	Sells               []*msg.Trade
	BuyVolume           int64
	SellVolume          int64
	MinimumContractSize int64
}

func (ts *badgerTradeStore) GetTradesBySideBuckets(party string) map[string]*MarketBucket {

	marketBuckets := make(map[string]*MarketBucket, 0)
	tradesByTimestamp, err := ts.GetByParty(party, nil)

	if err != nil {
		return marketBuckets
	}

	if ts.LogPositionStoreDebug {
		ts.log.Debugf("Total trades by timestamp for party %s = %d", party, len(tradesByTimestamp))
	}

	for idx, trade := range tradesByTimestamp {
		if _, ok := marketBuckets[trade.Market]; !ok {
			marketBuckets[trade.Market] = &MarketBucket{[]*msg.Trade{}, []*msg.Trade{}, 0, 0, 1}
		}
		if trade.Buyer == party {
			marketBuckets[trade.Market].Buys = append(marketBuckets[trade.Market].Buys, tradesByTimestamp[idx])
			marketBuckets[trade.Market].BuyVolume += int64(tradesByTimestamp[idx].Size)
		}
		if trade.Seller == party {
			marketBuckets[trade.Market].Sells = append(marketBuckets[trade.Market].Sells, tradesByTimestamp[idx])
			marketBuckets[trade.Market].SellVolume += int64(tradesByTimestamp[idx].Size)
		}
	}

	return marketBuckets
}