package main

import (
	"Reminders/internal/server"
)

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
}
