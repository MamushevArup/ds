package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) poll(c *gin.Context) {
	id := c.Param("id")
	type poll struct {
		Question string   `json:"question"`
		Options  []string `json:"options"`
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
