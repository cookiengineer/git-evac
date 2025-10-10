package views

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/interfaces"

func RegisterTo(main *app.Main) {

	main.RegisterView("repositories", func(element *dom.Element) interfaces.View {
		return ToRepositories(element)
	})

	main.RegisterView("backups", func(element *dom.Element) interfaces.View {
		// TODO: Change this to ToBackups(element)
		return ToRepositories(element)
	})

}
