package components

import "gooey"
import "gooey/dom"
import "strconv"

var Footer *dom.Element = nil

func init() {

	element := gooey.Document.QuerySelector("footer")

	if element != nil {
		Footer = element
	}

}

func InitFooter() {

	Footer.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

		target := event.Target

		if target.TagName == "BUTTON" {

			action := target.GetAttribute("data-action")
			actions := make(map[string]string)

			rows := Table.QuerySelectorAll("tbody tr[data-select=\"true\"]")

			for r := 0; r < len(rows); r++ {

				row := rows[r]
				id := row.GetAttribute("data-id")
				has_action := row.QuerySelector("button[data-action=\"" + action + "\"]") != nil

				if has_action == true {
					actions[id] = action
				}

			}

			if len(actions) > 0 {
				RenderDialog(actions)
				Dialog.Open()
			}

		}

	}))

}

