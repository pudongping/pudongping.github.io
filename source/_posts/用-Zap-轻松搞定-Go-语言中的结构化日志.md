---
title: 用 Zap 轻松搞定 Go 语言中的结构化日志
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Go
tags:
  - Go
  - Golang
  - Zap
abbrlink: b16382e2
date: 2024-11-13 17:29:13
img:
coverImg:
password:
summary:
---


在开发现代应用程序时，日志记录是一个不可或缺的部分。它不仅能帮助我们跟踪程序的运行状态，还能在出现问题时提供宝贵的调试信息。

在 Go 语言中，有许多日志库可供选择，但在性能和灵活性方面，Zap 是其中的佼佼者。

今天，我将带你深入了解如何在 Go 项目中使用 Zap 进行结构化日志记录，并且展示如何定制日志输出，以满足生产环境的需求。

## 为什么选择 Zap？

Zap 是 Uber 开发的一款高性能日志库，专为那些需要快速、结构化日志记录的场景而设计。与其他日志库相比，Zap 的性能更为优越，尤其是在需要频繁记录日志的高并发环境中。

此外，Zap 提供了两种日志记录接口：**Logger** 和 **SugaredLogger**。

- **Logger** 提供了最基础的、类型安全的结构化日志记录方式。虽然使用时稍显复杂，但在性能上无可匹敌，非常适合对性能有极致要求的场景。
- **SugaredLogger** 则在 Logger 之上进行了封装，提供了更为便捷的日志记录方法。尽管在性能上稍逊于 Logger，但它的灵活性和易用性使其成为大多数场景下的首选。

## 基础日志记录示例

为了更好地理解 Zap 的使用，让我们从一个简单的例子开始。

```go
package log_demo

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() {
	// 初始化 Logger，这里我们使用了开发环境的配置
	logger, _ = zap.NewDevelopment()
}

func ZapPrintLog() {
	InitLogger()
	defer logger.Sync() // 确保日志缓冲区中的所有日志都被写入

	// 记录一条信息级别的日志，并附带结构化的键值对
	logger.Info("This is a log message", zap.String("key1", "value1"), zap.Float64s("key2", []float64{1.0, 2.0, 3.0}))
}
```

在这个例子中，我们通过 `zap.NewDevelopment()` 初始化了一个 Logger，并使用 `logger.Info` 方法记录了一条信息级别的日志。`zap.String` 和 `zap.Float64s` 是 Zap 提供的用于结构化日志的字段构造器，它们将日志内容按键值对的形式记录下来。

## 更便捷的 SugaredLogger

虽然结构化日志非常有用，但有时我们只需要快速记录一些信息。此时，`SugaredLogger` 就派上了用场。它支持类似 `fmt.Printf` 的日志记录方式，使代码更加简洁。

```go
package log_demo

import (
	"go.uber.org/zap"
)

var sugaredLogger *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	sugaredLogger = logger.Sugar()
}

func ZapPrintLog1() {
	InitLogger()
	defer sugaredLogger.Sync()

	// 使用 SugaredLogger 记录日志
	sugaredLogger.Infof("This is a formatted log: %s", "example")
	sugaredLogger.Infow("This is a log message", "key1", "value1", "key2", []float64{1.0, 2.0, 3.0})
}
```

在上面的代码中，`SugaredLogger` 提供了 `Infof` 和 `Infow` 方法，前者允许你像使用 `fmt.Printf` 一样格式化日志，后者则结合了结构化日志的优点，使日志记录更加灵活。

## 自定义日志配置

在实际生产环境中，我们可能需要对日志输出进行更精细的控制，比如将日志输出到文件、控制台，或者对日志进行按大小或时间切割。

**Zap 自身不支持日志切割**，但我们可以借助第三方库 [lumberjack](https://github.com/natefinch/lumberjack) 来实现这一功能。

```go
package log_demo

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitCustomLogger 初始化一个定制的 Logger
func InitCustomLogger() {
	// 创建一个日志切割器
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "zap.log",    // 日志文件路径
		MaxSize:    10,           // 单个日志文件最大尺寸（MB）
		MaxBackups: 5,            // 最多保留5个备份
		MaxAge:     30,           // 日志保留最长天数
		Compress:   true,         // 启用日志压缩
	})

	// 配置日志编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder  // 时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志级别大写
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 创建 Logger Core
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// zap.AddCaller() 会在日志中加入调用函数的文件名和行号
	// zap.AddCallerSkip(1) 会跳过调用函数的文件名和行号
	// 当我们不是直接使用初始化好的logger实例记录日志，而是将其包装成一个函数等，此时日录日志的函数调用链会增加，想要获得准确的调用信息就需要通过AddCallerSkip函数来跳过
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func ZapPrintLog2() {
	InitCustomLogger()
	defer logger.Sync()

	logger.Debug("This is a custom log message", zap.String("key1", "value1"), zap.Float64s("key2", []float64{1.0, 2.0, 3.0}))
}
```

在这个示例中，我们通过 `lumberjack.Logger` 实现了日志文件的自动切割与管理。`zapcore.NewConsoleEncoder` 配置了日志的编码格式，确保日志输出不仅有结构化的信息，还带有清晰的时间戳和日志级别标识。此外，我们使用了 `zap.AddCaller()` 和 `zap.AddCallerSkip(1)`，这两个函数可以在日志中添加调用函数的文件名和行号，帮助我们更快地定位日志来源。

## 应用场景示例：记录调试信息

假设我们在开发一个 web 应用时需要记录一些调试信息。此时，使用我们之前定义的 `InitCustomLogger` 函数可以非常方便地记录这些信息。

```go
func main() {
    ZapPrintLog2() // 调用定制的日志打印函数

    // 假设这里有其他业务逻辑
    for i := 0; i < 3; i++ {
        logger.Info("Processing iteration", zap.Int("iteration", i))
    }
}
```

在这个简短的示例中，`logger.Info` 会在每次循环中记录当前迭代次数，并将日志输出到指定的日志文件中。

## 结语

`Zap` 是一个功能强大且灵活的日志库，无论你是需要极致性能，还是希望日志记录更为简单直观，Zap 都能满足你的需求。

通过本文的讲解，你不仅了解了如何在 Go 中使用 Zap 进行结构化日志记录，还学习了如何定制日志输出，以应对实际生产环境中的需求。

掌握 Zap 的使用，将使你的 Go 项目在日志管理方面更上一层楼。如果你还没有尝试过 Zap，现在就是开始的好时机！