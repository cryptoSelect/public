package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(host, user, passwd, dbname string, port int) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		host,
		user,
		passwd,
		dbname,
		port,
		"disable",
	)

	// 配置 GORM 日志
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 显示迁移日志
	}

	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	log.Println("Database connected successfully")
}

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(models ...interface{}) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	log.Println("Starting database migration...")

	for _, model := range models {
		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("Failed to migrate %T: %v", model, err)
			return fmt.Errorf("migration failed for %T: %w", model, err)
		}
		log.Printf("Successfully migrated %T", model)
	}

	log.Println("Database migration completed successfully")
	return nil
}

// DropTableIfExists 如果表存在则删除（谨慎使用）
func DropTableIfExists(model interface{}) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	if err := DB.Migrator().DropTable(model); err != nil {
		return fmt.Errorf("failed to drop table %T: %w", model, err)
	}

	log.Printf("Table %T dropped successfully", model)
	return nil
}

// CreateTable 创建表（如果不存在）
func CreateTable(model interface{}) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	if err := DB.Migrator().CreateTable(model); err != nil {
		return fmt.Errorf("failed to create table %T: %w", model, err)
	}

	log.Printf("Table %T created successfully", model)
	return nil
}

// HasTable 检查表是否存在
func HasTable(model interface{}) bool {
	if DB == nil {
		return false
	}

	return DB.Migrator().HasTable(model)
}
