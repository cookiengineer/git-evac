//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/schemas"
import "sort"
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

	table.Component = &component
	table.Name      = ""
	table.Schema    = nil
	table.selected  = make(map[string]bool)

	return &table

}

func (table *RepositoriesTable) Disable() bool {

	var result bool

	elements := table.Component.Element.QuerySelectorAll("button, input[type=\"checkbox\"]")

	if len(elements) > 0 {

		for _, element := range elements {
			element.SetAttribute("disabled", "")
		}

		result = true

	}

	return result

}

func (table *RepositoriesTable) Enable() bool {

	var result bool

	elements := table.Component.Element.QuerySelectorAll("button, input[type=\"checkbox\"]")

	if len(elements) > 0 {

		for _, element := range elements {
			element.RemoveAttribute("disabled")
		}

		result = true

	}

	return result

}

func (table *RepositoriesTable) Mount() bool {

	if table.Component != nil {

		table.Component.InitEvent("action")
		table.Component.InitEvent("select")

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

						} else {

							for identifier, _ := range table.selected {
								table.selected[identifier] = false
							}

						}

					} else {

						is_active  := event.Target.Value.Get("checked").Bool()
						identifier := event.Target.QueryParent("tr").GetAttribute("data-id")

						if is_active == true {

							table.selected[identifier] = true

						} else {

							input := table.Component.Element.QuerySelector("thead input[data-action=\"select\"]")

							if input != nil {
								input.Value.Set("checked", false)
							}

							table.selected[identifier] = false

						}

						event.PreventDefault()
						event.StopPropagation()

					}

					table.Render()
					table.Component.FireEventListeners("select", table.Selected())

				} else if action == "fix" || action == "commit" || action == "pull" || action == "push" {

					tr := event.Target.QueryParent("tr")
					id := tr.GetAttribute("data-id")

					if id != "" {

						actions := make(map[string]any)
						actions[id] = action

						table.Component.FireEventListeners("action", actions)

					}

				}

			}

		}))

		return true

	} else {
		return false
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

func (table *RepositoriesTable) Render() *dom.Element {

	if table.Component.Element != nil {

		table.Component.Element.SetAttribute("data-name", table.Name)
		table.Component.Element.SetAttribute("data-type", "repositories")

		thead := table.Component.Element.QuerySelector("thead")
		tbody := table.Component.Element.QuerySelector("tbody")

		if thead == nil && tbody == nil {

			table.Component.Element.ReplaceChildren([]*dom.Element{
				dom.Document.CreateElement("thead"),
				dom.Document.CreateElement("tbody"),
			})

			thead = table.Component.Element.QuerySelector("thead")
			tbody = table.Component.Element.QuerySelector("tbody")

		}

		if thead != nil {

			tr := dom.Document.CreateElement("tr")

			html := ""
			html += "<th><input type=\"checkbox\" title=\"Toggle all repositories\" data-action=\"select\"/></th>"
			html += "<th>Repository</th>"
			html += "<th>Remotes</th>"
			html += "<th>Branches</th>"
			html += "<th>Actions</th>"

			tr.SetInnerHTML(html)
			thead.ReplaceChildren([]*dom.Element{
				tr,
			})

		}

		if tbody != nil {

			elements := make([]*dom.Element, 0)

			if table.Schema != nil {

				owner_names := make([]string, 0)

				for _, owner := range table.Schema.Owners {
					owner_names = append(owner_names, owner.Name)
				}

				sort.Strings(owner_names)

				for _, owner_name := range owner_names {

					repository_names := make([]string, 0)

					for _, repository := range table.Schema.Owners[owner_name].Repositories {
						repository_names = append(repository_names, repository.Name)
					}

					sort.Strings(repository_names)

					for _, repository_name := range repository_names {

						id := owner_name + "/" + repository_name
						repository := table.Schema.Owners[owner_name].Repositories[repository_name]
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

						html += "<td><label>" + id + "</label></td>"
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

func (table *RepositoriesTable) Reset() {

	table.Schema   = nil
	table.selected = make(map[string]bool)

}

func (table *RepositoriesTable) Deselect(identifiers []string) {

	for _, id := range identifiers {

		_, ok := table.selected[id]

		if ok == true {
			table.selected[id] = false
		}

	}

}

func (table *RepositoriesTable) Select(identifiers []string) {

	for _, id := range identifiers {

		_, ok := table.selected[id]

		if ok == true {
			table.selected[id] = true
		}

	}

}

func (table *RepositoriesTable) Selected() map[string]any {

	result := make(map[string]any)

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
							result[id] = action
						}

					}

				}

			}

		}

	}

	return result

}

func (table *RepositoriesTable) SetSchema(schema *schemas.Repositories) bool {

	if schema != nil && len(schema.Owners) > 0 {

		table.Schema = schema
		table.selected = make(map[string]bool)

		for _, owner := range table.Schema.Owners {

			for _, repository := range owner.Repositories {
				table.selected[owner.Name + "/" + repository.Name] = false
			}

		}

		return true

	}

	return false

}

func (table *RepositoriesTable) String() string {

	html := "<table"
	html += " data-name=\"" + table.Name + "\""
	html += " data-type=\"repositories\""
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

	if table.Schema != nil {

		owner_names := make([]string, 0)

		for _, owner := range table.Schema.Owners {
			owner_names = append(owner_names, owner.Name)
		}

		sort.Strings(owner_names)

		for _, owner_name := range owner_names {

			repository_names := make([]string, 0)

			for _, repository := range table.Schema.Owners[owner_name].Repositories {
				repository_names = append(repository_names, repository.Name)
			}

			sort.Strings(repository_names)

			for _, repository_name := range repository_names {

				id := owner_name + "/" + repository_name
				repository := table.Schema.Owners[owner_name].Repositories[repository_name]
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

				html += "<tr data-id=\"" + id + "\""

				if table.selected[id] == true {
					html += " data-select=\"true\""
				}

				html += ">"

				if table.selected[id] == true {
					html += "<td><input type=\"checkbox\" data-action=\"select\" checked/></td>"
				} else {
					html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
				}

				html += "<td>" + id + "</td>"
				html += "<td>" + strings.Join(remotes, " ") + "</td>"
				html += "<td>" + strings.Join(branches, " ") + "</td>"
				html += "<td>" + strings.Join(actions, " ") + "</td>"

				html += "</tr>"

			}

		}

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
