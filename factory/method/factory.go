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
