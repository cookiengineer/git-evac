package views

import "github.com/cookiengineer/gooey/components/app"

func RegisterTo(main *app.Main) {

	main.RegisterView("repositories", app.WrapView(ToRepositories))

	// TODO: Change this to ToBackups(element)
	main.RegisterView("backups",      app.WrapView(ToRepositories))

}
