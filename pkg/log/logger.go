package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *zap.Logger
)

func init() {
	InitLogger(zapcore.InfoLevel)
}

func InitLogger(logLevel zapcore.Level) {
	once.Do(func() {
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:      "timestamp",
			LevelKey:     "level",
			MessageKey:   "message",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		}

		core := zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),
				zapcore.AddSync(zapcore.AddSync(zapcore.Lock(zapcore.AddSync(os.Stdout)))),
				logLevel,
			),
		)

		logger = zap.New(core, zap.AddCaller())
	})
}

func GetLogger() *zap.Logger {
	if logger == nil {
		panic("Logger not initialized. Call InitLogger first.")
	}
	return logger
}
