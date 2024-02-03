package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// responsible for defining range in the game
type game struct {
	UserID string `json:"user_id"`
	Upper  int    `json:"upper"`
	Lower  int    `json:"lower"`
}

var find = "number generated with command !guess <number> try to find out it"

func (h *handler) game(c *gin.Context) {
	// user should guess and server should generate
	// unmarshal json
	var g game
	if err := c.BindJSON(&g); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if g.Upper < g.Lower {
		newErrorResponse(c, http.StatusBadRequest, "upper should be greater than lower try again")
		return
	}
	// validate input and throw for the main logic to the service
	_, err := h.srv.Game.GenerateNumber(context.TODO(), g.UserID, g.Upper, g.Lower)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(201, gin.H{"message": find})
}
