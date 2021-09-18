package main

import "github.com/ProjectAthenaa/scheduler-service"

func main() {
	go handleServer()
	scheduler.NewScheduler().ProcessEvents()
}
