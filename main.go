package main

import (
	"go-holiday/schedule"
	"go-holiday/server"
	"log/slog"
)

func main() {
	initConfig()

	err := schedule.StartSchedule()
	if err != nil {
		slog.Error("failed to start schedule", "error", err)
		return
	}

	httpServer, err := server.CreateServer()
	if err != nil {
		slog.Error("failed to create server", "error", err)
		return
	}

	err = httpServer.Run()
	if err != nil {
		slog.Error("failed to run server", "error", err)
		return
	}
}
