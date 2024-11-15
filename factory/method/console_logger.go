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
