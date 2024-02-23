package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	logger *zap.SugaredLogger
}

var instance *logger

func New() {
	config := zap.NewProductionConfig()
	addEncoder(&config)
	addLevel(&config)
	zapLogger, _ := config.Build()

	logger := &logger{
		logger: zapLogger.Sugar(),
	}
	instance = logger
}

func addEncoder(config *zap.Config) {
	config.Encoding = "console"
	config.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.RFC3339TimeEncoder,
	}
}

func addLevel(config *zap.Config) {
	atom := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.Level = atom
}

func Debug(args ...any) {
	instance.logger.Debug(args...)
}

func Debugf(template string, args ...any) {
	instance.logger.Debugf(template, args...)
}

func Info(args ...any) {
	instance.logger.Info(args...)
}

func Infof(template string, args ...any) {
	instance.logger.Infof(template, args...)
}

func Warn(args ...any) {
	instance.logger.Warn(args...)
}

func Warnf(template string, args ...any) {
	instance.logger.Warnf(template, args...)
}

func Error(args ...any) {
	instance.logger.Error(args...)
}

func Errorf(template string, args ...any) {
	instance.logger.Errorf(template, args...)
}

func Fatal(args ...any) {
	instance.logger.Fatal(args...)
}

func Fatalf(template string, args ...any) {
	instance.logger.Fatalf(template, args...)
}
