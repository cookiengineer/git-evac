package controllers

import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/interfaces"

func RegisterTo(main *app.Main) {

	main.RegisterController("repositories", func(main *app.Main, view interfaces.View) interfaces.Controller {
		return NewRepositories(main, view)
	})

	main.RegisterController("backups", func(main *app.Main, view interfaces.View) interfaces.Controller {
		return NewBackups(main, view)
	})

	main.RegisterController("settings", func(main *app.Main, view interfaces.View) interfaces.Controller {
		return NewSettings(main, view)
	})

}
