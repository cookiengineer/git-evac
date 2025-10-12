//go:build wasm

package components

import "github.com/cookiengineer/gooey/components"

func RegisterTo(document *components.Document) {

	document.Register("table", components.WrapComponent(ToRepositoriesTable))

}
