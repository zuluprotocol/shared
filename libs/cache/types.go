package cache

import (
	"fmt"

	"code.vegaprotocol.io/shared/libs/num"
	"code.vegaprotocol.io/vega/protos/vega"
)

type MarketData struct {
	staticMidPrice *num.Uint
	markPrice      *num.Uint
	targetStake    *num.Uint
	suppliedStake  *num.Uint
	openVolume     int64
	tradingMode    vega.Market_TradingMode
}

func SetMarketData(m *MarketData) func(md *MarketData) {
	return func(md *MarketData) {
		md.staticMidPrice = m.staticMidPrice.Clone()
		md.markPrice = m.markPrice.Clone()
		md.targetStake = m.targetStake.Clone()
		md.suppliedStake = m.suppliedStake.Clone()
		md.tradingMode = m.tradingMode
	}
}

func SetStaticMidPrice(staticMidPrice *num.Uint) func(md *MarketData) {
	return func(md *MarketData) {
		md.staticMidPrice = staticMidPrice.Clone()
	}
}

func SetMarkPrice(markPrice *num.Uint) func(md *MarketData) {
	return func(md *MarketData) {
		md.markPrice = markPrice.Clone()
	}
}

func SetTargetStake(targetStake *num.Uint) func(md *MarketData) {
	return func(md *MarketData) {
		md.targetStake = targetStake.Clone()
	}
}

func SetSuppliedStake(suppliedStake *num.Uint) func(md *MarketData) {
	return func(md *MarketData) {
		md.suppliedStake = suppliedStake.Clone()
	}
}

func SetTradingMode(tradingMode vega.Market_TradingMode) func(md *MarketData) {
	return func(md *MarketData) {
		md.tradingMode = tradingMode
	}
}

func SetOpenVolume(openVolume int64) func(md *MarketData) {
	return func(md *MarketData) {
		md.openVolume = openVolume
	}
}

func (md MarketData) StaticMidPrice() *num.Uint {
	return md.staticMidPrice.Clone()
}

func (md MarketData) MarkPrice() *num.Uint {
	return md.markPrice.Clone()
}

func (md MarketData) TargetStake() *num.Uint {
	return md.targetStake.Clone()
}

func (md MarketData) SuppliedStake() *num.Uint {
	return md.suppliedStake.Clone()
}

func (md MarketData) TradingMode() vega.Market_TradingMode {
	return md.tradingMode
}

func (md MarketData) OpenVolume() int64 {
	return md.openVolume
}

func FromVegaMD(marketData *vega.MarketData) (*MarketData, error) {
	staticMidPrice, err := num.ConvertUint256(marketData.StaticMidPrice)
	if err != nil {
		return nil, fmt.Errorf("invalid static mid price: %s", err)
	}

	markPrice, err := num.ConvertUint256(marketData.MarkPrice)
	if err != nil {
		return nil, fmt.Errorf("invalid mark price: %s", err)
	}

	targetStake, err := num.ConvertUint256(marketData.TargetStake)
	if err != nil {
		return nil, fmt.Errorf("invalid target stake: %s", err)
	}

	suppliedStake, err := num.ConvertUint256(marketData.SuppliedStake)
	if err != nil {
		return nil, fmt.Errorf("invalid supplied stake: %s", err)
	}

	return &MarketData{
		staticMidPrice: staticMidPrice,
		markPrice:      markPrice,
		targetStake:    targetStake,
		suppliedStake:  suppliedStake,
		tradingMode:    marketData.MarketTradingMode,
	}, nil
}

type Balance struct {
	general num.Uint
	margin  num.Uint
	bond    num.Uint
}

func GeneralAndBond(b Balance) *num.Uint {
	return num.Sum(&b.general, &b.bond)
}

func General(b Balance) *num.Uint {
	return &b.general
}

func Margin(b Balance) *num.Uint {
	return &b.margin
}

func Bond(b Balance) *num.Uint {
	return &b.bond
}

func SetBalanceByType(typ vega.AccountType, balance *num.Uint) func(*Balance) {
	switch typ {
	case vega.AccountType_ACCOUNT_TYPE_GENERAL:
		return SetGeneral(balance)
	case vega.AccountType_ACCOUNT_TYPE_MARGIN:
		return SetMargin(balance)
	case vega.AccountType_ACCOUNT_TYPE_BOND:
		return SetBond(balance)
	}
	return func(*Balance) {}
}

func SetGeneral(general *num.Uint) func(*Balance) {
	return func(b *Balance) {
		b.general = fromPtr(general)
	}
}

func SetMargin(margin *num.Uint) func(*Balance) {
	return func(b *Balance) {
		b.margin = fromPtr(margin)
	}
}

func SetBond(bond *num.Uint) func(*Balance) {
	return func(b *Balance) {
		b.bond = fromPtr(bond)
	}
}

// nolint:nonamedreturns
func fromPtr[T any](ptr *T) (t T) {
	if ptr == nil {
		return
	}
	return *ptr
}