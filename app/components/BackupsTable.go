//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/schemas"
import "sort"
import "strings"

type BackupsTable struct {
	Name       string                `json:"name"`
	Component  *components.Component `json:"component"`
	Schemas struct {
		Backups      *schemas.Backups      `json:"backups"`
		Repositories *schemas.Repositories `json:"repositories"`
	} `json:"schemas"`
	selected   map[string]bool
}

func ToBackupsTable(element *dom.Element) *BackupsTable {

	var table BackupsTable

	component := components.NewComponent(element)

	table.Component            = &component
	table.Name                 = ""
	table.Schemas.Backups      = nil
	table.Schemas.Repositories = nil
	table.selected             = make(map[string]bool)

	return &table

}

func (table *BackupsTable) Disable() bool {

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

func (table *BackupsTable) Enable() bool {

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

func (table *BackupsTable) Mount() bool {

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

				} else if action == "backup" || action == "restore" {

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

func (table *BackupsTable) Query(query string) interfaces.Component {

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

func (table *BackupsTable) Render() *dom.Element {

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
			html += "<th>Backup</th>"
			html += "<th>Datetime</th>"
			html += "<th>Size</th>"
			html += "<th>Actions</th>"

			tr.SetInnerHTML(html)
			thead.ReplaceChildren([]*dom.Element{
				tr,
			})

		}

		if tbody != nil {

			elements := make([]*dom.Element, 0)

			if table.Schemas.Backups != nil && table.Schemas.Repositories != nil {

				owner_names := make([]string, 0)

				for _, owner := range table.Schemas.Backups.Owners {
					owner_names = append(owner_names, owner.Name)
				}

				for _, owner := range table.Schemas.Repositories.Owners {
					owner_names = append(owner_names, owner.Name)
				}

				sort.Strings(owner_names)

				for _, owner_name := range owner_names {

					repository_names := make([]string, 0)

					_, ok1 := table.Schemas.Backups.Owners[owner_name]

					if ok1 == true {

						for _, backup := range table.Schemas.Backups.Owners[owner_name].Backups {
							repository_names = append(repository_names, backup.Name)
						}

					}

					_, ok2 := table.Schemas.Repositories.Owners[owner_name]

					if ok2 == true {

						for _, repository := range table.Schemas.Repositories.Owners[owner_name].Repositories {
							repository_names = append(repository_names, repository.Name)
						}

					}

					sort.Strings(repository_names)

					for _, repository_name := range repository_names {

						id := owner_name + "/" + repository_name

						// TODO: If has repository and no backup, then render:
						// Repository name, backup name, file size, action is backup and restore

						// TODO: Else If has backup, then render:
						// Repository name, backup name, file size, action is restore

						// TODO: Else if has repository, then render:
						// Repository name, action is backup


						repository := table.Schema.Owners[owner_name].Backups[repository_name]
						// TODO: backup := table.Schema.Owners[owner_name].Backups[repository_name]
						actions := make([]string, 0)

						// TODO: backup action when repository is available
						// TODO: restore action when backup file is available
						if repository.HasRemoteChanges == true {
							actions = append(actions, "<button data-action=\"fix\">Fix</button>")
						} else if repository.HasLocalChanges == true {
							actions = append(actions, "<button data-action=\"commit\">Commit</button>")
						} else {
							actions = append(actions, "<button data-action=\"pull\">Pull</button>")
							actions = append(actions, "<button data-action=\"push\">Push</button>")
						}

						sort.Strings(actions)

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

						html += "<td>" + id + "</td>"
						html += "<td>TODO</td>" // /path/to/tar.gz
						html += "<td>TODO</td>" // formatted time.Time
						html += "<td>TODO</td>" // Filesize if exists (otherwise 0.00 MB)
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



// TODO: Render



