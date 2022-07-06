package main

import (
	"studybuddy-backend-fast/api"
	"studybuddy-backend-fast/config"
)

func main() {
	config.Init()
	api.Run()
}
