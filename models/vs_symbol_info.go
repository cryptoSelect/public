package models

type VsCoinInfo struct {
	ID        uint    `gorm:"primaryKey;comment:主键ID" json:"id"`
	VSTokenID int64   `json:"vsTokenId" gorm:"uniqueIndex;comment:ValueScan Token ID"`
	Name      string  `json:"name" gorm:"comment:币种名称"`
	Symbol    string  `json:"symbol" gorm:"index;comment:币种符号"`
	MarketCap float64 `json:"marketCap,string" gorm:"type:double precision;comment:市值"`
}

func (VsCoinInfo) TableName() string {
	return "vs_coin_info"
}
