//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "sort"
import "strings"
import "time"

type ScheduleProgress struct {
	Start    time.Time `json:"start"`
	Stop     time.Time `json:"Stop"`
	Finished bool      `json:"finished"`
}

type ScheduleTable struct {
	Name       string                      `json:"name"`
	Schema     map[string]string           `json:"schema"`
	Component  *components.Component       `json:"component"`
	progress   map[string]*ScheduleProgress
}

func NewScheduleTable(name string, schema map[string]string) ScheduleTable {

	var table ScheduleTable

	element := dom.Document.CreateElement("table")
	component := components.NewComponent(element)

	table.Schema    = make(map[string]string)
	table.Component = &component
	table.Name      = strings.TrimSpace(strings.ToLower(name))
	table.progress  = make(map[string]*ScheduleProgress)

	table.SetSchema(schema)

	return table

}

func ToScheduleTable(element *dom.Element) *ScheduleTable {

	var table ScheduleTable

	component := components.NewComponent(element)

	table.Schema    = make(map[string]string)
	table.Component = &component
	table.Name      = ""
	table.progress  = make(map[string]*ScheduleProgress)

	return &table

}

func (table *ScheduleTable) Disable() bool {
	return false
}

func (table *ScheduleTable) Enable() bool {
	return false
}

func (table *ScheduleTable) Mount() bool {

	if table.Component != nil {

		table.Component.InitEvent("action")

	}

	if table.Component.Element != nil {

		name := table.Component.Element.GetAttribute("data-name")

		if name != "" {
			table.Name = strings.TrimSpace(strings.ToLower(name))
		}

		table.Component.Element.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {

			event.PreventDefault()
			event.StopPropagation()

		}))

		return true

	} else {
		return false
	}

}

func (table *ScheduleTable) Render() *dom.Element {

	if table.Component.Element != nil {

		table.Component.Element.SetAttribute("data-name", table.Name)
		table.Component.Element.SetAttribute("data-type", "schedule")

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
			html += "<th>Repository</th>"
			html += "<th>Action</th>"
			html += "<th>Progress</th>"

			tr.SetInnerHTML(html)
			thead.ReplaceChildren([]*dom.Element{
				tr,
			})

		}

		if tbody != nil {

			elements := make([]*dom.Element, 0)

			if table.Schema != nil {

				for _, action := range []string{
					"backup",
					"restore",
					"clone",
					"fix",
					"commit",
					"pull",
					"push",
				} {

					filtered := make([]string, 0)

					for repo_name, repo_action := range table.Schema {

						if repo_action == action {
							filtered = append(filtered, repo_name)
						}

					}

					sort.Strings(filtered)

					for _, repository := range filtered {

						tr := dom.Document.CreateElement("tr")
						tr.SetAttribute("data-id", repository)

						html := ""
						html += "<td>" + repository + "</td>"
						html += "<td>" + action + "</td>"

						progress, ok := table.progress[repository]

						if progress != nil && ok == true {

							if !progress.Start.IsZero() && !progress.Stop.IsZero() && progress.Finished == true {
								html += "<td><progress data-finished=\"true\" min=\"0\" max=\"100\" value=\"100\">" + formatDuration(progress.Start, progress.Stop) + "</progress></td>"
							} else if !progress.Start.IsZero() && !progress.Stop.IsZero() && progress.Finished == false {
								html += "<td><progress data-finished=\"false\" min=\"0\" max=\"100\" value=\"100\">" + formatDuration(progress.Start, progress.Stop) + "</progress></td>"
							} else if !progress.Start.IsZero() && progress.Stop.IsZero() {
								html += "<td><progress min=\"0\" max=\"100\">" + formatDuration(progress.Start, progress.Stop) + "</progress></td>"
							} else {
								html += "<td><progress min=\"0\" max=\"100\" value=\"0\"></td>"
							}

						} else {
							html += "<td><progress min=\"0\" max=\"100\" value=\"0\"></td>"
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

func (table *ScheduleTable) Reset() {

	table.Schema   = make(map[string]string)
	table.progress = make(map[string]*ScheduleProgress)

	table.Render()

}

func (table *ScheduleTable) Query(query string) interfaces.Component {

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

func (table *ScheduleTable) SetSchema(schema map[string]string) {

	if len(schema) > 0 {

		table.Schema   = schema
		table.progress = make(map[string]*ScheduleProgress)

		for repository, _ := range table.Schema {

			table.progress[repository] = &ScheduleProgress{
				Start:    time.Time{},
				Stop:     time.Time{},
				Finished: false,
			}

		}

	}

}

func (table *ScheduleTable) Start(repository string) bool {

	if repository == "*" || repository == "all" || repository == "any" {

		for repository, _ := range table.Schema {

			progress, ok := table.progress[repository]

			if ok == true {

				if progress.Start.IsZero() {

					progress.Start = time.Now()
					progress.Stop = time.Time{}
					progress.Finished = false

				}

			} else {

				table.progress[repository] = &ScheduleProgress{
					Start: time.Time{},
					Stop: time.Time{},
					Finished: false,
				}

			}

		}

		table.Render()

		return true

	} else {

		progress, ok1 := table.progress[repository]
		_, ok2 := table.Schema[repository]

		if ok1 == true {

			if progress.Start.IsZero() {

				progress.Start = time.Now()
				progress.Stop = time.Time{}
				progress.Finished = false

				table.Render()

				return true

			}

		} else if ok2 == true {

			table.progress[repository] = &ScheduleProgress{
				Start: time.Time{},
				Stop: time.Time{},
				Finished: false,
			}

			return true

		}

	}

	return false

}

func (table *ScheduleTable) Finish(repository string) bool {

	progress, ok := table.progress[repository]

	if ok == true {

		if !progress.Start.IsZero() {

			progress.Stop = time.Now()
			progress.Finished = true

			table.Render()

		} else {

			progress.Start = time.Now()
			progress.Stop = time.Now()
			progress.Finished = true

			table.Render()

		}

		return true

	}

	return false

}

func (table *ScheduleTable) Stop(repository string) bool {

	if repository == "*" || repository == "all" || repository == "any" {

		for repository, _ := range table.Schema {

			progress, ok := table.progress[repository]

			if ok == true {
				progress.Stop = time.Now()
			}

		}

		table.Render()

		return true

	} else {

		progress, ok := table.progress[repository]

		if ok == true {

			progress.Stop = time.Now()

			table.Render()

			return true

		}

	}

	return false

}

func (table *ScheduleTable) String() string {

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

		// owner_names := make([]string, 0)

		// for _, owner := range table.Schema.Owners {
		// 	owner_names = append(owner_names, owner.Name)
		// }

		// sort.Strings(owner_names)

		// for _, owner_name := range owner_names {

		// 	repository_names := make([]string, 0)

		// 	for _, repository := range table.Schema.Owners[owner_name].Repositories {
		// 		repository_names = append(repository_names, repository.Name)
		// 	}

		// 	sort.Strings(repository_names)

		// 	for _, repository_name := range repository_names {

		// 		id := owner_name + "/" + repository_name
		// 		repository := table.Schema.Owners[owner_name].Repositories[repository_name]
		// 		actions := make([]string, 0)
		// 		branches := make([]string, 0)
		// 		remotes := make([]string, 0)

		// 		for _, branch_name := range repository.Branches {
		// 			branches = append(branches, "<label>" + branch_name + "</label>")
		// 		}

		// 		for remote_name, _ := range repository.Remotes {
		// 			remotes = append(remotes, "<label>" + remote_name + "</label>")
		// 		}

		// 		if repository.HasRemoteChanges == true {
		// 			actions = append(actions, "<button data-action=\"fix\">Fix</button>")
		// 		} else if repository.HasLocalChanges == true {
		// 			actions = append(actions, "<button data-action=\"commit\">Commit</button>")
		// 		} else {
		// 			actions = append(actions, "<button data-action=\"pull\">Pull</button>")
		// 			actions = append(actions, "<button data-action=\"push\">Push</button>")
		// 		}

		// 		sort.Strings(actions)
		// 		sort.Strings(branches)
		// 		sort.Strings(remotes)

		// 		html += "<tr data-id=\"" + id + "\""

		// 		if table.selected[id] == true {
		// 			html += " data-select=\"true\""
		// 		}

		// 		html += ">"

		// 		if table.selected[id] == true {
		// 			html += "<td><input type=\"checkbox\" data-action=\"select\" checked/></td>"
		// 		} else {
		// 			html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
		// 		}

		// 		html += "<td>" + id + "</td>"
		// 		html += "<td>" + strings.Join(remotes, " ") + "</td>"
		// 		html += "<td>" + strings.Join(branches, " ") + "</td>"
		// 		html += "<td>" + strings.Join(actions, " ") + "</td>"

		// 		html += "</tr>"

		// 	}

		// }

	}

	html += "</tbody>"
	html += "</table>"

	return html

}

func (table *ScheduleTable) Unmount() bool {

	if table.Component.Element != nil {
		table.Component.Element.RemoveEventListener("click", nil)
	}

	return true

}
