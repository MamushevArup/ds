package handler

import (
	"encoding/json"
	"github.com/MamushevArup/discord-bot/internal/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type Handler interface {
	InitRoutes() *gin.Engine
}

type handler struct {
	srv *service.Service
}

func NewBot(srv *service.Service) Handler {
	return &handler{
		srv: srv,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/hello/:id", h.hello)
	router.POST("/game", h.game)
	router.GET("/guess/:id/:number", h.guess)
	router.GET("/help", h.help)
	return router
}

func (h *handler) help(c *gin.Context) {
	open, err := os.Open("./helper.json")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	all, err := io.ReadAll(open)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var hm map[string]interface{}
	err = json.Unmarshal(all, &hm)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, hm)
}
