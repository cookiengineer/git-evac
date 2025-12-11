//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/schemas"
import "git-evac/structs"
import "slices"
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
		table.Component.Element.SetAttribute("data-type", "backups")

		thead := table.Component.Element.QuerySelector("thead")
		tbody := table.Component.Element.QuerySelector("tbody")

		if thead == nil && tbody == nil {

			table.Component.Element.ReplaceChildren([]*dom.Element{
				document.CreateElement("thead"),
				document.CreateElement("tbody"),
			})

			thead = table.Component.Element.QuerySelector("thead")
			tbody = table.Component.Element.QuerySelector("tbody")

		}

		if thead != nil {

			tr := document.CreateElement("tr")

			html := ""
			html += "<th><input type=\"checkbox\" title=\"Toggle all repositories\" data-action=\"select\"/></th>"
			html += "<th>Repository</th>"
			html += "<th>Backup</th>"
			html += "<th>Time</th>"
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

					if !slices.Contains(owner_names, owner.Name) {
						owner_names = append(owner_names, owner.Name)
					}

				}

				sort.Strings(owner_names)

				for _, owner_name := range owner_names {

					repository_names := make([]string, 0)

					backup_owner, ok1 := table.Schemas.Backups.Owners[owner_name]

					if ok1 == true {

						for _, backup := range table.Schemas.Backups.Owners[owner_name].Backups {
							repository_names = append(repository_names, backup.Name)
						}

					}

					repository_owner, ok2 := table.Schemas.Repositories.Owners[owner_name]

					if ok2 == true {

						for _, repository := range table.Schemas.Repositories.Owners[owner_name].Repositories {

							if !slices.Contains(repository_names, repository.Name) {
								repository_names = append(repository_names, repository.Name)
							}

						}

					}

					sort.Strings(repository_names)

					for _, repository_name := range repository_names {

						id := owner_name + "/" + repository_name

						var backup *structs.Backup = nil

						if backup_owner != nil {
							backup = backup_owner.GetBackup(repository_name)
						}

						var repository *structs.Repository = nil

						if repository_owner != nil {
							repository = repository_owner.GetRepository(repository_name)
						}

						tr := document.CreateElement("tr")
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

						if repository != nil && backup != nil {

							html += "<td><label>" + id + "</label></td>"
							html += "<td><label data-type=\"file\" data-path=\"" + backup.File + "\">" + id + ".tar.gz</label></td>"
							html += "<td><label>" + formatTime(backup.Time) + "</label></td>"
							html += "<td><label>" + formatSize(backup.Size) + "</label></td>"
							html += "<td>" + strings.Join([]string{
								"<button data-action=\"backup\">Backup</button>",
								"<button data-action=\"restore\">Restore</button>",
							}, " ") + "</td>"

						} else if repository != nil && backup == nil {

							html += "<td><label>" + id + "</label></td>"
							html += "<td></td>"
							html += "<td></td>"
							html += "<td></td>"
							html += "<td><button data-action=\"backup\">Backup</button></td>"

						} else if repository == nil && backup != nil {

							html += "<td></td>"
							html += "<td><label data-type=\"file\" data-path=\"" + backup.File + "\">" + id + ".tar.gz</label></td>"
							html += "<td><label>" + formatTime(backup.Time) + "</label></td>"
							html += "<td><label>" + formatSize(backup.Size) + "</label></td>"
							html += "<td><button data-action=\"restore\">Restore</button></td>"

						} else if repository == nil && backup == nil {
							// Should never happen
						}

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

func (table *BackupsTable) Reset() {

	table.Schemas.Backups      = nil
	table.Schemas.Repositories = nil
	table.selected             = make(map[string]bool)

}

func (table *BackupsTable) Deselect(identifiers []string) {

	for _, id := range identifiers {

		_, ok := table.selected[id]

		if ok == true {
			table.selected[id] = false
		}

	}

}

func (table *BackupsTable) Select(identifiers []string) {

	for _, id := range identifiers {

		_, ok := table.selected[id]

		if ok == true {
			table.selected[id] = true
		}

	}

}

func (table *BackupsTable) Selected() map[string]any {

	result := make(map[string]any)

	if table.Schemas.Backups != nil && table.Schemas.Repositories != nil {

		for id, is_selected := range table.selected {

			if is_selected == true {

				id_owner      := id[0:strings.Index(id, "/")]
				id_repository := id[strings.Index(id, "/")+1:]

				var backup *structs.Backup = nil

				if backup_owner, ok := table.Schemas.Backups.Owners[id_owner]; ok == true {
					backup = backup_owner.GetBackup(id_repository)
				}

				var repository *structs.Repository = nil

				if repository_owner, ok := table.Schemas.Repositories.Owners[id_owner]; ok == true {
					repository = repository_owner.GetRepository(id_repository)
				}

				if repository != nil && backup != nil {
					result[id] = "backup-or-restore"
				} else if repository != nil && backup == nil {
					result[id] = "backup"
				} else if repository == nil && backup != nil {
					result[id] = "restore"
				} else if repository == nil && backup == nil {
					// No action
				}

			}

		}

	}

	return result

}

func (table *BackupsTable) SetSchema(schema1 *schemas.Backups, schema2 *schemas.Repositories) bool {

	if schema1 != nil && schema2 != nil {

		table.Schemas.Backups = schema1
		table.Schemas.Repositories = schema2
		table.selected = make(map[string]bool)

		for _, owner := range table.Schemas.Backups.Owners {

			for _, backup := range owner.Backups {
				table.selected[owner.Name + "/" + backup.Name] = false
			}

		}

		for _, owner := range table.Schemas.Repositories.Owners {

			for _, repository := range owner.Repositories {
				table.selected[owner.Name + "/" + repository.Name] = false
			}

		}

		return true

	}

	return false

}

func (table *BackupsTable) String() string {

	html := "<table"
	html += " data-name=\"" + table.Name + "\""
	html += " data-type=\"backups\""
	html += ">"

	html += "<thead>"
	html += "<tr>"
	html += "<th><input type=\"checkbox\" title=\"Toggle all repositories\" data-action=\"select\"/></th>"
	html += "<th>Repository</th>"
	html += "<th>Backup</th>"
	html += "<th>Time</th>"
	html += "<th>Size</th>"
	html += "<th>Actions</th>"
	html += "</tr>"
	html += "</thead>"

	html += "<tbody>"

	if table.Schemas.Backups != nil && table.Schemas.Repositories != nil {

		owner_names := make([]string, 0)

		for _, owner := range table.Schemas.Backups.Owners {
			owner_names = append(owner_names, owner.Name)
		}

		for _, owner := range table.Schemas.Repositories.Owners {

			if !slices.Contains(owner_names, owner.Name) {
				owner_names = append(owner_names, owner.Name)
			}

		}

		sort.Strings(owner_names)

		for _, owner_name := range owner_names {

			repository_names := make([]string, 0)

			backup_owner, ok1 := table.Schemas.Backups.Owners[owner_name]

			if ok1 == true {

				for _, backup := range table.Schemas.Backups.Owners[owner_name].Backups {
					repository_names = append(repository_names, backup.Name)
				}

			}

			repository_owner, ok2 := table.Schemas.Repositories.Owners[owner_name]

			if ok2 == true {

				for _, repository := range table.Schemas.Repositories.Owners[owner_name].Repositories {

					if !slices.Contains(repository_names, repository.Name) {
						repository_names = append(repository_names, repository.Name)
					}

				}

			}

			sort.Strings(repository_names)

			for _, repository_name := range repository_names {

				id := owner_name + "/" + repository_name

				var backup *structs.Backup = nil

				if backup_owner != nil {
					backup = backup_owner.GetBackup(repository_name)
				}

				var repository *structs.Repository = nil

				if repository_owner != nil {
					repository = repository_owner.GetRepository(repository_name)
				}

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

				if repository != nil && backup != nil {

					html += "<td><label>" + id + "</label></td>"
					html += "<td><label data-type=\"file\" data-path=\"" + backup.File + "\">" + id + ".tar.gz</label></td>"
					html += "<td><label>" + formatTime(backup.Time) + "</label></td>"
					html += "<td><label>" + formatSize(backup.Size) + "</label></td>"
					html += "<td>" + strings.Join([]string{
						"<button data-action=\"backup\">Backup</button>",
						"<button data-action=\"restore\">Restore</button>",
					}, " ") + "</td>"

				} else if repository != nil && backup == nil {

					html += "<td><label>" + id + "</label></td>"
					html += "<td></td>"
					html += "<td></td>"
					html += "<td></td>"
					html += "<td><button data-action=\"backup\">Backup</button></td>"

				} else if repository == nil && backup != nil {

					html += "<td></td>"
					html += "<td><label data-type=\"file\" data-path=\"" + backup.File + "\">" + id + ".tar.gz</label></td>"
					html += "<td><label>" + formatTime(backup.Time) + "</label></td>"
					html += "<td><label>" + formatSize(backup.Size) + "</label></td>"
					html += "<td><button data-action=\"restore\">Restore</button></td>"

				} else if repository == nil && backup == nil {
					// Should never happen
				}

				html += "</tr>"

			}

		}

	}

	html += "</tbody>"
	html += "</table>"

	return html

}

func (table *BackupsTable) Unmount() bool {

	if table.Component.Element != nil {
		table.Component.Element.RemoveEventListener("click", nil)
	}

	return true

}
