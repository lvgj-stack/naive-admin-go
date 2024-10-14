package log

import (
	"go.uber.org/zap"
)

func InitLog() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
