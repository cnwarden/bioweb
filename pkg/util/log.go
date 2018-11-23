package log

import (
	"go.uber.org/zap"
	"fmt"
	"os"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	os.Mkdir("logs", os.ModePerm)
	fmt.Println("init logger")
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		"stdout",
		"logs/app.log",
	}
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	Logger, _ = config.Build()

	defer Logger.Sync()
}
