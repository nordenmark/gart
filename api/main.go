package main

import (
	"gart/server"
)

func main() {
	serv := server.CreateServer()

	serv.Run()
}
