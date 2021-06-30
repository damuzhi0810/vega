package steps

import (
	"fmt"
	"strings"
	"time"

	"code.vegaprotocol.io/vega/execution"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/cucumber/godog/gherkin"
)

type MappedMD struct {
	md     types.MarketData
	u64Map map[string]*uint64
	strMap map[string]*string
	tMap   map[string]*int64
	i64Map map[string]*int64
	tm     *types.Market_TradingMode
	tr     *types.AuctionTrigger
	et     *types.AuctionTrigger
}

type ErrStack []error

func TheMarketDataShouldBe(engine *execution.Engine, mID string, data *gherkin.DataTable) error {
	actual, err := engine.GetMarketData(mID)
	if err != nil {
		return err
	}
	// create a copy (deep copy), override the values we've gotten with those from the table so we can compare the objects
	expect := mappedMD(actual)
	// special fields first, these need to be compared manually
	u64Set := expect.parseU64(data)
	i64Set := expect.parseI64(data)
	tSet := expect.parseTimes(data)
	strSet := expect.parseStr(data)
	expect.parseSpecial(data)
	if pm := getPriceBounds(data); len(pm) > 0 {
		expect.md.PriceMonitoringBounds = pm
	}
	// this might be a sparse check
	if lp := getLPFeeShare(data); len(lp) > 0 {
		expect.md.LiquidityProviderFeeShare = lp
	}
	cmp := mappedMD(actual)
	parsed := mappedMD(expect.md)
	errs := make([]error, 0, len(u64Set)+len(i64Set)+len(strSet)+2)
	if expect.tm != nil && *expect.tm != expect.md.MarketTradingMode {
		errs = append(errs, fmt.Errorf("expected '%s' trading mode, instead got '%s'", *expect.tm, expect.md.MarketTradingMode))
	}
	if expect.tr != nil && *expect.tr != expect.md.Trigger {
		errs = append(errs, fmt.Errorf("expected '%s' auction trigger, instead got '%s'", *expect.tr, expect.md.Trigger))
	}
	if expect.et != nil && *expect.et != expect.md.ExtensionTrigger {
		errs = append(errs, fmt.Errorf("expected '%s' extension trigger, instead got '%s'", *expect.et, expect.md.ExtensionTrigger))
	}
	// compare uint64
	for _, u := range u64Set {
		e, g := cmp.u64Map[u], parsed.u64Map[u] // get pointers to both fields
		if *e != *g {
			errs = append(errs, fmt.Errorf("expected '%d' for %s, instead got '%d'", *e, u, *g))
		}
	}
	// compare int64
	for _, i := range i64Set {
		e, g := cmp.i64Map[i], parsed.i64Map[i]
		if *e != *g {
			errs = append(errs, fmt.Errorf("expected '%d' for %s, instead got '%d'", *e, i, *g))
		}
	}
	// compare times, which is basically identical to comparing i64
	for _, i := range tSet {
		e, g := cmp.tMap[i], parsed.tMap[i]
		if *e != *g {
			errs = append(errs, fmt.Errorf("expected '%d' for %s, instead got '%d'", *e, i, *g))
		}
	}
	// compare strings
	for _, s := range strSet {
		e, g := cmp.strMap[s], parsed.strMap[s]
		if *e != *g {
			errs = append(errs, fmt.Errorf("expected '%s' for %s, instead got '%s'", *e, s, *g))
		}
	}
	if err := cmpPriceBounds(expect, actual); len(err) > 0 {
		errs = append(errs, err...)
	}
	if err := cmpLPFeeShare(expect, actual); len(err) > 0 {
		errs = append(errs, err...)
	}
	// wrap all errors in a single error type for complete information
	if len(errs) > 0 {
		return ErrStack(errs)
	}
	// compare special fields (trading mode and auction trigger)
	return nil
}

func cmpLPFeeShare(expect *MappedMD, got types.MarketData) []error {
	errs := make([]error, 0, len(expect.md.LiquidityProviderFeeShare))
	for _, lpfs := range expect.md.LiquidityProviderFeeShare {
		match := false
		var found *types.LiquidityProviderFeeShare
		for _, g := range got.LiquidityProviderFeeShare {
			if lpfs.Party == g.Party {
				found = g
				match = lpfs.AverageEntryValuation == g.AverageEntryValuation && lpfs.EquityLikeShare == g.EquityLikeShare
				break
			}
		}
		if !match {
			if found == nil {
				errs = append(errs, fmt.Errorf("no LP fee share found for party %s", lpfs.Party))
			} else {
				errs = append(errs, fmt.Errorf(
					"expected LP fee share for party %s with avg valuation %s and equity like share %s, instead got avg. valuation %s and equity %s",
					lpfs.Party,
					lpfs.AverageEntryValuation,
					lpfs.EquityLikeShare,
					found.AverageEntryValuation,
					found.EquityLikeShare,
				))
			}
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

func cmpPriceBounds(expect *MappedMD, got types.MarketData) []error {
	errs := make([]error, 0, len(expect.md.PriceMonitoringBounds))
	for _, pmb := range expect.md.PriceMonitoringBounds {
		var bounds *types.PriceMonitoringBounds
		match := false
		for _, g := range got.PriceMonitoringBounds {
			if g.Trigger.Horizon == pmb.Trigger.Horizon {
				bounds = g
				match = pmb.MaxValidPrice == g.MaxValidPrice && pmb.MinValidPrice == g.MinValidPrice
				break
			}
		}
		if !match {
			if bounds == nil {
				errs = append(errs, fmt.Errorf("no price bound for horizon %d found", pmb.Trigger.Horizon))
			} else {
				errs = append(errs, fmt.Errorf(
					"expected price bounds %d-%d for horizon %d, instead got %d-%d",
					pmb.MinValidPrice,
					pmb.MaxValidPrice,
					pmb.Trigger.Horizon,
					bounds.MinValidPrice,
					bounds.MaxValidPrice,
				))
			}
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

func getPriceBounds(data *gherkin.DataTable) (ret []*types.PriceMonitoringBounds) {
	for _, row := range ParseTable(data) {
		h, ok := row.I64B("horizon")
		if !ok {
			return nil
		}
		expected := &types.PriceMonitoringBounds{
			MinValidPrice: row.MustU64("min bound"),
			MaxValidPrice: row.MustU64("max bound"),
			Trigger: &types.PriceMonitoringTrigger{
				Horizon: h,
			},
		}
		ret = append(ret, expected)
	}
	return ret
}

func getLPFeeShare(data *gherkin.DataTable) (ret []*types.LiquidityProviderFeeShare) {
	for _, r := range ParseTable(data) {
		avg, ok := r.StrB("average entry valuation")
		if !ok {
			return nil
		}
		ret = append(ret, &types.LiquidityProviderFeeShare{
			Party:                 r.MustStr("party"),
			EquityLikeShare:       r.MustStr("equity share"),
			AverageEntryValuation: avg,
		})
	}
	return ret
}

func (m *MappedMD) parseSpecial(data *gherkin.DataTable) {
	todo := map[string]struct{}{
		"trading mode":      {},
		"auction trigger":   {},
		"extension trigger": {},
	}
	for _, r := range ParseTable(data) {
		for k := range todo {
			if _, ok := r.StrB(k); ok {
				switch k {
				case "trading mode":
					tm := r.MustTradingMode(k)
					m.tm = &tm
				case "auction trigger":
					at := r.MustAuctionTrigger(k)
					m.tr = &at
				case "extension trigger":
					et := r.MustAuctionTrigger(k)
					m.et = &et
				}
				delete(todo, k)
			}
		}
		if len(todo) == 0 {
			return
		}
	}
}

// parses the data, and returns a slice of keys for the values that were provided
func (m *MappedMD) parseU64(data *gherkin.DataTable) []string {
	set := make([]string, 0, len(m.u64Map))
	for _, r := range ParseTable(data) {
		for k, ptr := range m.u64Map {
			if u, ok := r.U64B(k); ok {
				*ptr = u
				set = append(set, k)
				// avoid reassignments in following iterations
				delete(m.u64Map, k)
			}
		}
	}
	return set
}

func (m *MappedMD) parseTimes(data *gherkin.DataTable) []string {
	// already set start based off of the value in the map
	// does some trickery WRT auction end time, so we can check if the auction duration is N seconds
	end, start := int64(0), *m.tMap["auction start"]
	set := make([]string, 0, len(m.tMap))
	for _, r := range ParseTable(data) {
		for k, ptr := range m.tMap {
			if i, ok := r.I64B(k); ok {
				if k == "auction end" {
					end = i
					if end < start {
						i = start + int64(time.Duration(end)*time.Second)
					}
				}
				*ptr = i
				set = append(set, k)
				// again: avoid reassignments when parsing the next row
				delete(m.i64Map, k)
			}
		}
	}
	return set
}

func (m *MappedMD) parseI64(data *gherkin.DataTable) []string {
	set := make([]string, 0, len(m.i64Map))
	for _, r := range ParseTable(data) {
		for k, ptr := range m.i64Map {
			if i, ok := r.I64B(k); ok {
				*ptr = i
				set = append(set, k)
				// again: avoid reassignments when parsing the next row
				delete(m.i64Map, k)
			}
		}
	}
	return set
}

func (m *MappedMD) parseStr(data *gherkin.DataTable) []string {
	set := make([]string, 0, len(m.strMap))
	for _, r := range ParseTable(data) {
		for k, ptr := range m.strMap {
			if i, ok := r.StrB(k); ok {
				*ptr = i
				set = append(set, k)
				// again: avoid reassignments when parsing the next row
				delete(m.strMap, k)
			}
		}
	}
	return set
}

func mappedMD(md types.MarketData) *MappedMD {
	r := &MappedMD{
		md: md,
	}
	r.u64Map = map[string]*uint64{
		"mark price":               &r.md.MarkPrice,
		"best bid price":           &r.md.BestBidPrice,
		"best bid volume":          &r.md.BestBidVolume,
		"best offer price":         &r.md.BestOfferPrice,
		"best offer volume":        &r.md.BestOfferVolume,
		"best static bid price":    &r.md.BestStaticBidPrice,
		"best static bid volume":   &r.md.BestStaticBidVolume,
		"best static offer price":  &r.md.BestStaticOfferPrice,
		"best static offer volume": &r.md.BestStaticOfferVolume,
		"mid price":                &r.md.MidPrice,
		"static mid price":         &r.md.StaticMidPrice,
		"open interest":            &r.md.OpenInterest,
		"indicative price":         &r.md.IndicativePrice,
		"indicative volume":        &r.md.IndicativeVolume,
	}
	r.strMap = map[string]*string{
		"target stake":       &r.md.TargetStake,
		"supplied stake":     &r.md.SuppliedStake,
		"market value proxy": &r.md.MarketValueProxy,
		"market":             &r.md.Market, // this is a bit pointless, but might as well add it
	}
	r.tMap = map[string]*int64{
		"timestamp":     &r.md.Timestamp,
		"auction end":   &r.md.AuctionEnd,
		"auction start": &r.md.AuctionStart,
	}
	return r
}

// Error so we print out the wrong matches line by line
func (e ErrStack) Error() string {
	str := make([]string, 0, len(e))
	for _, v := range e {
		str = append(str, v.Error())
	}
	return strings.Join(str, "\n")
}