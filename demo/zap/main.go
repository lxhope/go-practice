package main

// doc: https://pkg.go.dev/go.uber.org/zap#section-readme

import (
	"time"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("zap sugar logger demo",
		"name", "zap",
		"time", time.Now(),
	)
}
