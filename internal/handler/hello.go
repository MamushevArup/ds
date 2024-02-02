package handler

import "github.com/gin-gonic/gin"

var helloMessage = "Hey let's start our journey use !help to see all commands available"

func (h *handler) hello(c *gin.Context) {
	_ = c.Param("id")

	c.JSON(200, gin.H{"message": helloMessage})
}
