package components

import "gooey"
import "gooey/dom"
import "gooey/elements"
import "slices"
import "sort"
import "strings"

var Dialog *elements.Dialog = nil

func init() {

	element := gooey.Document.QuerySelector("dialog")

	if element != nil {
		dialog := elements.ToDialog(element)
		Dialog = &dialog
	}

}

func InitDialog() {

	Dialog.Element.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

		target := event.Target
		action := target.GetAttribute("data-action")

		if target.TagName == "BUTTON" {

			if action == "open-terminal" {
			} else if action == "fix" {
			} else if action == "commit" {
			} else if action == "pull" {
			} else if action == "push" {
			} else if action == "cancel" {
				Dialog.Close()
			} else if action == "close" {
				Dialog.Close()
			}

		}

	}))

}

func RenderDialog(selected map[string]string) {

	if Dialog != nil {

		actions := make([]string, 0)
		ids := make([]string, 0)

		for id, action := range selected {

			ids = append(ids, id)

			if slices.Contains(actions, action) == false {
				actions = append(actions, action)
			}

		}

		sort.Strings(ids)

		html := ""

		for i := 0; i < len(ids); i++ {
			html += RenderDialogTableRow(ids[i])
		}

		tbody := Dialog.Element.QuerySelector("table tbody")

		if tbody != nil {
			tbody.SetInnerHTML(html)
		}

		RenderDialogFooter(actions)

	}

}

func RenderDialogTableRow(identifier string) string {

	html := ""

	html += "<tr data-id=\"" + identifier + "\">"
	html += "<td>" + identifier + "</td>"
	html += "<td><input type=\"checkbox\"/></td>"
	html += "</tr>"

	return html

}

func RenderDialogFooter(actions []string) {

	div := Dialog.Element.QuerySelector("div:last-of-type")

	if div != nil {

		html := ""

		for a := 0; a < len(actions); a++ {

			action := actions[a]
			label := strings.ToUpper(action[0:1]) + strings.ToLower(action[1:])

			html += "<button data-action=\"" + action + "\">" + label + "</button>"

		}

		div.SetInnerHTML(html)

	}

}
