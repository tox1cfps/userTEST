package main

import (
	"userTest/config"
	"userTest/internal/ui"
)

func main() {
	config.Connect()
	ui.UserMenu()
}
