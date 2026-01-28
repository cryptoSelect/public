package models

type CoinTradeInflowDto struct {
	ID                        uint    `gorm:"primaryKey;comment:主键ID" json:"id"`
	Symbol                    string  `json:"symbol" gorm:"index;comment:币种符号"`
	TimeParticleEnum          int     `json:"timeParticleEnum" gorm:"comment:时间粒度枚举"`
	Time                      string  `json:"time" gorm:"comment:时间粒度标识(如 5m/15m/1h)"`
	Stop                      bool    `json:"stop" gorm:"comment:是否统计现货(Stop)"`
	StopTradeInflow           float64 `json:"stopTradeInflow" gorm:"comment:现货资金流入(净值)"`
	StopTradeAmount           float64 `json:"stopTradeAmount" gorm:"comment:现货成交额"`
	StopTradeInflowChange     float64 `json:"stopTradeInflowChange" gorm:"comment:现货资金流入变化率"`
	StopTradeAmountChange     float64 `json:"stopTradeAmountChange" gorm:"comment:现货成交额变化率"`
	Contract                  bool    `json:"contract" gorm:"comment:是否统计合约(Contract)"`
	ContractTradeInflow       float64 `json:"contractTradeInflow" gorm:"comment:合约资金流入(净值)"`
	ContractTradeAmount       float64 `json:"contractTradeAmount" gorm:"comment:合约成交额"`
	ContractTradeInflowChange float64 `json:"contractTradeInflowChange" gorm:"comment:合约资金流入变化率"`
	ContractTradeAmountChange float64 `json:"contractTradeAmountChange" gorm:"comment:合约成交额变化率"`
	StopTradeIn               float64 `json:"stopTradeIn" gorm:"comment:现货流入"`
	StopTradeOut              float64 `json:"stopTradeOut" gorm:"comment:现货流出"`
	ContractTradeIn           float64 `json:"contractTradeIn" gorm:"comment:合约流入"`
	ContractTradeOut          float64 `json:"contractTradeOut" gorm:"comment:合约流出"`
}

func (CoinTradeInflowDto) TableName() string {
	return "trade_inflow"
}
