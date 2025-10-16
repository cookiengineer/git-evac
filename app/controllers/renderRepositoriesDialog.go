//go:build wasm

package controllers

import layout_components "github.com/cookiengineer/gooey/components/layout"
import "fmt"

func renderRepositoriesDialog(dialog *layout_components.Dialog, actions map[string]any) {

	fmt.Println("Render DIALOG NOW")
	fmt.Println(dialog)
	fmt.Println(actions)

}
