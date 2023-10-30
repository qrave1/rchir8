package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"log"
	"log/slog"
	"os"
	"rchir8/internal/config"
	"rchir8/internal/handler"
)

func main() {
	opts := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opts))
	slog.SetDefault(logger)

	var cfg = config.NewConfig()
	var s = securecookie.New(cfg.Hashkey, cfg.BlockKey)
	var cont = handler.NewController(logger, cfg, s)

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", cont.ReadCookie)
		api.POST("/", cont.SetCookie)
	}

	log.Fatal(r.Run())
}
