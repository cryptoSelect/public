package main

import (
	"log"

	"github.com/cryptoSelect/public/database"
	"github.com/cryptoSelect/public/models"
)

func main() {
	// 初始化数据库连接
	database.InitDB("localhost", "postgres", "password", "crypto_db", 5432)

	// 方法1: 自动迁移（推荐用于生产环境）
	// 会自动创建表，如果表已存在则只更新结构
	err := database.AutoMigrate(&models.SymbolRecord{}, &models.UserInfo{}, &models.Subscription{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// 方法2: 检查表是否存在
	if database.HasTable(&models.SymbolRecord{}) {
		log.Println("SymbolRecord table already exists")
	} else {
		log.Println("SymbolRecord table does not exist")
	}

	// 方法3: 强制创建表（如果表已存在会报错）
	// err = database.CreateTable(&models.SymbolRecord{})
	// if err != nil {
	// 	log.Printf("Create table failed: %v", err)
	// }

	// 方法4: 删除表（谨慎使用）
	// err = database.DropTableIfExists(&models.SymbolRecord{})
	// if err != nil {
	// 	log.Printf("Drop table failed: %v", err)
	// }

	log.Println("Database operations completed successfully")
}
