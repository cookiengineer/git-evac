package components

import "gooey"
import "gooey/dom"
import "app/storage"
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

func RenderFooter(selected map[string]string) {

	fixes := make([]string, 0)
	commits := make([]string, 0)
	pulls_or_pushes := make([]string, 0)

	for id, action := range selected {

		if action == "fix" {
			fixes = append(fixes, id)
		} else if action == "commit" {
			commits = append(commits, id)
		} else if action == "pull-and-push" {
			pulls_or_pushes = append(pulls_or_pushes, id)
		}

	}

	total_amount := 0

	if storage.Index != nil {

		for _, user := range storage.Index.Users {
			total_amount += len(user.Repositories)
		}

		for _, orga := range storage.Index.Organizations {
			total_amount += len(orga.Repositories)
		}

	}

	message := "Selected " + strconv.Itoa(len(selected)) + " of " + strconv.Itoa(total_amount) + " Repositories"

	buttons := ""

	if len(fixes) > 0 {
		buttons += "<button data-action=\"fix\">Fix " + strconv.Itoa(len(fixes)) + "</button>"
		buttons += "<button data-action=\"commit\" disabled>Commit " + strconv.Itoa(len(commits)) + "</button>"
		buttons += "<button data-action=\"pull\" disabled>Pull " + strconv.Itoa(len(pulls_or_pushes)) + "</button>"
		buttons += "<button data-action=\"push\" disabled>Push " + strconv.Itoa(len(pulls_or_pushes)) + "</button>"
	} else if len(commits) > 0 {
		buttons += "<button data-action=\"commit\">Commit " + strconv.Itoa(len(commits)) + "</button>"
		buttons += "<button data-action=\"pull\" disabled>Pull " + strconv.Itoa(len(pulls_or_pushes)) + "</button>"
		buttons += "<button data-action=\"push\" disabled>Push " + strconv.Itoa(len(pulls_or_pushes)) + "</button>"
	} else if len(pulls_or_pushes) > 0 {
		buttons += "<button data-action=\"pull\">Pull " + strconv.Itoa(len(pulls_or_pushes)) + "</button>"
		buttons += "<button data-action=\"push\">Push " + strconv.Itoa(len(pulls_or_pushes)) + "</button>"
	}

	element1 := Footer.QuerySelector("div:first-of-type")

	if element1 != nil {
		element1.SetInnerHTML(message)
	}

	element2 := Footer.QuerySelector("div:last-of-type")

	if element2 != nil {
		element2.SetInnerHTML(buttons)
	}

}
