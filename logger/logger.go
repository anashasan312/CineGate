package log

import (
	logger "github.com/rs/zerolog"
)

func Init(level string) {
	var logLevel logger.Level

	switch level {
	case "trace":
		logLevel = logger.TraceLevel
	case "debug":
		logLevel = logger.DebugLevel
	case "info":
		logLevel = logger.InfoLevel
	case "warn":
		logLevel = logger.WarnLevel
	case "error":
		logLevel = logger.ErrorLevel
	case "fatal":
		logLevel = logger.FatalLevel
	case "panic":
		logLevel = logger.PanicLevel
	case "disabled":
		logLevel = logger.Disabled
	default:
		logLevel = logger.InfoLevel

	}

	logger.SetGlobalLevel(logLevel)
}
