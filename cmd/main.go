package main

import "bank/internal/server"

func main() {
	server := server.NewServer()
	server.Run("127.0.0.1:8080")
}
