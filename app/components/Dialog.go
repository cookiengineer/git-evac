package components

import "fmt"
import "gooey"
import "gooey/dom"
import "gooey/elements"

var Dialog *elements.Dialog = nil

func init() {

	element := gooey.Document.QuerySelector("dialog")

	if element != nil {
		dialog := elements.ToDialog(element)
		Dialog = &dialog
	}

}

func InitDialog() {

	button := Dialog.Element.QuerySelector("button[data-action]")

	if button != nil {

		button.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			action := event.Target.GetAttribute("data-action")

			if action == "close" {
				Dialog.Close()
			}

		}))

	}

}

func RenderDialog(selected map[string]string) {

	if Dialog != nil {

		// TODO: Render Dialog content

		fmt.Println("TODO: Render Dialog Contents")

	}

}
