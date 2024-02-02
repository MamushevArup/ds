package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// responsible for defining range in the game
type game struct {
	Upper int `json:"upper"`
	Lower int `json:"lower"`
}

func (h *handler) game(c *gin.Context) {
	// user should guess and server should generate
	var g game
	if err := c.BindJSON(&g); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if g.Upper < g.Lower {
		newErrorResponse(c, http.StatusBadRequest, "upper should be greater than lower try again")
		return
	}
	// pass the lower and upper bound
	number, err := h.srv.Gamer.GenerateNumber(context.TODO(), g.Upper, g.Lower)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(201, gin.H{"message": number})
}
