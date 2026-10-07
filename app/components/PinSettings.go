//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/structs"
import "git-evac/types"
import "sort"

// PinSettings is the dialog content used by the Repositories view to manage
// per-repository overrides (pins) for a single repository.
type PinSettings struct {
	Component *components.Component
	OwnerName string
	RepoName  string
	owner     *structs.SettingsOwner
}

func NewPinSettings(owner *structs.SettingsOwner, repo_name string) PinSettings {

	var pin PinSettings

	element := document.CreateElement("div")
	component := components.NewComponent(element)

	pin.Component = &component
	pin.owner = owner

	if owner != nil {
		pin.OwnerName = owner.Name
	}

	pin.RepoName = repo_name

	pin.Build()

	return pin

}

func (pin *PinSettings) Disable() bool {
	return false
}

func (pin *PinSettings) Enable() bool {
	return false
}

func (pin *PinSettings) SetScheduler(scheduler interfaces.Scheduler) {

	if pin.Component != nil {
		pin.Component.SetScheduler(scheduler)
	}

}

func (pin *PinSettings) Mount() bool {
	return pin.Component.Element != nil
}

func (pin *PinSettings) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) == 1 {

		if pin.Component.Element != nil {

			if utils.MatchesQuery(pin.Component.Element, selectors[0]) == true {
				return pin
			}

		}

	}

	return nil

}

func (pin *PinSettings) Render() *dom.Element {
	return pin.Component.Element
}

func (pin *PinSettings) String() string {

	if pin.Component != nil {
		return pin.Component.String()
	}

	return ""

}

func (pin *PinSettings) Unmount() bool {
	return true
}

func (pin *PinSettings) Build() {

	if pin.Component.Element == nil || pin.owner == nil {
		return
	}

	pinned := pin.owner.GetRepository(pin.RepoName)

	remotes := make(map[string]*types.Remote)

	for name, remote := range pin.owner.SnapshotRemotes() {
		remotes[name] = remote
	}

	if pinned != nil {

		for name, remote := range pinned.SnapshotRemotes() {

			if _, ok := remotes[name]; ok == false {
				remotes[name] = remote
			}

		}

	}

	names := make([]string, 0)

	for name := range remotes {
		names = append(names, name)
	}

	sort.Strings(names)

	element := pin.Component.Element
	element.SetInnerHTML("")

	fieldset := document.CreateElement("fieldset")

	legend := document.CreateElement("legend")
	legend.SetInnerHTML("Remotes")
	fieldset.Append(legend)

	for _, name := range names {

		owner_remote := pin.owner.GetRemote(name)
		pinned_remote := (*types.Remote)(nil)

		if pinned != nil {
			pinned_remote = pinned.GetRemote(name)
		}

		include := true
		url := ""

		if pinned != nil {

			if _, ok := pinned.SnapshotRemotes()[name]; ok == true {
				if pinned_remote == nil {
					include = false
				} else {
					url = pinned_remote.GetURL()
				}
			} else if owner_remote != nil {
				url = owner_remote.GetURL()
			}

		} else if owner_remote != nil {
			url = owner_remote.GetURL()
		}

		div := document.CreateElement("div")
		div.SetAttribute("data-remote", name)

		label := document.CreateElement("label")

		checkbox := document.CreateElement("input")
		checkbox.SetAttribute("type", "checkbox")
		checkbox.SetAttribute("data-field", "include")

		if include == true {
			checkbox.SetAttribute("checked", "")
		}

		label.Append(checkbox)

		text := document.CreateElement("span")
		text.SetInnerHTML(name)
		label.Append(text)

		div.Append(label)

		input := document.CreateElement("input")
		input.SetAttribute("type", "text")
		input.SetAttribute("data-field", "url")
		setValue(input, url)
		div.Append(input)

		fieldset.Append(div)

	}

	element.Append(fieldset)

	fieldset_id := document.CreateElement("fieldset")

	legend_id := document.CreateElement("legend")
	legend_id.SetInnerHTML("Identity")
	fieldset_id.Append(legend_id)

	div_id := document.CreateElement("div")
	div_id.SetAttribute("data-identity", "default")

	label_id := document.CreateElement("label")
	label_id.SetInnerHTML("Default Identity")
	div_id.Append(label_id)

	select_id := document.CreateElement("select")
	select_id.SetAttribute("data-field", "default-identity")

	inherit := document.CreateElement("option")
	inherit.SetAttribute("value", "")
	inherit.SetInnerHTML("Inherit (" + pin.owner.DefaultIdentity + ")")
	select_id.Append(inherit)

	for name := range pin.owner.SnapshotIdentities() {

		option := document.CreateElement("option")
		option.SetAttribute("value", name)
		option.SetInnerHTML(name)
		select_id.Append(option)

	}

	selected := ""

	if pinned != nil {
		selected = pinned.GetDefaultIdentity()
	}

	setValue(select_id, selected)

	div_id.Append(select_id)
	fieldset_id.Append(div_id)
	element.Append(fieldset_id)

}

func (pin *PinSettings) GetPinned() *structs.SettingsRepository {

	pinned := structs.NewSettingsRepository(pin.RepoName)

	if pin.Component.Element == nil || pin.owner == nil {
		return pinned
	}

	rows := pin.Component.Element.QuerySelectorAll("fieldset div[data-remote]")

	for _, row := range rows {

		name := row.GetAttribute("data-remote")

		if name == "" {
			continue
		}

		checkbox := row.QuerySelector("input[data-field=\"include\"]")
		input := row.QuerySelector("input[data-field=\"url\"]")

		include := true

		if checkbox != nil && checkbox.Value != nil {
			include = checkbox.Value.Get("checked").Bool()
		}

		if include == false {
			pinned.Remotes[name] = nil
			continue
		}

		url := readValue(input)
		owner_remote := pin.owner.GetRemote(name)

		if owner_remote == nil || owner_remote.GetURL() != url {
			pinned.SetRemote(types.NewRemote(name, url))
		}

	}

	select_element := pin.Component.Element.QuerySelector("select[data-field=\"default-identity\"]")

	if select_element != nil {

		selected := readValue(select_element)

		if selected != "" && selected != pin.owner.DefaultIdentity {
			pinned.DefaultIdentity = selected
		}

	}

	return pinned

}
