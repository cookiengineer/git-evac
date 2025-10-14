//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/schemas"
import app_schemas "git-evac-app/schemas"
import "sort"
import "strconv"
import "strings"

type RepositoriesTable struct {
	Name       string                `json:"name"`
	Schema     *schemas.Repositories `json:"schema"`
	Component  *components.Component `json:"component"`
	selected   map[string]bool
}

func ToRepositoriesTable(element *dom.Element) *RepositoriesTable {

	var table RepositoriesTable

	component := components.NewComponent(element)

	table.Schema    = nil
	table.Component = &component
	table.Name      = ""
	table.selected  = make(map[string]bool)

	table.Mount()

	return &table

}

func (table *RepositoriesTable) Disable() bool {

	var result bool

	inputs := table.Component.Element.QuerySelectorAll("input[type=\"checkbox\"]")

	if len(inputs) > 0 {

		for _, element := range inputs {
			element.SetAttribute("disabled", "")
		}

		result = true

	}

	return result

}

func (table *RepositoriesTable) Enable() bool {

	var result bool

	inputs := table.Component.Element.QuerySelectorAll("input[type=\"checkbox\"]")

	if len(inputs) > 0 {

		for _, element := range inputs {
			element.RemoveAttribute("disabled")
		}

		result = true

	}

	return result

}

func (table *RepositoriesTable) Mount() bool {

	if table.Component != nil {
		table.Component.InitEvent("action")
	}

	if table.Component.Element != nil {

		name := table.Component.Element.GetAttribute("data-name")

		if name != "" {
			table.Name = strings.TrimSpace(strings.ToLower(name))
		}

		table.Component.Element.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {

			if event.Target != nil {

				action := event.Target.GetAttribute("data-action")

				if action == "select" {

					th := event.Target.QueryParent("th")

					if th != nil {

						is_active := event.Target.Value.Get("checked").Bool()

						if is_active == true {

							for identifier, _ := range table.selected {
								table.selected[identifier] = true
							}

							table.Render()

						} else {

							for identifier, _ := range table.selected {
								table.selected[identifier] = false
							}

							table.Render()

						}

					} else {

						is_active  := event.Target.Value.Get("checked").Bool()
						identifier := event.Target.QueryParent("tr").GetAttribute("data-id")

						if is_active == true {

							table.selected[identifier] = true
							table.Render()

						} else {

							input := table.Component.Element.QuerySelector("thead input[data-action=\"select\"]")

							if input != nil {
								input.Value.Set("checked", false)
							}

							table.selected[identifier] = false
							table.Render()

						}

						event.PreventDefault()
						event.StopPropagation()

					}

				}

			}

		}))

		return true

	} else {
		return false
	}

}

func (table *RepositoriesTable) Render() *dom.Element {

	if table.Component.Element != nil {

		if table.Name != "" {
			table.Component.Element.SetAttribute("data-name", table.Name)
		}

		tbody := table.Component.Element.QuerySelector("tbody")

		if tbody != nil {

			elements := make([]*dom.Element, 0)

			if table.Schema != nil {

				for _, owner := range table.Schema.Owners {

					for _, repository := range owner.Repositories {

						id := owner.Name + "/" + repository.Name
						actions := make([]string, 0)
						branches := make([]string, 0)
						remotes := make([]string, 0)

						for _, branch_name := range repository.Branches {
							branches = append(branches, "<label>" + branch_name + "</label>")
						}

						for remote_name, _ := range repository.Remotes {
							remotes = append(remotes, "<label>" + remote_name + "</label>")
						}

						if repository.HasRemoteChanges == true {
							actions = append(actions, "<button data-action=\"fix\">Fix</button>")
						} else if repository.HasLocalChanges == true {
							actions = append(actions, "<button data-action=\"commit\">Commit</button>")
						} else {
							actions = append(actions, "<button data-action=\"pull\">Pull</button>")
							actions = append(actions, "<button data-action=\"push\">Push</button>")
						}

						sort.Strings(actions)
						sort.Strings(branches)
						sort.Strings(remotes)

						tr := dom.Document.CreateElement("tr")
						tr.SetAttribute("data-id", id)

						if table.selected[id] == true {
							tr.SetAttribute("data-select", "true")
						}

						html := ""

						if table.selected[id] == true {
							html += "<td><input type=\"checkbox\" data-action=\"select\" checked/></td>"
						} else {
							html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
						}

						html += "<td>" + owner.Name + "/" + repository.Name + "</td>"
						html += "<td>" + strings.Join(remotes, " ") + "</td>"
						html += "<td>" + strings.Join(branches, " ") + "</td>"
						html += "<td>" + strings.Join(actions, " ") + "</td>"

						tr.SetInnerHTML(html)
						elements = append(elements, tr)

					}

				}

			}

			tbody.ReplaceChildren(elements)

		}

	}

	return table.Component.Element

}

func (table *RepositoriesTable) Deselect(names []string) {

	for _, name := range names {

		_, ok := table.selected[name]

		if ok == true {
			table.selected[name] = false
		}

	}

}

func (table *RepositoriesTable) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) == 1 {

		if table.Component.Element != nil {

			if utils.MatchesQuery(table.Component.Element, selectors[0]) == true {
				return table
			}

		}

	}

	return nil

}

func (table *RepositoriesTable) Select(identifiers []string) {

	for _, id := range identifiers {

		_, ok := table.selected[id]

		if ok == true {
			table.selected[id] = true
		}

	}

}

func (table *RepositoriesTable) Selected() app_schemas.Selected {

	result := app_schemas.Selected(map[string]string{})

	if table.Schema != nil {

		for id, is_selected := range table.selected {

			if is_selected == true {

				id_owner      := id[0:strings.Index(id, "/")]
				id_repository := id[strings.Index(id, "/")+1:]

				_, ok1 := table.Schema.Owners[id_owner]

				if ok1 == true {

					repository, ok2 := table.Schema.Owners[id_owner].Repositories[id_repository]

					if ok2 == true {

						action := ""

						if repository.HasRemoteChanges == true {
							action = "fix"
						} else if repository.HasLocalChanges == true {
							action = "commit"
						} else {
							action = "pull-or-push"
						}

						if action != "" {
							result.Set(id, action)
						}

					}

				}

			}

		}

	}

	return result

}

func (table *RepositoriesTable) SetSchema(schema *schemas.Repositories) {

	if schema != nil && len(schema.Owners) > 0 {

		table.Schema = schema
		table.selected = make(map[string]bool)

		for _, owner := range table.Schema.Owners {

			for _, repository := range owner.Repositories {
				table.selected[owner.Name + "/" + repository.Name] = false
			}

		}

	}

}

func (table *RepositoriesTable) String() string {

	html := "<table"

	if table.Name != "" {
		html += " data-name=\"" + table.Name + "\""
	}

	html += ">"

	html += "<thead>"
	html += "<tr>"
	html += "<th><input type=\"checkbox\" title=\"Toggle all repositories\" data-action=\"select\"/></th>"
	html += "<th>Repository</th>"
	html += "<th>Remotes</th>"
	html += "<th>Branches</th>"
	html += "<th>Actions</th>"

	html += "</tr>"
	html += "</thead>"

	html += "<tbody>"

	for _, position := range table.sorted {

		html += "<tr data-id=\"" + strconv.FormatInt(int64(position), 10) + "\""

		if table.selected[position] == true {
			html += " data-select=\"true\""
		}

		html += ">"

		if table.selected[position] == true {
			html += "<td><input type=\"checkbox\" data-action=\"select\" checked/></td>"
		} else {
			html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
		}

		values, _ := table.Dataset.Get(position).String()

		for _, property := range table.Properties {

			val, ok := values[property]

			if ok == true {
				html += "<td>" + val + "</td>"
			} else {
				html += "<td></td>"
			}

		}

		html += "</tr>"

	}

	html += "</tbody>"
	html += "</table>"

	return html

}

func (table *RepositoriesTable) Unmount() bool {

	if table.Component.Element != nil {
		table.Component.Element.RemoveEventListener("click", nil)
	}

	return true

}
