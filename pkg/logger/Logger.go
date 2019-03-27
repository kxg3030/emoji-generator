package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger *zap.Logger

func EncodeTime(i time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(i.Format("2006-01-02 15:04:05"))
}