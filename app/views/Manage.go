package views

import "gooey"
import "gooey/app"
import "gooey/dom"
import "git-evac-app/api"
import app_schemas "git-evac-app/schemas"
import "git-evac/server/schemas"
import "git-evac/structs"
import "slices"
import "sort"
import "strconv"
import "strings"

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

				selected := app_schemas.Selected{}
				selected[id] = action

				view.Main.Storage.Write("selected-batch", selected)
				view.renderDialog()

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

				selected := app_schemas.Selected{}
				view.Main.Storage.Read("selected-batch", &selected)

				if action == "confirm" {

					buttons := dialog.QuerySelectorAll("footer button[data-action]")

					for b := 0; b < len(buttons); b++ {
						buttons[b].SetAttribute("disabled", "")
					}

					for id, action := range selected {

						label      := dialog.QuerySelector("table tbody tr[data-id=\"" + id + "\"] label[data-state]")
						owner      := id[0:strings.Index(id, "/")]
						repository := id[strings.Index(id, "/")+1:]

						if action == "fix" {

							go func(label *dom.Element, owner string, repository string) {

								response, err := api.TerminalOpen(owner, repository)

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

								response, err := api.GitClone(owner, repository)

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

								response, err := api.GitCommit(owner, repository)

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

								response, err := api.GitPull(owner, repository)

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

								response, err := api.GitPush(owner, repository)

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
				selected := app_schemas.Selected{}

				view.Main.Storage.Read("selected", &selected)

				if action != "" {

					if action == "pull" {

						selected.Filter("pull-or-push")

						for name := range selected {
							selected.Set(name, "pull")
						}

					} else if action == "push" {

						selected.Filter("pull-or-push")

						for name := range selected {
							selected.Set(name, "push")
						}

					} else {
						selected.Filter(action)
					}

					view.Main.Storage.Write("selected-batch", selected)
					view.renderDialog()

					dialog.SetAttribute("open", "")

				}

			}

		}))

	}

}

func (view Manage) Enter() bool {

	schema, err := api.Repositories()

	if err == nil {
		view.Main.Storage.Write("repositories", schema)
	}

	view.renderTable()

	return true

}

func (view Manage) Leave() bool {
	return true
}

func (view Manage) Refresh() {

	schema, err := api.Repositories()

	if err == nil {
		view.Main.Storage.Write("repositories", schema)
	}

	selected := app_schemas.Selected{}
	view.Main.Storage.Read("selected", &selected)

	selected_batch := app_schemas.Selected{}
	view.Main.Storage.Write("selected-batch", selected_batch)

	view.renderTable()
	view.renderFooter()
	view.renderDialog()

}

func (view Manage) Update() {

	selected := app_schemas.Selected{}
	table    := view.GetElement("table")

	view.Main.Storage.Read("selected", &selected)

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

		view.Main.Storage.Write("selected", selected)
		view.renderFooter()

	}

}

func (view Manage) UpdateRepository(updated schemas.Repository) {

	repositories := schemas.Repositories{}

	view.Main.Storage.Read("repositories", &repositories)

	found := false

	for owner_name, owner := range repositories.Owners {

		for repo_name, repo := range owner.Repositories {

			if repo.Folder == updated.Repository.Folder {
				repositories.Owners[owner_name].Repositories[repo_name] = &updated.Repository
				found = true
				break
			}

		}

	}

	if found == true {
		view.Main.Storage.Write("repositories", repositories)
	}

	view.renderTable()

}

func (view Manage) renderTable() {

	schema := schemas.Repositories{}
	table  := view.GetElement("table")

	view.Main.Storage.Read("repositories", &schema)

	if table != nil {

		html := ""

		owners := make([]string, 0)

		for name := range schema.Owners {
			owners = append(owners, name)
		}

		sort.Strings(owners)

		for _, owner := range owners {

			repositories := make([]string, 0)

			for name := range schema.Owners[owner].Repositories {
				repositories = append(repositories, name)
			}

			sort.Strings(repositories)

			for _, repo := range repositories {
				html += view.renderTableRow(owner, schema.Owners[owner].Repositories[repo])
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

func (view Manage) renderDialog() {

	dialog   := view.GetElement("dialog")
	selected := app_schemas.Selected{}

	view.Main.Storage.Read("selected-batch", &selected)

	if dialog != nil {

		ids := make([]string, 0)
		actions := make([]string, 0)

		for id, action := range selected {

			ids = append(ids, id)

			if !slices.Contains(actions, action) {
				actions = append(actions, action)
			}

		}

		sort.Strings(ids)

		h3 := dialog.QuerySelector("h3")

		if h3 != nil {

			title := ""

			if len(actions) == 1 {
				title = strings.ToUpper(actions[0][0:1]) + strings.ToLower(actions[0][1:]) + " " + strconv.Itoa(len(selected)) + " Repositories"
			} else {
				title = "Manage " + strconv.Itoa(len(selected)) + " Repositories"
			}

			h3.SetInnerHTML(title)

		}

		tbody := dialog.QuerySelector("table tbody")

		if tbody != nil {

			html := ""

			for i := 0; i < len(ids); i++ {
				html += view.renderDialogTableRow(ids[i], selected[ids[i]])
			}

			tbody.SetInnerHTML(html)

		}

	}

}

func (view Manage) renderDialogTableRow(identifier string, action string) string {

	html := ""

	html += "<tr data-id=\"" + identifier + "\">"
	html += "<td><label data-state=\"waiting\" title=\"waiting...\"></label></td>"
	html += "<td><label>" + identifier + "</label></td>"
	html += "</tr>"

	return html

}

func (view Manage) renderFooter() {

	repositories := schemas.Repositories{}
	selected     := app_schemas.Selected{}
	footer       := view.GetElement("footer")

	view.Main.Storage.Read("repositories", &repositories)
	view.Main.Storage.Read("selected", &selected)

	if footer != nil {

		clones := selected.Count("clone")
		fixes := selected.Count("fix")
		commits := selected.Count("commit")
		pulls_or_pushes := selected.Count("pull-or-push")

		total := 0

		for _, owner := range repositories.Owners {
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
