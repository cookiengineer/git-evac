package controllers

import "github.com/cookiengineer/gooey/components/app"

func RegisterTo(main *app.Main) {

	main.RegisterController("repositories", app.WrapController(NewRepositories))
	main.RegisterController("backups",      app.WrapController(NewBackups))
	main.RegisterController("settings",     app.WrapController(NewSettings))

}
