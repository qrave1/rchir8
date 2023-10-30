package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"log"
	"rchir8/internal/config"
	"rchir8/internal/handler"
)

func main() {
	var cfg = config.NewConfig()
	var s = securecookie.New(cfg.Hashkey, cfg.BlockKey)
	var cont = handler.NewController(cfg, s)

	r := gin.Default()
	r.GET("/", cont.ReadCookie)
	r.POST("/", cont.SetCookie)

	log.Fatal(r.Run())
}
