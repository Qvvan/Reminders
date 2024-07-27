package main

import (
	"Reminders/internal/handlers"
	"Reminders/internal/server"
)

func init() {
	server.InitServer()
	handlers.SetLogger(server.GetLogger())
}

func main() {
	server.StartServer()
}
