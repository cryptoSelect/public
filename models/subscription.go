package models

import "time"

// Subscription 订阅表（public 库）
type Subscription struct {
	ID        uint      `gorm:"primaryKey;comment:主键ID"`
	Symbol    string    `gorm:"index:idx_sub_user_symbol;not null;comment:交易对(如 BTCUSDT)"`
	UserID    uint      `gorm:"index:idx_sub_user_symbol;not null;comment:用户ID"`
	CreatedAt time.Time `gorm:"comment:创建时间"`
}

func (Subscription) TableName() string {
	return "subscription"
}
