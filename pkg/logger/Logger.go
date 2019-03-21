package logger

import (
	"emoji/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger *zap.Logger

func EncodeTime(i time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(i.Format(config.HourFormat))
}