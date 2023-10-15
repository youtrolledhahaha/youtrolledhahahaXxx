package main

import (
	"github.com/youtrolledhahaha/youtrolledhahahaXxxclient/app"
	"github.com/youtrolledhahaha/youtrolledhahahaXxxclient/app/environment"
	"github.com/youtrolledhahaha/youtrolledhahahaXxxclient/app/ui"
)

var (
	Version       = "dev"
	Port          = ""
	ServerAddress = ""
	Token         = ""
)

func main() {
	ui.ShowMenu(Version, ServerAddress, Port)

	app.New(environment.Load(ServerAddress, Port, Token)).Run()
}
