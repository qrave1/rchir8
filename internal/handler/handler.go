package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"log/slog"
	"rchir8/internal/config"

	"net/http"
)

type UserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Controller struct {
	log *slog.Logger
	cfg *config.Config
	s   *securecookie.SecureCookie
}

func NewController(log *slog.Logger, cfg *config.Config, s *securecookie.SecureCookie) *Controller {
	return &Controller{log: log, cfg: cfg, s: s}
}

func (cont *Controller) SetCookie(c *gin.Context) {
	cont.log.Info("Start processing SetCookie")
	defer cont.log.Info("Finish processing SetCookie")
	var userData UserData
	err := c.ShouldBindJSON(&userData)
	if err != nil {
		cont.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userData.Email == "" || userData.Password == "" {
		cont.log.Error("bad data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad data"})
		return
	}

	value := map[string]string{
		userData.Email: userData.Password,
	}

	encoded, err := cont.s.Encode(cont.cfg.Cookie, value)
	if err != nil {
		cont.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(
		cont.cfg.Cookie,
		encoded,
		3600,
		"/",
		"localhost",
		true,
		true,
	)

	c.JSON(http.StatusCreated, gin.H{"status": "OK"})
}

func (cont *Controller) ReadCookie(c *gin.Context) {
	cont.log.Info("Start processing ReadCookie")
	defer cont.log.Info("Finish processing ReadCookie")
	value := make(map[string]string)
	if cookie, err := c.Cookie(cont.cfg.Cookie); err == nil {
		if err = cont.s.Decode(cont.cfg.Cookie, cookie, &value); err != nil {
			cont.log.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
	var email, pass string
	for k, v := range value {
		email = k
		pass = v
	}

	c.JSON(http.StatusOK, gin.H{
		"Cookie":   cont.cfg.Cookie,
		"Email":    email,
		"Password": pass,
	})
}
