package views

import "gooey"
import "gooey/app"
import "gooey/dom"
import app_schemas "git-evac-app/schemas"
import "git-evac/server/schemas"
import "git-evac/structs"
import "sort"
import "strconv"

type Manage struct {
	Main *app.Main `json:"main"`
	app.BaseView
}

func NewManage(main *app.Main) Manage {

	var view Manage

	view.Main     = main
	view.Elements = make(map[string]*dom.Element)

	view.SetElement("table", gooey.Document.QuerySelector("main table"))
	view.SetElement("dialog", gooey.Document.QuerySelector("dialog"))
	view.SetElement("footer", gooey.Document.QuerySelector("footer"))

	view.Init()

	return view

}

func (view Manage) Init() {

	table := view.GetElement("table")

	if table != nil {

		table.QuerySelector("thead input[type=\"checkbox\"]").AddEventListener("change", dom.ToEventListener(func(event dom.Event) {

			target := event.Target

			if target.TagName == "INPUT" {

				is_checked := target.Value.Get("checked").Bool()
				rows := table.QuerySelectorAll("tr[data-id]")

				for _, row := range rows {

					if is_checked == true {
						row.SetAttribute("data-select", "true")
					} else {
						row.SetAttribute("data-select", "false")
					}

					input := row.QuerySelector("input[type=\"checkbox\"]")

					if input != nil {
						input.Value.Set("checked", is_checked)
					}

				}

				view.Update()

			}

		}))

		table.QuerySelector("tbody").AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			target := event.Target

			if target.TagName == "INPUT" {

				row := target.ParentNode().ParentNode()
				is_checked := target.Value.Get("checked").Bool()

				if is_checked == true {
					row.SetAttribute("data-select", "true")
				} else {
					row.SetAttribute("data-select", "false")
				}

				view.Update()

			} else if target.TagName == "BUTTON" {

				row := target.ParentNode().ParentNode()
				id := row.GetAttribute("data-id")
				action := target.GetAttribute("data-action")

				actions := make(map[string]string)
				actions[id] = action

				if len(actions) > 0 {
					view.RenderDialog()
					// TODO: RenderDialog(actions)
					// TODO: Dialog.Open()
				}

			} else if target.TagName == "TD" {

				row := target.ParentNode()
				is_checked := row.GetAttribute("data-select") == "true"

				if is_checked == true {
					row.SetAttribute("data-select", "false")
					row.QuerySelector("input[type=\"checkbox\"]").Value.Set("checked", false)
				} else {
					row.SetAttribute("data-select", "true")
					row.QuerySelector("input[type=\"checkbox\"]").Value.Set("checked", true)
				}

				view.Update()

			} else if target.TagName == "EM" {

				row := target.ParentNode().ParentNode()
				is_checked := row.GetAttribute("data-select") == "true"

				if is_checked == true {
					row.SetAttribute("data-select", "false")
					row.QuerySelector("input[type=\"checkbox\"]").Value.Set("checked", false)
				} else {
					row.SetAttribute("data-select", "true")
					row.QuerySelector("input[type=\"checkbox\"]").Value.Set("checked", true)
				}

				view.Update()

			}

		}))

	}


}

func (view Manage) Enter() bool {

	view.RenderTable()

	return true

}

func (view Manage) Leave() bool {
	return true
}

func (view Manage) Update() {

	selected := app_schemas.Selected{}
	table    := view.GetElement("table")

	view.Main.ReadItem("selected", &selected)

	if table != nil {

		selected.Reset()

		elements := table.QuerySelectorAll("tr[data-select=\"true\"]")

		for _, element := range elements {

			id := element.GetAttribute("data-id")
			buttons := element.QuerySelectorAll("button[data-action]")

			for _, button := range buttons {

				action := button.GetAttribute("data-action")

				if action == "clone" {
					selected.Set(id, action)
				} else if action == "fix" {
					selected.Set(id, action)
				} else if action == "commit" {
					selected.Set(id, action)
				} else if action == "pull" || action == "push" {
					selected.Set(id, "pull-or-push")
				}

			}

		}

		view.Main.SaveItem("selected", selected)
		view.RenderFooter()

	}

}

func (view Manage) RenderTable() {

	index := schemas.Index{}
	table := view.GetElement("table")

	view.Main.ReadItem("index", &index)

	if table != nil {

		html := ""

		for name, owner := range index.Owners {

			for _, repo := range owner.Repositories {
				html += view.renderTableRow(name, repo)
			}

		}

		tbody := table.QuerySelector("tbody")

		if tbody != nil {
			tbody.SetInnerHTML(html)
		}

	}

}

func (view Manage) renderTableRow(owner string, repository *structs.Repository) string {

	var result string

	id := owner + "/" + repository.Name

	result += "<tr data-id=\"" + id + "\" data-select=\"false\">";
	result += "<td><input type=\"checkbox\" data-id=\"" + id + "\" name=\"" + id + "\"/></td>";
	result += "<td><label>" + owner + "/" + repository.Name + "</label></td>";
	result += "<td>"

	remotes := make([]string, 0)

	for name, _ := range repository.Remotes {
		remotes = append(remotes, name)
	}

	sort.Strings(remotes)

	for r := 0; r < len(remotes); r++ {

		remote := remotes[r]
		url := repository.Remotes[remote].URL
		label := toRemoteLabel(remote, url)

		if label != "" {

			if repository.CurrentRemote == remote {
				result += "<label data-remote=\"" + label + "\" class=\"active\">" + remote + "</label>"
			} else {
				result += "<label data-remote=\"" + label + "\">" + remote + "</label>"
			}

		} else {

			if repository.CurrentRemote == remote {
				result += "<label class=\"active\">" + remote + "</label>"
			} else {
				result += "<label>" + remote + "</label>"
			}

		}

		if r < len(remotes) - 1 {
			result += " "
		}

	}

	result += "</td>"
	result += "<td>"

	sort.Strings(repository.Branches)

	for b, branch := range repository.Branches {

		if repository.CurrentBranch == branch {
			result += "<label class=\"active\">" + branch + "</label>"
		} else {
			result += "<label>" + branch + "</label>"
		}

		if b < len(repository.Branches) - 1 {
			result += " "
		}

	}

	result += "</td>"
	result += "<td>"

	if repository.HasRemoteChanges == true {
		result += "<button data-action=\"fix\">Fix</button>";
	} else if repository.HasLocalChanges == true {
		result += "<button data-action=\"commit\">Commit</button>";
	} else {
		result += "<button data-action=\"pull\">Pull</button>";
		result += "<button data-action=\"push\">Push</button>";
	}

	result += "</td>"
	result += "</tr>"

	return result

}

func (view Manage) RenderDialog() {

	// TODO: Render Dialog Title and Contents

}

func (view Manage) RenderFooter() {

	index    := schemas.Index{}
	selected := app_schemas.Selected{}
	footer   := view.GetElement("footer")

	view.Main.ReadItem("index", &index)
	view.Main.ReadItem("selected", &selected)

	if footer != nil {

		clones := selected.Count("clone")
		fixes := selected.Count("fix")
		commits := selected.Count("commit")
		pulls_or_pushes := selected.Count("pull-or-push")

		total := 0

		for _, owner := range index.Owners {
			total += len(owner.Repositories)
		}

		message := "Selected " + strconv.Itoa(selected.Length()) + " of " + strconv.Itoa(total) + " Repositories"

		div1 := footer.QuerySelector("div:first-of-type")

		if div1 != nil {
			div1.SetInnerHTML(message)
		}

		buttons := ""

		if clones > 0 {

			buttons += "<button data-action=\"clone\">Clone " + strconv.Itoa(clones) + "</button>"

			if fixes > 0 {
				buttons += "<button data-action=\"fix\" disabled>Fix " + strconv.Itoa(fixes) + "</button>"
			}

			if commits > 0 {
				buttons += "<button data-action=\"commit\" disabled>Commit " + strconv.Itoa(commits) + "</button>"
			}

			if pulls_or_pushes > 0 {
				buttons += "<button data-action=\"pull\" disabled>Pull " + strconv.Itoa(pulls_or_pushes) + "</button>"
				buttons += "<button data-action=\"push\" disabled>Push " + strconv.Itoa(pulls_or_pushes) + "</button>"
			}

		} else if fixes > 0 {

			buttons += "<button data-action=\"fix\">Fix " + strconv.Itoa(fixes) + "</button>"

			if commits > 0 {
				buttons += "<button data-action=\"commit\" disabled>Commit " + strconv.Itoa(commits) + "</button>"
			}

			if pulls_or_pushes > 0 {
				buttons += "<button data-action=\"pull\" disabled>Pull " + strconv.Itoa(pulls_or_pushes) + "</button>"
				buttons += "<button data-action=\"push\" disabled>Push " + strconv.Itoa(pulls_or_pushes) + "</button>"
			}

		} else if commits > 0 {

			buttons += "<button data-action=\"commit\">Commit " + strconv.Itoa(commits) + "</button>"

			if pulls_or_pushes > 0 {
				buttons += "<button data-action=\"pull\" disabled>Pull " + strconv.Itoa(pulls_or_pushes) + "</button>"
				buttons += "<button data-action=\"push\" disabled>Push " + strconv.Itoa(pulls_or_pushes) + "</button>"
			}

		} else if pulls_or_pushes > 0 {
			buttons += "<button data-action=\"pull\">Pull " + strconv.Itoa(pulls_or_pushes) + "</button>"
			buttons += "<button data-action=\"push\">Push " + strconv.Itoa(pulls_or_pushes) + "</button>"
		}

		div2 := footer.QuerySelector("div:last-of-type")

		if div2 != nil {
			div2.SetInnerHTML(buttons)
		}

	}

}
