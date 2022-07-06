package main

import (
	"villa-akmali/api"
	"villa-akmali/config"
)


func main() {
	config.Init()
	api.Run()
}