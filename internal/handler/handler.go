package handler

import (
	"github.com/MamushevArup/discord-bot/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.POST("/poll/:id", h.poll)
	return router
}

func (h *handler) poll(c *gin.Context) {
	id := c.Param("id")
	type poll struct {
		Question string         `json:"question"`
		Options  map[int]string `json:"options"`
	}
	var p poll
	if err := c.BindJSON(&p); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	err := h.srv.Poll.CreatePoll(id, p.Question, p.Options)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(201, gin.H{"message": "Poll created to vote use !vote <question> <option> command"})
}
