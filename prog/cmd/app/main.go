package main

import (
	menu "prog/cmd/modes/techUI"
	"prog/cmd/registry"
	"log"
)

func main() {
	app := registry.App{}

	err := app.Config.ParseConfig("config.json", "../../config")
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}

	if app.Config.Mode == "tech" {
		app.Logger.Info("Start with tech ui!")
		menu.RunMenu(app.Services)
	} else {
		app.Logger.Error("Wrong app mode", "mode", app.Config.Mode)
	}
}