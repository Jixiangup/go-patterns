package method

func Usage() {
	logger, _ := CreateLoggerFactory(Console)
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
}
