# Crypto Public

> Crypto 公共组件库 - 提供数据库模型、数据库连接和通用工具

## 📋 项目简介

Crypto Public 是一个 Go 语言的公共组件库，为加密货币相关项目提供通用的数据模型、数据库连接和基础工具。该项目专注于提供标准化的数据结构和数据库操作，便于多个项目复用。

## 🚀 功能特性

- **标准化数据模型**：提供币种信息、交易记录、资金流向等核心数据模型
- **数据库支持**：基于 GORM 的 PostgreSQL 数据库连接和操作
- **自动迁移**：支持数据库表结构的自动迁移和版本管理
- **完整注释**：所有字段都有详细的中文注释，便于理解和维护

## 📁 项目结构

```
public/
├── config/          # 配置文件目录
├── database/        # 数据库连接和操作
├── models/          # 数据模型定义
│   ├── symbol.go           # 交易对记录模型
│   ├── trade_inflow.go     # 资金流向模型
│   └── vs_symbol_info.go   # ValueScan 币种信息模型
├── main.go          # 示例程序
├── go.mod           # Go 模块定义
└── README.md        # 项目文档
```

## 🗄️ 数据模型

### 1. SymbolRecord - 交易对记录

存储交易对的技术指标和交易数据：

```go
type SymbolRecord struct {
    ID              uint      // 主键ID
    Symbol          string    // 交易对 (e.g. BTCUSDT)
    Cycle           string    // 周期 (e.g. 5m, 1h)
    Price           float64   // 当前价格
    Volume          float64   // 成交量
    TakerBuyVolume  float64   // 主动买入量
    TakerBuyRatio   float64   // 主动买入占比
    Rsi             float64   // RSI值
    Rate            float64   // 资金费率
    // ... 更多字段
}
```

**表名**: `symbol_records`

### 2. CoinTradeInflowDto - 资金流向

存储现货和合约的资金流入流出数据：

```go
type CoinTradeInflowDto struct {
    ID                        uint    // 主键ID
    Symbol                    string  // 币种符号
    TimeParticleEnum          int     // 时间粒度枚举
    Time                      string  // 时间粒度标识(如 5m/15m/1h)
    StopTradeInflow           float64 // 现货资金流入(净值)
    ContractTradeInflow       float64 // 合约资金流入(净值)
    // ... 更多资金流向字段
}
```

**表名**: `trade_inflow`

### 3. VsCoinInfo - ValueScan 币种信息

存储 ValueScan 平台的币种基础信息：

```go
type VsCoinInfo struct {
    ID        uint    // 主键ID
    VSTokenID int64   // ValueScan Token ID (唯一)
    Name      string  // 币种名称
    Symbol    string  // 币种符号
    MarketCap float64 // 市值
}
```

**表名**: `vs_coin_info`

## 🛠️ 安装使用

### 环境要求

- Go 1.25.4+
- PostgreSQL 12+

### 安装依赖

```bash
go mod tidy
```

### 数据库配置

```go
import "github.com/your-org/crypto-public/database"

// 初始化数据库连接
database.InitDB("localhost", "postgres", "password", "crypto_db", 5432)
```

### 自动迁移

```go
import "github.com/your-org/crypto-public/models"

// 自动迁移所有表
err := database.AutoMigrate(
    &models.SymbolRecord{},
    &models.CoinTradeInflowDto{},
    &models.VsCoinInfo{},
)
if err != nil {
    log.Fatalf("Migration failed: %v", err)
}
```

## 📖 使用示例

### 基础使用

```go
package main

import (
    "log"
    "github.com/your-org/crypto-public/database"
    "github.com/your-org/crypto-public/models"
)

func main() {
    // 初始化数据库
    database.InitDB("localhost", "postgres", "password", "crypto_db", 5432)
    
    // 自动迁移
    err := database.AutoMigrate(&models.SymbolRecord{})
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }
    
    // 创建记录
    record := models.SymbolRecord{
        Symbol: "BTCUSDT",
        Cycle:  "5m",
        Price:  45000.0,
        Volume: 1000.0,
    }
    
    database.DB.Create(&record)
    log.Println("Record created successfully")
}
```

### 查询示例

```go
// 查询特定交易对的最新记录
var record models.SymbolRecord
database.DB.Where("symbol = ? AND cycle = ?", "BTCUSDT", "5m").
    Order("updated_at DESC").
    First(&record)

// 分页查询资金流向数据
var inflows []models.CoinTradeInflowDto
database.DB.Where("symbol = ?", "BTC").
    Limit(100).
    Offset(0).
    Find(&inflows)
```

## 🔧 数据库操作

### 支持的操作

- `AutoMigrate()` - 自动迁移表结构
- `HasTable()` - 检查表是否存在
- `CreateTable()` - 创建表
- `DropTableIfExists()` - 删除表

### 连接配置

```go
// 数据库连接字符串格式
dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai"
```

## 📊 索引设计

### SymbolRecord 表索引

- `idx_symbol_cycle` - (symbol, cycle) 复合唯一索引

### CoinTradeInflowDto 表索引

- `symbol` - 币种符号索引

### VsCoinInfo 表索引

- `vs_token_id` - 唯一索引
- `symbol` - 币种符号索引

## 🤝 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📝 版本历史

- **v1.0.0** - 初始版本
  - 基础数据模型
  - 数据库连接和迁移功能
  - 完整的项目文档

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 提交 Issue
- 发起 Discussion
- 邮件联系

---

⭐ 如果这个项目对你有帮助，请给个 Star！
