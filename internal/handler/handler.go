package handler

import (
	"github.com/MamushevArup/discord-bot/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler is communication tool and encapsulation
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

// InitRoutes define routes
func (h *handler) InitRoutes() *gin.Engine {

	router := gin.Default()

	router.GET("/hello/:id", h.hello)
	router.GET("/guess/:id/:number", h.guess)
	router.GET("/help", h.help)

	router.POST("/game", h.game)
	router.POST("/poll/:id", h.poll)
	router.POST("/vote/:id", h.vote)

	return router
}
