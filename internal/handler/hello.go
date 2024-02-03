package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

var helloMessage = "Hey let's start our journey use !help to see all commands available"

func (h *handler) hello(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "id is not specified")
		return
	}
	err := h.srv.Hello.CreateUser(context.Background(), id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, gin.H{"message": helloMessage})
}
