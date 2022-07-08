package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const kopeksInRub = 100

func (s *Server) transferMoney(ctx *gin.Context) {
	var transData transaction
	ctx.BindJSON(&transData)

	if transData.Sum < 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"description": "negative sum to transfer",
		})
		return
	}

	s.accStorage.Lock()
	defer s.accStorage.Unlock()

	sumToTransfer := uint(transData.Sum * kopeksInRub) // Rubles to kopeks

	err := s.accStorage.SubtractMoney(transData.Sender, sumToTransfer)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"description": "incorrect request's data",
		})
		return
	}

	err = s.accStorage.AddMoney(transData.Receiver, sumToTransfer)
	if err != nil {
		// If we can't increase the receiver's balance, we "return" money to sender.
		// We do that as this point is reachable in case when we have already taken
		// sender's money away. Kind of "rollback".
		s.accStorage.AddMoney(transData.Sender, sumToTransfer)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"description": "incorrect request's data",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"description": "transaction successfuly completed",
	})
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
	s.accStorage.RLock()
	defer s.accStorage.RUnlock()

	accountMoney := uint(regData.Money * kopeksInRub)

	s.accStorage.AddAccount(accountMoney)
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

	s.accStorage.RLock()
	defer s.accStorage.RUnlock()
	account, ok := s.accStorage.GetAccount(uint(id))
	if !ok {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"description": "account not found"})
		return
	}

	accForResponse := AccountForResponse{
		Id:    account.GetId(),
		Money: float32(account.GetMoney()) / kopeksInRub,
	}
	ctx.IndentedJSON(http.StatusOK, accForResponse)
}
