package models

import "time"

// UserInfo 用户信息表（public 库）
type UserInfo struct {
	ID         uint      `gorm:"primaryKey;comment:主键ID"`
	Email      string    `gorm:"uniqueIndex;not null;comment:邮箱"`
	Password   string    `gorm:"not null;comment:密码(bcrypt哈希)"`
	TelegramID string    `gorm:"column:telegram_id;comment:Telegram用户ID"`
	CreatedAt  time.Time `gorm:"comment:创建时间"`
	UpdatedAt  time.Time `gorm:"comment:更新时间"`
}

func (UserInfo) TableName() string {
	return "user_info"
}
