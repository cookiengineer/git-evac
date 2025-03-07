package main

import "gooey"
import "gooey/app"
import "git-evac-app/views"
import "time"

func main() {

	element := gooey.Document.QuerySelector("main")

	if element != nil {

		main := app.Main{}
		main.Init(element)

		view := element.GetAttribute("data-view")

		if view == "repositories" {
			main.SetView("repositories", views.NewRepositories(&main))
			main.ChangeView("repositories")
		} else if view == "backups" {
			main.SetView("backups", views.NewBackups(&main))
			main.ChangeView("backups")
		} else if view == "settings" {
			main.SetView("settings", views.NewSettings(&main))
			main.ChangeView("settings")
		}

	}

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
