<h1>工厂方法模式</h1>

__工厂方法模式（Factory Method Pattern）__ 是一种创建型设计模式，它提供了一个接口，用于创建对象，但由子类决定要实例化的类是哪一个。通过这种方式，工厂方法模式让类的实例话推迟到子类中进行，从而实现了对象的创建与使用的解藕。

<h2>目录</h2>

<!-- TOC -->
  * [关键概念](#关键概念)
  * [适用场景](#适用场景)
  * [优缺点](#优缺点)
    * [优点](#优点)
    * [缺点](#缺点)
  * [举例](#举例)
  * [实现](#实现)
    * [接口定义](#接口定义)
    * [实例实现](#实例实现)
  * [使用](#使用)
<!-- TOC -->

## 关键概念

1. __工厂方法__: 通常这是一个定义在接口中的方法，用于创建对象。该方法由子类实现，用于返回具体的产品实例。
2. __产品接口__: 这是所有具体产品类实现的接口，它定义了产品的行为。
3. __具体产品__: 这是实现产品接口的具体产品实例，它定义了产品的具体行为。表示实际被创建的对象。
4. __具体工厂__: 这是实现工厂方法的具体工厂函数，它返回具体产品的实例。

## 适用场景

- __创建过程复杂的对象__: 当对象的创建过程需要多步或者比较复杂的逻辑时，可以使用工厂方法将对象的创建过程封装起来，避免客户端直接创建对象。
- __实例之间具有相似的构建过程，但实现不同__: 当需要创建一系列相关或相似的实例时，使用工厂方法可以通过工厂来统一管理不同产品的创建
- __调用方不需要关心具体产品类的实现__: 调用方不需要知道具体的类实现细节，只需要依赖工厂接口即可。

## 优缺点

### 优点

- __解藕__: 客户端代码不需要知道具体的产品类，实现了高内聚和低耦合，易于扩展和维护。
- __可扩展性__: 添加新的实例时，只需要添加新的具体实例，不需要修改原有代码和现有客户端代码(非侵入性)。
- __灵活性__: 可以通过配置工厂方法来动态控制创建哪些具体的实例，增加了代码的灵活性。

### 缺点

- __增加实例的数量__: 每个实例都需要一个对应的接口实现，可能会导致系统实例的数量增加。
- __不适用于实例种类较少的情况__: 如果实例的种类很少，使用工厂方法模式可能显得过于复杂，增加了不必要的代码。

## 举例

假设需要设计一个简单的日志系统，支持两种日志记录方式：控制台输出和文件输出。可以使用工厂方法模式来实现这个日志系统。

## 实现

### 接口定义

```go
package method

// OutputType 日志输出类型
type OutputType string

const (
	// Console 控制台输出
	Console OutputType = "console"
	// File 文件输出
	File OutputType = "file"
)

// Logger 日志接口
type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
}

// LoggerFactory 日志工厂接口 也可以定义为函数 因为Golang可以直接定义函数无需存在实例之中
type LoggerFactory interface {
	CreateLogger(outputType OutputType) (Logger, error)
}

// CreateLoggerFactory 获取日志实例工厂函数
func CreateLoggerFactory(outputType OutputType) (Logger, error) {
	var logger Logger
	switch outputType {
	case Console:
		logger = new(ConsoleLogger)
	case File:
		logger = new(FileLogger)
	default:
		return nil, nil
	}
	return logger, nil

}
```

### 实例实现

> [控制台日志](console_logger.go)实现

```go
package method

import "log"

// ConsoleLogger 控制台日志 实现自 Logger 接口
type ConsoleLogger struct {
}

func (c *ConsoleLogger) Debug(message string, args ...interface{}) {
	log.Printf("ConsoleLogger [debug] "+message, args...)
	println()
}

func (c *ConsoleLogger) Info(message string, args ...interface{}) {
	log.Printf("ConsoleLogger [info] "+message, args...)
	println()
}

func (c *ConsoleLogger) Warn(message string, args ...interface{}) {
	log.Printf("ConsoleLogger [warn] "+message, args...)
	println()
}

func (c *ConsoleLogger) Error(message string, args ...interface{}) {
	log.Printf("ConsoleLogger [error] "+message, args...)
	println()
}
```

> [文件日志](file_logger.go)实现

```go
package method

import "log"

// FileLogger 文件日志 实现自 Logger 接口
type FileLogger struct {
}

func (f *FileLogger) Debug(message string, args ...interface{}) {
	log.Printf("FileLogger [debug] "+message, args...)
	println()
}

func (f *FileLogger) Info(message string, args ...interface{}) {
	log.Printf("FileLogger [info] "+message, args...)
	println()
}

func (f *FileLogger) Warn(message string, args ...interface{}) {
	log.Printf("FileLogger [warn] "+message, args...)
	println()
}

func (f *FileLogger) Error(message string, args ...interface{}) {
	log.Printf("FileLogger [error] "+message, args...)
	println()
}
```

## 使用

> 见 [usage.go](usage.go)

```go
logger, _ := CreateLoggerFactory(Console)
logger.Debug("debug message")
logger.Info("info message")
logger.Warn("warn message")
logger.Error("error message")
```