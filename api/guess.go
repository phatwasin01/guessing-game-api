package api

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GuessRequest struct {
	Number int `json:"number"`
}
type GuessResponse struct {
	NewNumber int `json:"new_number"`
}

var hiddenNumber = 8

func (server *Server) guessNumber(ctx *gin.Context) {
	var req GuessRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if req.Number != hiddenNumber {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong guess"})
		return
	}
	hiddenNumber = rand.Intn(10)
	var res GuessResponse
	res.NewNumber = hiddenNumber
	ctx.JSON(http.StatusCreated, res)

}
