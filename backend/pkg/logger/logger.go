package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New constructs a zap.Logger configured for JSON output.
func New(level string) (*zap.Logger, error) {
	zapCfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(parseLevel(level)),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return zapCfg.Build()
}

func parseLevel(lvl string) zapcore.Level {
	switch lvl {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// Sync flushes logger buffers, ignoring errors.
func Sync(log *zap.Logger) {
	if log == nil {
		return
	}
	_ = log.Sync()
}
