package components

func Update() {

	selected := make(map[string]string)
	elements := Table.QuerySelectorAll("tr[data-select=\"true\"]")

	for _, element := range elements {

		id := element.GetAttribute("data-id")
		buttons := element.QuerySelectorAll("button[data-action]")

		for _, button := range buttons {

			action := button.GetAttribute("data-action")

			if action == "fix" {
				selected[id] = action
			} else if action == "commit" {
				selected[id] = action
			} else if action == "pull" || action == "push" {
				selected[id] = "pull-and-push"
			}

		}

	}

	RenderFooter(selected)

}
