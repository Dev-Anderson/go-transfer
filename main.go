package main

import "go-transfer/server"

func main() {
	s := server.NewServer()
	s.Run()
}
