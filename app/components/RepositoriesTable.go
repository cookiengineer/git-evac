//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/data"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/server/schemas"
import "strconv"
import "strings"
import "fmt"

type RepositoriesTable struct {
	Name       string                `json:"name"`
	Labels     []string              `json:"labels"`
	Properties []string              `json:"properties"`
	Types      []string              `json:"types"`
	Schema     *schemas.Repositories `json:"schema"`
	Component  *components.Component `json:"component"`
	Selectable bool                  `json:"selectable"`
	selected   []bool
	sorted     []int
	sortby     string
}

func ToRepositoriesTable(element *dom.Element) *RepositoriesTable {

	var table RepositoriesTable

	component := components.NewComponent(element)

	table.Schema     = nil
	table.Component  = &component
	table.Name       = ""
	table.Labels     = make([]string, 0)
	table.Properties = make([]string, 0)
	table.Types      = make([]string, 0)
	table.Selectable = element.HasAttribute("data-selectable")
	table.selected   = make([]bool, 0)
	table.sorted     = make([]int, 0)
	table.sortby     = ""

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

		selectable := table.Component.Element.GetAttribute("data-selectable")

		if selectable == "true" {
			table.Selectable = true
		} else {
			table.Selectable = false
		}

		thead := table.Component.Element.QuerySelector("thead")

		if thead != nil && len(table.Labels) == 0 && len(table.Properties) == 0 && len(table.Types) == 0 {

			elements   := thead.QuerySelectorAll("th")
			labels     := make([]string, 0)
			properties := make([]string, 0)
			types      := make([]string, 0)
			selectable := table.Selectable

			if len(elements) > 0 {

				checkbox := elements[0].QuerySelector("input[type=\"checkbox\"]")

				if checkbox != nil {
					elements   = elements[1:]
					selectable = true
				}

				for e := 0; e < len(elements); e++ {

					element := elements[e]

					label    := strings.TrimSpace(element.TextContent)
					property := element.GetAttribute("data-property")
					typ      := element.GetAttribute("data-type")

					if typ == "" {
						typ = "string"
					}

					if label != "" && property != "" {

						labels     = append(labels, label)
						properties = append(properties, property)
						types      = append(types, typ)

					}

				}

			}

			table.Labels     = labels
			table.Properties = properties
			table.Types      = types
			table.Selectable = selectable

		}

		tbody := table.Component.Element.QuerySelector("tbody")

		if tbody != nil {

			rows     := tbody.QuerySelectorAll("tr")
			dataset  := data.NewDataset(len(rows))
			sorted   := make([]int, len(rows))
			selected := make([]bool, len(rows))

			for r, row := range rows {

				var id int = -1

				id_str := row.GetAttribute("data-id")

				if id_str != "" {

					tmp, err := strconv.ParseInt(id_str, 10, 0)

					if err == nil {
						id = int(tmp)
					}

				} else {
					id = int(r)
				}

				elements := row.QuerySelectorAll("td")

				if len(elements) > 0 {

					checkbox := elements[0].QuerySelector("input[type=\"checkbox\"]")
					values   := make(map[string]string)
					types    := make(map[string]string)

					if checkbox != nil {
						elements = elements[1:]
					}

					for e := 0; e < len(elements); e++ {

						key := table.Properties[e]
						typ := table.Types[e]
						val := strings.TrimSpace(elements[e].TextContent)

						if key != "" && typ != "" && val != "" {
							values[key] = val
							types[key]  = typ
						}

					}

					if len(values) == len(types) {

						if id != -1 && id >= 0 && id < dataset.Length() {
							dataset.Set(id, data.ParseData(values, types))
							selected[id] = row.HasAttribute("data-select")
							sorted[r] = id
						} else {
							dataset.Set(id, data.ParseData(values, types))
							selected[r] = row.HasAttribute("data-select")
							sorted[r] = id
						}

					}

				}

			}

			table.Dataset = &dataset
			table.sorted = sorted
			table.selected = selected

		} else {

			console.Group("Table Body: Invalid Markup")
			console.Error("Expected <tr>...</tr>")
			console.GroupEnd("Table Body: Invalid Markup")

		}

		table.Component.Element.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {

			if event.Target != nil {

				action := event.Target.GetAttribute("data-action")

				if action == "select" {

					th := event.Target.QueryParent("th")

					if th != nil {

						is_active := event.Target.Value.Get("checked").Bool()

						if is_active == true {

							for s := 0; s < len(table.selected); s++ {
								table.selected[s] = true
							}

							table.Render()

						} else {

							for s := 0; s < len(table.selected); s++ {
								table.selected[s] = false
							}

							table.Render()

						}

					} else {

						is_active := event.Target.Value.Get("checked").Bool()
						tmp       := event.Target.QueryParent("tr").GetAttribute("data-id")

						if is_active == true {

							num, err := strconv.ParseInt(tmp, 10, 64)

							if err == nil {

								index := int(num)

								if index >= 0 && index < table.Dataset.Length() {

									table.selected[index] = true
									table.Render()

								}

							}

						} else {

							num, err := strconv.ParseInt(tmp, 10, 64)

							if err == nil {

								index := int(num)

								if index >= 0 && index < table.Dataset.Length() {

									input := table.Component.Element.QuerySelector("thead input[data-action=\"select\"]")

									if input != nil {
										input.Value.Set("checked", false)
									}

									table.selected[index] = false
									table.Render()

								}

							}

						}

						event.PreventDefault()
						event.StopPropagation()

					}

				} else if action == "sort" {

					thead := table.Component.Element.QuerySelector("thead")
					th    := event.Target.QueryParent("th")

					if thead != nil && th != nil {

						property := th.GetAttribute("data-property")
						ths      := thead.QuerySelectorAll("th")

						for _, th := range ths {
							th.RemoveAttribute("data-sort")
						}

						if table.sortby != property {

							th.SetAttribute("data-sort", "ascending")

							table.sorted = table.Dataset.SortByProperty(property)
							table.sortby = property

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

		if table.Selectable == true {
			table.Component.Element.SetAttribute("data-selectable", "true")
		}

		tbody := table.Component.Element.QuerySelector("tbody")

		if tbody != nil {

			elements := make([]*dom.Element, 0)

			for _, position := range table.sorted {

				tr := dom.Document.CreateElement("tr")

				tr.SetAttribute("data-id", strconv.FormatInt(int64(position), 10))

				if table.selected[position] == true {
					tr.SetAttribute("data-select", "true")
				}

				html := ""

				if table.Selectable == true {

					if table.selected[position] == true {
						html += "<td><input type=\"checkbox\" data-action=\"select\" checked/></td>"
					} else {
						html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
					}

				}

				whatever := table.Dataset.Get(position)

				fmt.Println(position, whatever)

				values, _ := table.Dataset.Get(position)

				for _, property := range table.Properties {

					value, ok := values[property]

					if property == "repository" {

						html += "<td>" + value + "</td>"

					} else if property == "remotes" {

						labels := make([]string, 0)

						fmt.Println("remote", value)

						for _, remote := range value {
							labels = append(labels, remote)
						}

						html += "<td>" + labels + "</td>"

					} else if property == "branches" {

					} else if property == "actions" {

					}

					if ok == true {
					} else {
						html += "<td></td>"
					}

				}

				tr.SetInnerHTML(html)

				elements = append(elements, tr)

			}

			tbody.ReplaceChildren(elements)

		}

	}

	return table.Component.Element

}

func (table *RepositoriesTable) Deselect(indexes []int) {

	for _, index := range indexes {
		table.selected[index] = false
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

func (table *RepositoriesTable) Select(indexes []int) {

	for _, index := range indexes {
		table.selected[index] = true
	}

}

func (table *RepositoriesTable) Selected() ([]int, []data.Data) {

	result_indexes := make([]int, 0)
	result_dataset := make([]data.Data, 0)

	for s, value := range table.selected {

		if value == true {

			data := table.Dataset.Get(s)

			if data != nil {
				result_indexes = append(result_indexes, s)
				result_dataset = append(result_dataset, *data)
			}

		}

	}

	return result_indexes, result_dataset

}

func (table *RepositoriesTable) SetSchema(repositories schemas.Repositories) {

	table.Repositories = &repositories

	// TODO: Get repositories.Length() manually

	table.selected = make([]bool, dataset.Length())
	table.sortby = ""
	table.sorted = make([]int, dataset.Length())

	for d := 0; d < dataset.Length(); d++ {
		table.sorted[d] = d
	}

}

func (table *RepositoriesTable) SetLabelsAndPropertiesAndTypes(labels []string, properties []string, types []string) bool {

	var result bool

	if len(labels) == len(properties) && len(labels) == len(types) {

		table.Labels     = labels
		table.Properties = properties
		table.Types      = types

		result = true

	}

	return result

}

func (table *RepositoriesTable) SortBy(prop string) bool {

	var result bool

	thead := table.Component.Element.QuerySelector("thead")

	if thead != nil {

		ths   := thead.QuerySelectorAll("th")
		found := false

		for _, th := range ths {
			th.RemoveAttribute("data-sort")
		}

		for _, th := range ths {

			property := th.GetAttribute("data-property")

			if property == prop {
				th.SetAttribute("data-sort", "ascending")
				found = true
				break
			}

		}

		if found == true {
			table.sorted = table.Dataset.SortByProperty(prop)
			table.sortby = prop
			result = true
		}

	}

	return result

}

func (table *RepositoriesTable) String() string {

	html := "<table"

	if table.Name != "" {
		html += " data-name=\"" + table.Name + "\""
	}

	if table.Selectable == true {
		html += " data-selectable=\"" + strconv.FormatBool(table.Selectable) + "\""
	}

	html += ">"

	html += "<thead>"
	html += "<tr>"

	if table.Selectable == true {
		html += "<th><input type=\"checkbox\" data-action=\"select\"/></th>"
	}

	for l, label := range table.Labels {

		property := table.Properties[l]
		typ      := table.Types[l]

		html += "<th data-property=\"" + property + "\" data-type=\"" + typ + "\""

		if table.sortby == property {
			html += " data-sort=\"ascending\""
		}

		html += ">"
		html += "<label data-action=\"sort\">"
		html += label
		html += "</label>"
		html += "</th>"

	}

	html += "</tr>"
	html += "</thead>"

	html += "<tbody>"

	for _, position := range table.sorted {

		html += "<tr data-id=\"" + strconv.FormatInt(int64(position), 10) + "\""

		if table.selected[position] == true {
			html += " data-select=\"true\""
		}

		html += ">"

		if table.Selectable == true {

			if table.selected[position] == true {
				html += "<td><input type=\"checkbox\" data-action=\"select\" checked/></td>"
			} else {
				html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
			}

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
