//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/interfaces"
import content_components "github.com/cookiengineer/gooey/components/content"

func RegisterTo(document *components.Document) {

	document.Register("table", func(element *dom.Element) interfaces.Component {

		typ := element.GetAttribute("data-type")

		if typ == "repositories" {
			return ToRepositoriesTable(element)
		} else if typ == "backups" {
			// TODO: BackupsTable
			return nil
		} else if typ == "scheduler" {
			return ToSchedulerTable(element)
		} else {
			return content_components.ToTable(element)
		}

	})

}
