package signalStra

import (
	"fmt"
	"github.com/rz1998/invest-basic/types/investBasic"
)

type SSignalShort struct {
	Date       string `json:"date"`
	UniqueCode string `json:"uniqueCode"`
	IsShort    bool   `json:"isShort"`
	IsLong     bool   `json:"isLong"`
}

type SSignalStra struct {
	// 策略名（必填）
	NameStra string `json:"nameStra"`
	// 信号产生时间（选填）
	TimeSignalCreated int64 `json:"timeSignalCreated"`
	// 信号建议下单时间（选填）
	TimeSignalSuggested int64 `json:"timeSignalSuggested"`
	// 信号对应证券（必填）
	UniqueCode string `json:"uniqueCode"`
	// 交易方向 1买；2卖（必填）
	TradeDir investBasic.EDirTrade `json:"tradeDir"`
	// 信号类型（选填）
	TypeSignal int `json:"typeSignal"`
	// 仓位 比如：半仓0.5，全仓1（选填，不填时平均分配）
	Priority float64 `json:"priority"`
	// 下单价格（必填）
	Price int64 `json:"price"`
	// 下单数量（选填）
	Vol int64 `json:"vol"`
}

func (signal *SSignalStra) String() string {
	return fmt.Sprintf("%s,%d,%d,%s,%d,%d,%f,%d,%d",
		signal.NameStra, signal.TimeSignalCreated, signal.TimeSignalSuggested,
		signal.UniqueCode, signal.TradeDir, signal.TypeSignal, signal.Priority, signal.Price, signal.Vol)
}

func (signal *SSignalStra) Title() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s",
		"nameStra", "timeSignalCreated", "timeSignalSuggested",
		"uniqueCode", "tradeDir", "typeSignal", "priority", "price", "vol")
}
