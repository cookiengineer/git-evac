package components

import "gooey"
import "gooey/dom"
import "git-evac/server/schemas"
import "git-evac/structs"
import "sort"

var Table *dom.Element = nil

func init() {

	element := gooey.Document.QuerySelector("table")

	if element != nil {
		Table = element
	}

}

func InitTable() {

	Table.QuerySelector("thead input[type=\"checkbox\"]").AddEventListener("change", dom.ToEventListener(func(event dom.Event) {

		target := event.Target

		if target.TagName == "INPUT" {

			is_checked := target.Value.Get("checked").Bool()
			rows := Table.QuerySelectorAll("tr[data-id]")

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

			Update()

		}

	}))

	Table.QuerySelector("tbody").AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

		target := event.Target

		if target.TagName == "INPUT" {

			row := target.ParentNode().ParentNode()
			is_checked := target.Value.Get("checked").Bool()

			if is_checked == true {
				row.SetAttribute("data-select", "true")
			} else {
				row.SetAttribute("data-select", "false")
			}

			Update()

		} else if target.TagName == "BUTTON" {

			row := target.ParentNode().ParentNode()
			id := row.GetAttribute("data-id")
			action := target.GetAttribute("data-action")

			actions := make(map[string]string)
			actions[id] = action

			if len(actions) > 0 {
				RenderDialog(actions)
				Dialog.Open()
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

			Update()

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

			Update()

		}

	}))

}

func RenderTable(index *schemas.Index) {

	html := ""

	for user_name, user := range index.Users {

		for _, repo := range user.Repositories {
			html += RenderTableRow(user_name, "user", repo)
		}

	}

	for orga_name, orga := range index.Organizations {

		for _, repo := range orga.Repositories {
			html += RenderTableRow(orga_name, "organization", repo)
		}

	}

	if Table != nil {

		tbody := Table.QuerySelector("tbody")

		if tbody != nil {
			tbody.SetInnerHTML(html)
		}

	}

}

func RenderTableRow(owner string, typ string, repository *structs.Repository) string {

	var result string

	id := owner + "/" + repository.Name

	result += "<tr data-id=\"" + id + "\" data-select=\"false\">";
	result += "<td><input type=\"checkbox\" data-id=\"" + id + "\" name=\"" + id + "\"/></td>";
	result += "<td><label data-type=\"" + typ + "\">" + owner + "/" + repository.Name + "</label></td>";
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
