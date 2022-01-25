package logging

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	w := io.Writer(os.Stdout)
	return zap.New(zapcore.NewCore(DefaultEncoder(), zapcore.AddSync(w), zap.DebugLevel), zap.AddCaller())
}

func DefaultEncoder() zapcore.Encoder {
	conf := zap.NewDevelopmentEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(conf)
	return encoder
}
