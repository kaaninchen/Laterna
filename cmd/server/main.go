package main

import (
	"github.com/kaaninchen/laterna/server/cmd/internal/config"
	"github.com/kaaninchen/laterna/server/cmd/internal/db"
	"github.com/kaaninchen/laterna/server/cmd/internal/http"
	"github.com/kaaninchen/laterna/server/cmd/internal/logger"
)

func main() {
	config.AppConfig = config.LoadConfig("config.json")
	logger.InitLoggers()

	db.ConnectDB()
	http.StartHTTPServer()
}
