package handler

import (
	"github.com/MamushevArup/discord-bot/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	router.POST("/vote/:id", h.vote)
	return router
}

func (h *handler) vote(c *gin.Context) {
	id := c.Param("id")
	type vote struct {
		Question string `json:"question"`
		Option   string `json:"option"`
	}
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "no id provided")
		return
	}
	var v vote
	if err := c.BindJSON(&v); err != nil || v.Option == "" || v.Question == "" {
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	counter, err := h.srv.Poll.Vote(id, v.Question, v.Option)
	if err != nil {
		log.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	cnt := strconv.Itoa(counter)
	c.JSON(201, gin.H{"message": "your vote counted for this vote counter " + cnt})
}
