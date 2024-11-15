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
