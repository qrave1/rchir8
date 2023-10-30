package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"rchir8/internal/config"

	"net/http"
)

type UserData struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Controller struct {
	cfg *config.Config
	s   *securecookie.SecureCookie
}

func NewController(cfg *config.Config, s *securecookie.SecureCookie) *Controller {
	return &Controller{cfg: cfg, s: s}
}

func (cont *Controller) SetCookie(c *gin.Context) {
	var userData UserData
	err := c.Bind(&userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userData.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	value := map[string]string{
		userData.Email: userData.Password,
	}

	encoded, err := cont.s.Encode(cont.cfg.Cookie, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
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
	value := make(map[string]string)
	if cookie, err := c.Cookie(cont.cfg.Cookie); err == nil {
		if err = cont.s.Decode(cont.cfg.Cookie, cookie, &value); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
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
