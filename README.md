# zapX - Uber Zap 日志库增强封装

[![Go Report Card](https://goreportcard.com/badge/github.com/huajianxiaowanzi/zapX)](https://goreportcard.com/report/github.com/huajianxiaowanzi/zapX)
[![Go 参考文档](https://pkg.go.dev/badge/github.com/huajianxiaowanzi/zapX.svg)](https://pkg.go.dev/github.com/huajianxiaowanzi/zapX)

zapX 是对 Uber Zap 高性能日志库的增强封装，提供更友好的 API 接口和开箱即用的最佳实践配置。

## ✨ 核心特性

- 🌈 **彩色终端输出** - 不同日志级别自动着色，关键信息一目了然
- 📅 **智能日志分割** - 按日期和级别自动分类存储，日志管理更轻松
- 🚀 **极简 API 设计** - 比原生 Zap 更直观易用的接口设计
- 🔧 **灵活配置** - 支持控制台/文件输出、日志级别等多维度配置
- 🔍 **上下文追踪** - 支持结构化日志字段，方便问题排查
- ⏱ **标准化时间** - ISO8601 标准时间格式，兼容各类日志分析工具

## 安装

```bash
go get github.com/huajianxiaowanzi/zapX/log@v0.1.0
```

## 🚀 快速开始
### 1. 基础使用

```go
package main

import "github.com/huajianxiaowanzi/zapX/log"

func main() {
    // 初始化日志（默认配置：控制台输出 + Info级别）
    log.Init()
    
    // 记录不同级别日志
    log.Debug("调试信息")      // 默认不显示
    log.Info("服务启动成功")   // 绿色输出
    log.Warn("磁盘空间不足")   // 黄色输出
    log.Error("连接失败")     // 红色输出
    
    // 格式化输出
    log.Infof("当前用户: %s", "张三")
    log.Errorf("错误码: %d", 500)
}
```

### 2. 添加上下文
```go
// 添加单个字段
log.With(zap.String("requestID", "req-123")).Info("收到请求")

// 添加多个字段
log.With(
zap.String("userID", "u1001"),
zap.Int("loginAttempt", 3),
).Warn("用户登录频繁失败")

// 错误日志专用
err := errors.New("connection timeout")
log.WithError(err).Error("数据库操作失败")
```

### 3. 自定义配置
```go
func main() {
    // 自定义配置
    cfg := log.LogConfig{
        ConsoleOutput: true,   // 启用控制台输出, 默认为true
        FileOutput:    true,    // 启用文件输出, 默认为false
        Level:         zapcore.DebugLevel, // 设置日志级别, 默认为zapcore.InfoLevel
    }
    
    log.Init(cfg)
    
    // 这些日志会同时输出到控制台和文件
    log.Debug("调试信息")  // 现在会显示
    log.Info("系统状态正常")
}
```

## 🗄️ 日志文件结构
日志文件按日期自动组织，示例结构：
```text
logs/
├── 2023-10-01/
│   ├── out.log    # 普通日志(Info/Warn)
│   └── err.log    # 错误日志(Error+)
├── 2023-10-02/
│   ├── out.log
│   └── err.log
└── current -> 2023-10-02  # 当日日志软链接
```
