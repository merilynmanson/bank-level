package server

import (
	"bank/internal/storage"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	accStorage *storage.AccountStorage
}

func NewServer() *Server {
	return &Server{
		router:     gin.Default(),
		accStorage: storage.NewAccountStorage(),
	}
}

func (s *Server) initHandlers() {
	s.router.POST("/accounts", s.createAccount)
	s.router.GET("/accounts/:id", s.getAccount)
	s.router.POST("/transactions", s.transferMoney)
}

func (s *Server) Run(addr string) {
	s.initHandlers()
	s.router.Run(addr)
}
