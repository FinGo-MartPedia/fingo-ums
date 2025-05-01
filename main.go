package main

import (
	"github.com/fingo-martPedia/fingo-ums/cmd"
	"github.com/fingo-martPedia/fingo-ums/helpers"
)

func main() {
	helpers.SetupLogger()

	helpers.SetupDatabase()

	go cmd.ServeGRPC()

	cmd.ServeHTTP()
}
