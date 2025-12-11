//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac-app/structs"
import "sort"
import "strings"
import "time"

type scheduler_progress struct {
	Start    time.Time `json:"start"`
	Stop     time.Time `json:"Stop"`
	Finished bool      `json:"finished"`
}

type SchedulerTable struct {
	Name       string                      `json:"name"`
	Schema     map[string]string           `json:"schema"`
	Component  *components.Component       `json:"component"`
	Scheduler  *structs.Scheduler          `json:"scheduler"`
	progress   map[string]*scheduler_progress
}

func NewSchedulerTable(name string, schema map[string]string) SchedulerTable {

	var table SchedulerTable

	element := document.CreateElement("table")
	component := components.NewComponent(element)

	table.Schema    = make(map[string]string)
	table.Component = &component
	table.Name      = strings.TrimSpace(strings.ToLower(name))
	table.Scheduler = structs.NewScheduler()
	table.progress  = make(map[string]*scheduler_progress)

	table.SetSchema(schema)

	return table

}

func ToSchedulerTable(element *dom.Element) *SchedulerTable {

	var table SchedulerTable

	component := components.NewComponent(element)

	table.Schema    = make(map[string]string)
	table.Component = &component
	table.Name      = ""
	table.Scheduler = structs.NewScheduler()
	table.progress  = make(map[string]*scheduler_progress)

	return &table

}

func (table *SchedulerTable) Disable() bool {
	return false
}

func (table *SchedulerTable) Enable() bool {
	return false
}

func (table *SchedulerTable) Mount() bool {

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

func (table *SchedulerTable) Render() *dom.Element {

	if table.Component.Element != nil {

		table.Component.Element.SetAttribute("data-name", table.Name)
		table.Component.Element.SetAttribute("data-type", "scheduler")

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

						tr := document.CreateElement("tr")
						tr.SetAttribute("data-id", repository)

						html := ""
						html += "<td>" + repository + "</td>"
						html += "<td>" + action + "</td>"

						progress, ok := table.progress[repository]

						if progress != nil && ok == true {

							if !progress.Start.IsZero() && !progress.Stop.IsZero() && progress.Finished == true {
								// Started, stopped, and finished
								html += "<td><progress data-finished=\"true\" min=\"0\" max=\"100\" value=\"100\">" + formatDuration(progress.Start, progress.Stop) + "</progress></td>"
							} else if !progress.Start.IsZero() && !progress.Stop.IsZero() && progress.Finished == false {
								// Started, stopped, and errored
								html += "<td><progress data-finished=\"false\" min=\"0\" max=\"100\" value=\"100\">" + formatDuration(progress.Start, progress.Stop) + "</progress></td>"
							} else if !progress.Start.IsZero() && progress.Stop.IsZero() {
								// Started
								html += "<td><progress>" + formatDuration(progress.Start, progress.Stop) + "</progress></td>"
							} else {
								// Not started
								html += "<td><progress disabled></progress></td>"
							}

						} else {
							// Not started
							html += "<td><progress disabled></progress></td>"
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

func (table *SchedulerTable) Reset() {

	table.Schema   = make(map[string]string)
	table.Scheduler.Reset()
	table.progress = make(map[string]*scheduler_progress)

	table.Render()

}

func (table *SchedulerTable) Query(query string) interfaces.Component {

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

func (table *SchedulerTable) SetSchema(schema map[string]string) {

	if len(schema) > 0 {

		table.Schema   = schema
		table.progress = make(map[string]*scheduler_progress)

		for repository, action := range table.Schema {

			if strings.Contains(repository, "/") {

				repo := strings.TrimSpace(repository[strings.LastIndex(repository, "/")+1:])
				owner := strings.TrimSpace(repository[0:strings.Index(repository, "/")])

				if owner != "" && repo != "" {

					table.Scheduler.Add(action, owner, repo)

					table.progress[owner + "/" + repo] = &scheduler_progress{
						Start:    time.Time{},
						Stop:     time.Time{},
						Finished: false,
					}

				}

			}


		}

	}

}

func (table *SchedulerTable) Start() {

	for repository, _ := range table.Schema {

		progress, ok := table.progress[repository]

		if ok == true {

			if progress.Start.IsZero() {

				progress.Start = time.Now()
				progress.Stop = time.Time{}
				progress.Finished = false

			}

		}

	}

	table.Render()
	table.Scheduler.Start()

	for action := range table.Scheduler.Results {

		if action.Error != nil {

			progress, ok := table.progress[action.Owner + "/" + action.Repository]

			if ok == true {
				progress.Stop = time.Now()
				progress.Finished = false
			}

		} else if action.Response != nil {

			progress, ok := table.progress[action.Owner + "/" + action.Repository]

			if ok == true {
				progress.Stop = time.Now()
				progress.Finished = true
			}

		}

		table.Render()

	}

}

func (table *SchedulerTable) Stop() {

	table.Scheduler.Stop()

	for repository, _ := range table.Schema {

		progress, ok := table.progress[repository]

		if ok == true {

			if progress.Start.IsZero() {
				progress.Start = time.Now()
				progress.Stop = time.Now()
				progress.Finished = false
			} else {
				progress.Stop = time.Now()
				progress.Finished = false
			}

		}

	}

	table.Render()

}

func (table *SchedulerTable) String() string {

	html := "<table"
	html += " data-name=\"" + table.Name + "\""
	html += " data-type=\"scheduler\""
	html += ">"

	html += "<thead>"
	html += "<tr>"
	html += "<th>Repository</th>"
	html += "<th>Action</th>"
	html += "<th>Progress</th>"
	html += "</tr>"
	html += "</thead>"

	html += "<tbody>"

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

				html += "<tr data-id=\"" + repository + "\">"
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

				html += "</tr>"

			}

		}

	}

	html += "</tbody>"
	html += "</table>"

	return html

}

func (table *SchedulerTable) Unmount() bool {

	if table.Component.Element != nil {
		table.Component.Element.RemoveEventListener("click", nil)
	}

	return true

}

func (table *SchedulerTable) Wait() {
	table.Scheduler.Wait()
}
