package main

import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/layout"
import "github.com/cookiengineer/gooey/components/ui"
import app_components "git-evac-app/components"
import app_controllers "git-evac-app/controllers"
import app_views "git-evac-app/views"
import "time"

func main() {

	main := app.NewMain()

	// Register Gooey Components
	content.RegisterTo(main.Document)
	layout.RegisterTo(main.Document)
	ui.RegisterTo(main.Document)

	// Register App Components
	app_components.RegisterTo(main.Document)

	// Register App Controllers
	app_controllers.RegisterTo(main)

	// Register App Views
	app_views.RegisterTo(main)

	// Start the App
	main.Mount()
	main.Render()

	for true {
		time.Sleep(1 * time.Second)
	}

}
