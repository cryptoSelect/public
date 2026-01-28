package models

import "time"

type SymbolRecord struct {
	ID              uint      `gorm:"primaryKey;comment:主键ID"`                                          // 主键ID
	Symbol          string    `gorm:"index:idx_symbol_cycle,unique;comment:交易对 (e.g. BTCUSDT)"`         // 交易对 (e.g. BTCUSDT)
	Cycle           string    `gorm:"index:idx_symbol_cycle,unique;comment:周期 (e.g. 5m, 1h)"`           // 周期 (e.g. 5m, 1h)
	Price           float64   `json:"price" gorm:"comment:当前价格"`                                        // 当前价格
	Volume          float64   `json:"volume" gorm:"comment:成交量"`                                        // 成交量
	TakerBuyVolume  float64   `json:"taker_buy_volume" gorm:"comment:主动买入量"`                            // 主动买入量
	TakerBuyRatio   float64   `json:"taker_buy_ratio" gorm:"comment:主动买入占比"`                            // 主动买入占比
	Rsi             float64   `json:"rsi" gorm:"comment:RSI值"`                                          // RSI值
	Rate            float64   `json:"rate" gorm:"comment:资金费率"`                                         // 资金费率
	RateCycle       int       `json:"rate_cycle" gorm:"comment:费率结算周期(小时)"`                             // 费率结算周期
	CrossType       int       `json:"cross_type" gorm:"comment:MACD交叉类型(1表示金叉0轴上2金叉0轴下，3死叉0轴上，4死叉0轴下)"` // MACD交叉类型 (金叉/死叉)
	CrossTime       time.Time `json:"cross_time" gorm:"comment:交叉时间"`                                   // 交叉时间
	Shape           int       `json:"shape" gorm:"comment:缠论分型 (1表示顶分型，2表示底分型)"`                        // 缠论分型 (顶分型/底分型)
	VpSignal        string    `json:"vp_signal" gorm:"comment:量价分析信号"`                                  // 量价分析信号
	Change          float64   `json:"change" gorm:"comment:涨跌幅"`                                        // 涨跌幅
	Description     string    `json:"description" gorm:"type:text;comment:详情描述"`                        // 详情描述
	NextFundingTime int64     `json:"next_funding_time" gorm:"comment:下次结算时间"`                          // 下次结算时间
	UpdatedAt       time.Time `json:"updated_at" gorm:"comment:更新时间"`                                   // 更新时间
	Support         float64   `json:"support" gorm:"comment:支撑位"`                                       // 支撑位
	Resistance      float64   `json:"resistance" gorm:"comment:压力位"`                                    // 压力位
	SMCSignal       string    `json:"smc_signal" gorm:"comment:SMC信号 (BOS/CHoCH)"`                      // SMC信号
	Fvg             string    `json:"fvg" gorm:"comment:FVG缺口"`                                         // FVG缺口
	Ob              string    `json:"ob" gorm:"comment:订单块 OB"`                                         // 订单块 OB
}

func (SymbolRecord) TableName() string {
	return "symbol_records"
}
