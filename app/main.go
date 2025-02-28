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

		if view == "manage" {
			main.SetView("manage", views.NewManage(&main))
			main.ChangeView("manage")
		// } else if view == "backup" {
		//	main.SetView("backup", views.NewBackup(&main))
		// } else if view == "restore" {
		//	main.SetView("restore", views.NewBackup())
		//} else if view == "settings" {
		//	main.SetView("settings", views.NewSettings())
		}

	}

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
