package views

import "github.com/cookiengineer/gooey/components/app"

func RegisterTo(main *app.Main) {

	main.RegisterView("repositories", app.WrapView(ToRepositories))
	main.RegisterView("backups",      app.WrapView(ToBackups))
	main.RegisterView("settings",     app.WrapView(ToSettings))

}
