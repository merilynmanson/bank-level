package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) transferMoney(ctx *gin.Context) {
	var transData transaction
	ctx.BindJSON(&transData)

	err := s.accStorage.SubtractMoney(transData.Sender, transData.Sum)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"description": "sender's account not found",
		})
		return
	}
	err = s.accStorage.AddMoney(transData.Receiver, transData.Sum)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"description": "receiver's account not found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"description": "transaction successfuly completed",
	})
	s.accStorage.PrintAccs()
}

func (s *Server) createAccount(ctx *gin.Context) {
	var regData accRegData

	err := ctx.BindJSON(&regData)
	if err != nil {
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "incorrect params",
		})
		return
	}

	s.accStorage.AddAccount(regData.Money)
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"description": "account was successfuly created",
	})
}

func (s *Server) getAccount(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"description": "incorrect id parameter"})
		return
	}
	account, ok := s.accStorage.GetAccount(uint(id))
	if !ok {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"description": "account not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, *account)
}
