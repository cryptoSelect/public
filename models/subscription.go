package models

import "time"

// Subscription 订阅表（public 库），按 symbol+cycle 订阅
type Subscription struct {
	ID        uint      `gorm:"primaryKey;comment:主键ID"`
	Symbol    string    `gorm:"uniqueIndex:idx_sub_user_symbol_cycle;not null;comment:交易对(如 BTCUSDT)"`
	Cycle     string    `gorm:"uniqueIndex:idx_sub_user_symbol_cycle;not null;comment:周期(如 15m, 1h)"`
	UserID    uint      `gorm:"uniqueIndex:idx_sub_user_symbol_cycle;not null;comment:用户ID"`
	CreatedAt time.Time `gorm:"comment:创建时间"`
}

func (Subscription) TableName() string {
	return "subscription"
}
