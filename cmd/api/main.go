package main

import "github.com/seinyan/go-rest-api/internal/api"

func main() {
	s := api.NewServer()
	s.Start()
}