package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) guess(c *gin.Context) {
	// get the dynamic parameter from url
	id := c.Param("id")
	num := c.Param("number")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "id is empty")
		return
	}
	// validate and throw to the service
	feedback, err := h.srv.Guess.MatchNumbers(context.Background(), id, num)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, gin.H{"message": feedback})
}
