package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) poll(c *gin.Context) {
	id := c.Param("id")
	// use local struct
	type poll struct {
		Question string   `json:"question"`
		Options  []string `json:"options"`
	}
	// create instance of struct
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
