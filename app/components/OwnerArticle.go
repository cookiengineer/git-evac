//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/structs"
import "git-evac/types"
import "sort"
import "strconv"
import "strings"

type OwnerArticle struct {
	Name      string
	Component *components.Component
	scheduler interfaces.Scheduler
}

func ToOwnerArticle(element *dom.Element) *OwnerArticle {

	var article OwnerArticle

	component := components.NewComponent(element)

	article.Component = &component
	article.Name = strings.TrimSpace(element.GetAttribute("data-name"))

	return &article

}

func (article *OwnerArticle) Disable() bool {
	return false
}

func (article *OwnerArticle) Enable() bool {
	return false
}

func (article *OwnerArticle) SetScheduler(scheduler interfaces.Scheduler) {

	article.scheduler = scheduler

	if article.Component != nil {
		article.Component.SetScheduler(scheduler)
	}

}

func (article *OwnerArticle) Mount() bool {
	return article.Component.Element != nil
}

func (article *OwnerArticle) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) == 1 {

		if article.Component.Element != nil {

			if utils.MatchesQuery(article.Component.Element, selectors[0]) == true {
				return article
			}

		}

	}

	return nil

}

func (article *OwnerArticle) Render() *dom.Element {
	return article.Component.Element
}

func (article *OwnerArticle) String() string {

	if article.Component != nil {
		return article.Component.String()
	}

	return ""

}

func (article *OwnerArticle) Unmount() bool {
	return true
}

func (article *OwnerArticle) GetName() string {

	if article.Component.Element != nil {
		article.Name = strings.TrimSpace(article.Component.Element.GetAttribute("data-name"))
	}

	return article.Name

}

func (article *OwnerArticle) SetOwner(owner *structs.SettingsOwner) {

	if owner == nil || article.Component.Element == nil {
		return
	}

	article.Name = owner.Name

	element := article.Component.Element
	element.SetAttribute("data-name", owner.Name)

	children := make([]*dom.Element, 0)
	children = append(children, buildOwnerHeader(owner.Name))
	children = append(children, buildIdentitiesSection(owner))
	children = append(children, buildRemotesSection(owner))
	children = append(children, buildServicesSection(owner))

	element.ReplaceChildren(children)

}

func (article *OwnerArticle) GetOwner() *structs.SettingsOwner {

	if article.Component.Element == nil {
		return nil
	}

	element := article.Component.Element
	owner := structs.NewSettingsOwner(strings.TrimSpace(element.GetAttribute("data-name")))

	if select_element := element.QuerySelector("section[data-type=\"identities\"] select[data-field=\"default-identity\"]"); select_element != nil {
		owner.DefaultIdentity = readValue(select_element)
	}

	identities := element.QuerySelectorAll("section[data-type=\"identities\"] fieldset[data-type=\"identity\"]")

	for _, fieldset := range identities {

		identity := readIdentityFieldset(fieldset)

		if identity != nil {
			owner.SetIdentity(identity)
		}

	}

	remotes := element.QuerySelectorAll("section[data-type=\"remotes\"] fieldset[data-type=\"remote\"]")

	for _, fieldset := range remotes {

		remote := readRemoteFieldset(fieldset)

		if remote != nil {
			owner.SetRemote(remote)
		}

	}

	services := element.QuerySelectorAll("section[data-type=\"services\"] fieldset[data-type=\"service\"]")

	for _, fieldset := range services {

		service := readServiceFieldset(fieldset)

		if service != nil {
			owner.SetService(service)
		}

	}

	return owner

}

func readIdentityFieldset(fieldset *dom.Element) *types.Identity {

	name := readFieldValue(fieldset, "name")

	if name == "" {
		return nil
	}

	identity := types.NewIdentity(name)

	identity.SetSSHKey(readFieldValue(fieldset, "ssh-key"))
	identity.SetGitUserName(readFieldValue(fieldset, "git-user-name"))
	identity.SetGitUserEmail(readFieldValue(fieldset, "git-user-email"))

	return identity

}

func readRemoteFieldset(fieldset *dom.Element) *types.Remote {

	name := readFieldValue(fieldset, "name")

	if name == "" {
		return nil
	}

	url := readFieldValue(fieldset, "url")

	return types.NewRemote(name, url)

}

func readServiceFieldset(fieldset *dom.Element) *types.Service {

	name := readFieldValue(fieldset, "name")

	if name == "" {
		return nil
	}

	service := types.NewService(name)

	service.SetURL(readFieldValue(fieldset, "url"))
	service.SetToken(readFieldValue(fieldset, "token"))
	service.SetType(readFieldValue(fieldset, "type"))

	return service

}

func buildOwnerHeader(name string) *dom.Element {

	header := document.CreateElement("h3")
	header.SetAttribute("data-type", "organization")

	text := document.CreateElement("span")
	text.SetInnerHTML(name)
	header.Append(text)

	button := document.CreateElement("button")
	button.SetAttribute("data-action", "remove-owner")
	button.SetInnerHTML("Remove")
	header.Append(button)

	return header

}

func buildIdentitiesSection(owner *structs.SettingsOwner) *dom.Element {

	section := document.CreateElement("section")
	section.SetAttribute("data-type", "identities")

	names := make([]string, 0)

	for name := range owner.SnapshotIdentities() {
		names = append(names, name)
	}

	sort.Strings(names)

	section.Append(buildSectionHeader(strconv.Itoa(len(names))+" Identities", "add-identity"))
	section.Append(buildDefaultIdentitySelect(names, owner.DefaultIdentity))

	for _, name := range names {

		identity := owner.GetIdentity(name)

		if identity == nil {
			continue
		}

		fieldset := document.CreateElement("fieldset")
		fieldset.SetAttribute("data-type", "identity")
		fieldset.SetAttribute("data-name", name)

		legend := document.CreateElement("legend")
		legend.SetInnerHTML(name)
		legend.Append(buildRemoveButton("remove-identity"))
		fieldset.Append(legend)

		fieldset.Append(buildTextField("Name", "name", "text", identity.GetName()))
		fieldset.Append(buildTextField("SSH Key", "ssh-key", "text", identity.GetSSHKey()))
		fieldset.Append(buildTextField("Git User Name", "git-user-name", "text", identity.GetGitUserName()))
		fieldset.Append(buildTextField("Git User E-Mail", "git-user-email", "email", identity.GetGitUserEmail()))

		section.Append(fieldset)

	}

	return section

}

func buildRemotesSection(owner *structs.SettingsOwner) *dom.Element {

	section := document.CreateElement("section")
	section.SetAttribute("data-type", "remotes")

	names := make([]string, 0)

	for name := range owner.SnapshotRemotes() {
		names = append(names, name)
	}

	sort.Strings(names)

	section.Append(buildSectionHeader(strconv.Itoa(len(names))+" Remotes", "add-remote"))

	for _, name := range names {

		remote := owner.GetRemote(name)

		if remote == nil {
			continue
		}

		fieldset := document.CreateElement("fieldset")
		fieldset.SetAttribute("data-type", "remote")
		fieldset.SetAttribute("data-name", name)

		legend := document.CreateElement("legend")
		legend.SetInnerHTML(name)
		legend.Append(buildRemoveButton("remove-remote"))
		fieldset.Append(legend)

		fieldset.Append(buildTextField("Name", "name", "text", remote.GetName()))
		fieldset.Append(buildTextField("URL", "url", "text", remote.GetURL()))

		section.Append(fieldset)

	}

	return section

}

func buildServicesSection(owner *structs.SettingsOwner) *dom.Element {

	section := document.CreateElement("section")
	section.SetAttribute("data-type", "services")

	names := make([]string, 0)

	for name := range owner.SnapshotServices() {
		names = append(names, name)
	}

	sort.Strings(names)

	section.Append(buildSectionHeader(strconv.Itoa(len(names))+" Services", "add-service"))

	for _, name := range names {

		service := owner.GetService(name)

		if service == nil {
			continue
		}

		fieldset := document.CreateElement("fieldset")
		fieldset.SetAttribute("data-type", "service")
		fieldset.SetAttribute("data-name", name)

		legend := document.CreateElement("legend")
		legend.SetInnerHTML(name)
		legend.Append(buildRemoveButton("remove-service"))
		fieldset.Append(legend)

		fieldset.Append(buildTextField("Name", "name", "text", service.GetName()))
		fieldset.Append(buildTextField("API URL", "url", "text", service.GetURL()))
		fieldset.Append(buildTextField("Token", "token", "password", service.GetToken()))
		fieldset.Append(buildSelectField("Type", "type", []string{"forgejo", "github", "gitea", "gogs", "gitlab"}, service.GetType()))

		section.Append(fieldset)

	}

	return section

}

func buildSectionHeader(label string, action string) *dom.Element {

	header := document.CreateElement("h4")
	header.SetInnerHTML(label)

	button := document.CreateElement("button")
	button.SetAttribute("data-action", action)
	button.SetInnerHTML("Add")
	header.Append(button)

	return header

}

func buildRemoveButton(action string) *dom.Element {

	button := document.CreateElement("button")
	button.SetAttribute("data-action", action)
	button.SetInnerHTML("Remove")

	return button

}

func buildDefaultIdentitySelect(names []string, selected string) *dom.Element {

	return buildSelectField("Default Identity", "default-identity", names, selected)

}

func buildTextField(label string, field string, typ string, value string) *dom.Element {

	div := document.CreateElement("div")

	label_element := document.CreateElement("label")
	label_element.SetInnerHTML(label)
	div.Append(label_element)

	input := document.CreateElement("input")
	input.SetAttribute("type", typ)
	input.SetAttribute("data-field", field)
	setValue(input, value)
	div.Append(input)

	return div

}

func buildSelectField(label string, field string, values []string, selected string) *dom.Element {

	div := document.CreateElement("div")

	label_element := document.CreateElement("label")
	label_element.SetInnerHTML(label)
	div.Append(label_element)

	select_element := document.CreateElement("select")
	select_element.SetAttribute("data-field", field)

	for _, value := range values {

		option := document.CreateElement("option")
		option.SetAttribute("value", value)
		option.SetInnerHTML(value)

		if value == selected {
			option.SetAttribute("selected", "")
		}

		select_element.Append(option)

	}

	setValue(select_element, selected)
	div.Append(select_element)

	return div

}

func readFieldValue(fieldset *dom.Element, field string) string {

	if fieldset == nil {
		return ""
	}

	return readValue(fieldset.QuerySelector("input[data-field=\"" + field + "\"], select[data-field=\"" + field + "\"]"))

}

func readValue(element *dom.Element) string {

	if element == nil || element.Value == nil {
		return ""
	}

	value := element.Value.Get("value")

	if value.IsNull() || value.IsUndefined() {
		return ""
	}

	return strings.TrimSpace(value.String())

}

func setValue(element *dom.Element, value string) {

	if element != nil && element.Value != nil {
		element.Value.Set("value", value)
	}

}
