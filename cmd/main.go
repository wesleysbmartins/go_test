package main

import server "go_tests/internal/services"

func main() {
	api := &server.ServerApi{}
	api.Run()
}
