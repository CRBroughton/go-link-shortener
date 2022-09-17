package main

import (
	"goly/model"
	"goly/utils/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
