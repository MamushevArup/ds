package handler

import (
	"github.com/MamushevArup/discord-bot/internal/service"
	"github.com/gin-gonic/gin"
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
	return router
}
