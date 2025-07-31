package views

import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/bindings/dom"
import "git-evac-app/actions"
import app_schemas "git-evac-app/schemas"
import "git-evac/server/schemas"
import "git-evac/structs"
import "sort"
import "strconv"
import "strings"

type Repositories struct {
	Main     *app.Main             `json:"main"`
	Schema   *schemas.Repositories `json:"schema"`
	Selected *app_schemas.Selected `json:"selected"`
	app.BaseView
}

func NewRepositories(main *app.Main) Repositories {

	var view Repositories

	view.Main     = main
	view.Schema   = &schemas.Repositories{}
	view.Selected = &app_schemas.Selected{}
	view.Elements = make(map[string]*dom.Element)

	view.SetElement("table", dom.Document.QuerySelector("main table"))
	view.SetElement("dialog", dom.Document.QuerySelector("dialog"))
	view.SetElement("footer", dom.Document.QuerySelector("footer"))

	view.Init()

	return view

}

func (view Repositories) Init() {

	dialog := view.GetElement("dialog")
	footer := view.GetElement("footer")
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

				view.Selected.Reset()
				view.Selected.Set(id, action)
				view.renderDialog(action)

				dialog.SetAttribute("open", "")

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

			} else if target.TagName == "LABEL" {

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

	if dialog != nil {

		dialog.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			target := event.Target
			action := target.GetAttribute("data-action")

			if target.TagName == "BUTTON" {

				if action == "confirm" {

					buttons := dialog.QuerySelectorAll("footer button[data-action]")

					for b := 0; b < len(buttons); b++ {
						buttons[b].SetAttribute("disabled", "")
					}

					rows := dialog.QuerySelectorAll("table tbody tr[data-id]")

					for r := 0; r < len(rows); r++ {

						label      := rows[r].QuerySelector("label[data-state]")
						id         := rows[r].GetAttribute("data-id")
						action     := rows[r].GetAttribute("data-action")
						owner      := id[0:strings.Index(id, "/")]
						repository := id[strings.Index(id, "/")+1:]

						if action == "fix" {

							go func(label *dom.Element, owner string, repository string) {

								response, err := actions.Fix(owner, repository)

								if err == nil {

									if label != nil {
										label.SetAttribute("data-state", "success")
										label.SetAttribute("title", "success!")
									}

									view.UpdateRepository(*response)

								} else {

									if label != nil {
										label.SetAttribute("data-state", "failure")
										label.SetAttribute("title", "failure!")
									}

								}

								for b := 0; b < len(buttons); b++ {
									buttons[b].RemoveAttribute("disabled")
								}

							}(label, owner, repository)

						} else if action == "clone" {

							go func(label *dom.Element, owner string, repository string) {

								response, err := actions.Clone(owner, repository)

								if err == nil {

									if label != nil {
										label.SetAttribute("data-state", "success")
										label.SetAttribute("title", "success!")
									}

									view.UpdateRepository(*response)

								} else {

									if label != nil {
										label.SetAttribute("data-state", "failure")
										label.SetAttribute("title", "failure!")
									}

								}

								for b := 0; b < len(buttons); b++ {
									buttons[b].RemoveAttribute("disabled")
								}

							}(label, owner, repository)

						} else if action == "commit" {

							go func(label *dom.Element, owner string, repository string) {

								response, err := actions.Commit(owner, repository)

								if err == nil {

									if label != nil {
										label.SetAttribute("data-state", "success")
										label.SetAttribute("title", "success!")
									}

									view.UpdateRepository(*response)

								} else {

									if label != nil {
										label.SetAttribute("data-state", "failure")
										label.SetAttribute("title", "failure!")
									}

								}

								for b := 0; b < len(buttons); b++ {
									buttons[b].RemoveAttribute("disabled")
								}

							}(label, owner, repository)

						} else if action == "pull" {

							go func(label *dom.Element, owner string, repository string) {

								response, err := actions.Pull(owner, repository)

								if err == nil {

									if label != nil {
										label.SetAttribute("data-state", "success")
										label.SetAttribute("title", "success!")
									}

									view.UpdateRepository(*response)

								} else {

									if label != nil {
										label.SetAttribute("data-state", "failure")
										label.SetAttribute("title", "failure!")
									}

								}

								for b := 0; b < len(buttons); b++ {
									buttons[b].RemoveAttribute("disabled")
								}

							}(label, owner, repository)

						} else if action == "push" {

							go func(label *dom.Element, owner string, repository string) {

								response, err := actions.Push(owner, repository)

								if err == nil {

									if label != nil {
										label.SetAttribute("data-state", "success")
										label.SetAttribute("title", "success!")
									}

									view.UpdateRepository(*response)

								} else {

									if label != nil {
										label.SetAttribute("data-state", "failure")
										label.SetAttribute("title", "failure!")
									}

								}

								for b := 0; b < len(buttons); b++ {
									buttons[b].RemoveAttribute("disabled")
								}

							}(label, owner, repository)

						}

					}

				} else if action == "cancel" {
					dialog.RemoveAttribute("open")
				} else if action == "close" {
					dialog.RemoveAttribute("open")
				}

			}

		}))

	}

	if footer != nil {

		footer.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			target := event.Target

			if target.TagName == "BUTTON" {

				action := target.GetAttribute("data-action")

				if action != "" {
					view.renderDialog(action)
					dialog.SetAttribute("open", "")
				}

			}

		}))

	}

}

func (view Repositories) Enter() bool {

	schema, err := actions.Index()

	if err == nil {
		view.Schema.Owners = schema.Owners
		view.Main.Storage.Write("repositories", schema)
	}

	view.Render()

	return true

}

func (view Repositories) Leave() bool {
	return true
}

func (view Repositories) Update() {

	table := view.GetElement("table")

	if table != nil {

		view.Selected.Reset()

		elements := table.QuerySelectorAll("tr[data-select=\"true\"]")

		for _, element := range elements {

			id := element.GetAttribute("data-id")
			buttons := element.QuerySelectorAll("button[data-action]")

			for _, button := range buttons {

				action := button.GetAttribute("data-action")

				if action == "clone" {
					view.Selected.Set(id, action)
				} else if action == "fix" {
					view.Selected.Set(id, action)
				} else if action == "commit" {
					view.Selected.Set(id, action)
				} else if action == "pull" || action == "push" {
					view.Selected.Set(id, "pull-or-push")
				}

			}

		}

		view.renderFooter()

	}

}

func (view Repositories) UpdateRepository(updated schemas.Repository) {

	schema := schemas.Repositories{}
	view.Main.Storage.Read("repositories", &schema)

	has_changed := false

	for owner_name, owner := range schema.Owners {

		for repo_name, repo := range owner.Repositories {

			if repo.Folder == updated.Repository.Folder {
				schema.Owners[owner_name].Repositories[repo_name] = &updated.Repository
				has_changed = true
				break
			}

		}

	}

	if has_changed == true {
		view.Schema.Owners = schema.Owners
		view.Main.Storage.Write("repositories", schema)
		view.Render()
	}

}

func (view Repositories) Render() {

	table := view.GetElement("table")

	if table != nil {

		html := ""

		owners := make([]string, 0)

		for name := range view.Schema.Owners {
			owners = append(owners, name)
		}

		sort.Strings(owners)

		for _, owner := range owners {

			repositories := make([]string, 0)

			for name := range view.Schema.Owners[owner].Repositories {
				repositories = append(repositories, name)
			}

			sort.Strings(repositories)

			for _, repo := range repositories {
				html += view.renderTableRow(owner, view.Schema.Owners[owner].Repositories[repo])
			}

		}

		tbody := table.QuerySelector("tbody")

		if tbody != nil {
			tbody.SetInnerHTML(html)
		}

	}

}

func (view Repositories) renderTableRow(owner string, repository *structs.Repository) string {

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

func (view Repositories) renderDialog(action string) {

	dialog := view.GetElement("dialog")

	if dialog != nil {

		selected := view.Selected.Copy()

		if action == "fix" {
			selected.FilterByValue("fix")
		} else if action == "clone" {
			selected.FilterByValue("clone")
		} else if action == "commit" {
			selected.FilterByValue("commit")
		} else if action == "pull" {

			selected.FilterByValue("pull-or-push")

			for name := range selected {
				selected.Set(name, "pull")
			}

		} else if action == "push" {

			selected.FilterByValue("pull-or-push")

			for name := range selected {
				selected.Set(name, "push")
			}

		}

		h3 := dialog.QuerySelector("h3")

		if h3 != nil {

			title := strings.ToUpper(action[0:1]) + strings.ToLower(action[1:]) + " " + strconv.Itoa(len(selected)) + " Repositories"
			h3.SetInnerHTML(title)

		}

		tbody := dialog.QuerySelector("table tbody")

		if tbody != nil {

			ids := selected.Keys("")
			sort.Strings(ids)

			html := ""

			for i := 0; i < len(ids); i++ {
				html += view.renderDialogTableRow(ids[i], selected.Get(ids[i]))
			}

			tbody.SetInnerHTML(html)

		}

	}

}

func (view Repositories) renderDialogTableRow(identifier string, action string) string {

	html := ""
	html += "<tr data-id=\"" + identifier + "\" data-action=\"" + action + "\">"
	html += "<td><label data-state=\"waiting\" title=\"waiting...\"></label></td>"
	html += "<td><label>" + identifier + "</label></td>"
	html += "</tr>"

	return html

}

func (view Repositories) renderFooter() {

	footer := view.GetElement("footer")

	if footer != nil {

		clones := view.Selected.Count("clone")
		fixes := view.Selected.Count("fix")
		commits := view.Selected.Count("commit")
		pulls_or_pushes := view.Selected.Count("pull-or-push")

		total := 0

		for _, owner := range view.Schema.Owners {
			total += len(owner.Repositories)
		}

		message := "Selected " + strconv.Itoa(view.Selected.Length()) + " of " + strconv.Itoa(total) + " Repositories"

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
