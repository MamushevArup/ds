package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
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
