package views

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/interfaces"

func RegisterTo(main *app.Main) {

	main.RegisterView("repositories", func(element *dom.Element) interfaces.View {
		return ToRepositoriesView(element)
	})

	main.RegisterView("backups", func(element *dom.Element) interfaces.View {
		return ToRepositoriesView(element)
	})

}
