package views

import "fmt"

import "gooey"
import "gooey/app"
import "gooey/dom"
import "gooey/location"
import "git-evac-app/actions"
import "git-evac/server/schemas"
import "git-evac/structs"
import "sort"
import "strconv"

var fieldset_identifier int = 0

type Settings struct {
	Main   *app.Main         `json:"main"`
	Schema *schemas.Settings `json:"schema"`
	app.BaseView
}

func NewSettings(main *app.Main) Settings {

	var view Settings

	view.Main     = main
	view.Schema   = &schemas.Settings{}
	view.Elements = make(map[string]*dom.Element)

	view.SetElement("table", gooey.Document.QuerySelector("main table"))

	view.SetElement("settings-backup", gooey.Document.QuerySelector("main input#settings-backup"))
	view.SetElement("settings-folder", gooey.Document.QuerySelector("main input#settings-folder"))
	view.SetElement("settings-port",   gooey.Document.QuerySelector("main input#settings-port"))

	view.SetElement("main",   gooey.Document.QuerySelector("main"))
	view.SetElement("dialog", gooey.Document.QuerySelector("body > dialog"))
	view.SetElement("footer", gooey.Document.QuerySelector("body > footer"))

	view.Init()

	return view

}

func (view Settings) Init() {

	main   := view.GetElement("main")
	dialog := view.GetElement("dialog")
	footer := view.GetElement("footer")

	if dialog != nil {

		dialog.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			target := event.Target

			if target.TagName == "BUTTON" {

				action := target.GetAttribute("data-action")

				if action == "confirm" {

					input_name    := dialog.QuerySelector("#organization-name")
					input_remotes := dialog.QuerySelectorAll("input[type=\"checkbox\"]")
					input_origin  := dialog.QuerySelector("#organization-origin")

					name       := ""
					origin     := ""
					identities := make(map[string]structs.IdentitySettings)
					remotes    := make(map[string]structs.RemoteSettings)

					if input_name != nil {
						name = input_name.Value.Get("value").String()
					}

					for _, input := range input_remotes {

						value      := input.Value.Get("value").String()
						is_checked := input.Value.Get("checked").Bool()

						if is_checked == true {

							if value == "bitbucket" {

								remotes["bitbucket"] = structs.RemoteSettings{
									Name: "bitbucket",
									URL:  "git@bitbucket.org:{owner}/{repo}.git",
									Type: "bitbucket",
								}

							} else if value == "github" {

								remotes["github"] = structs.RemoteSettings{
									Name: "github",
									URL:  "git@github.com:{owner}/{repo}.git",
									Type: "github",
								}

							} else if value == "gitlab" {

								remotes["gitlab"] = structs.RemoteSettings{
									Name: "gitlab",
									URL:  "git@gitlab.com:{owner}/{repo}.git",
									Type: "gitlab",
								}

							} else if value == "gitea" {

								remotes["gitea"] = structs.RemoteSettings{
									Name: "gitea",
									URL:  "git@gitea.com:{owner}/{repo}.git",
									Type: "gitea",
								}

							}

						}

					}

					if input_origin != nil {
						origin = input_origin.Value.Get("value").String()
					}

					if name != "" && origin != "" {

						if origin == "bitbucket" {

							remotes["origin"] = structs.RemoteSettings{
								Name: "origin",
								URL:  "git@bitbucket.org:{owner}/{repo}.git",
								Type: "bitbucket",
							}

						} else if origin == "github" {

							remotes["origin"] = structs.RemoteSettings{
								Name: "origin",
								URL:  "git@github.com:{owner}/{repo}.git",
								Type: "github",
							}

						} else if origin == "gitlab" {

							remotes["origin"] = structs.RemoteSettings{
								Name: "origin",
								URL:  "git@gitlab.com:{owner}/{repo}.git",
								Type: "gitlab",
							}

						} else if origin == "gitea" {

							remotes["origin"] = structs.RemoteSettings{
								Name: "origin",
								URL:  "git@gitea.com:{owner}/{repo}.git",
								Type: "gitea",
							}

						} else {

							remotes["origin"] = structs.RemoteSettings{
								Name: "origin",
								URL:  origin,
								Type: "git",
							}

						}

						view.Schema.Settings.Organizations[name] = structs.OrganizationSettings{
							Name:       name,
							Identities: identities,
							Remotes:    remotes,
						}

						view.Render()
						view.updateFooter(true)
						dialog.RemoveAttribute("open")

					}

				} else if action == "close" {
					dialog.RemoveAttribute("open")
				} else if action == "cancel" {
					dialog.RemoveAttribute("open")
				}

			}

		}))

	}

	if footer != nil {

		footer.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			target := event.Target

			if target.TagName == "BUTTON" {

				action := target.GetAttribute("data-action")

				if action == "add-organization" {

					view.renderDialog()
					dialog.SetAttribute("open", "")

				} else if action == "cancel" {

					location.Location.Reload()

				} else if action == "save" {

					// TODO: Save the local app settings into a JSON and send it to the backend
					fmt.Println("TODO: Save Settings")

				}

			}

		}))

	}

	if main != nil {

		main.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			target := event.Target

			if target.TagName == "BUTTON" {

				action := target.GetAttribute("data-action")

				if action == "remove-organization" {

					section      := target.ParentNode()
					organization := section.GetAttribute("data-name")

					_, ok := view.Schema.Settings.Organizations[organization]

					if ok == true {
						delete(view.Schema.Settings.Organizations, organization)
						view.updateFooter(true)
					}

					section.Remove()

				} else if action == "remove-identity" {

					fieldset     := target.ParentNode()
					section      := fieldset.ParentNode().ParentNode()
					input        := fieldset.QuerySelector("legend input")
					organization := section.GetAttribute("data-name")
					identity     := input.Value.Get("value").String()

					_, ok1 := view.Schema.Settings.Organizations[organization]

					if ok1 == true {

						_, ok2 := view.Schema.Settings.Organizations[organization].Identities[identity]

						if ok2 == true {
							delete(view.Schema.Settings.Organizations[organization].Identities, identity)
							view.updateFooter(true)
						}

					}

					fieldset.Remove()

				} else if action == "remove-remote" {

					fieldset     := target.ParentNode()
					section      := fieldset.ParentNode().ParentNode()
					input        := fieldset.QuerySelector("legend input")
					organization := section.GetAttribute("data-name")
					remote       := input.Value.Get("value").String()

					_, ok1 := view.Schema.Settings.Organizations[organization]

					if ok1 == true {

						_, ok2 := view.Schema.Settings.Organizations[organization].Remotes[remote]

						if ok2 == true {
							delete(view.Schema.Settings.Organizations[organization].Remotes, remote)
							view.updateFooter(true)
						}

					}

					fieldset.Remove()

				} else if action == "add-identity" {

					section := target.ParentNode().ParentNode().ParentNode()
					name    := section.GetAttribute("data-name")
					article := section.QuerySelector("article:nth-of-type(1)")

					if article != nil {

						organization, ok := view.Schema.Settings.Organizations[name]

						if ok == true {

							identity := structs.NewIdentitySettings("new-identity")
							organization.SetIdentity(identity)

							fieldset := view.renderIdentityFieldset(identity.Name, identity)
							article.InsertAdjacentHTML("beforeend", fieldset)

						}

					}

				} else if action == "add-remote" {

					section  := target.ParentNode().ParentNode().ParentNode()
					name     := section.GetAttribute("data-name")
					article  := section.QuerySelector("article:nth-of-type(2)")

					if article != nil {

						organization, ok := view.Schema.Settings.Organizations[name]

						if ok == true {

							remote := structs.NewRemoteSettings("new-remote")
							organization.SetRemote(remote)

							fieldset := view.renderRemoteFieldset(remote.Name, remote)
							article.InsertAdjacentHTML("beforeend", fieldset)

						}

					}

				}

			}

		}))

		main.AddEventListener("change", dom.ToEventListener(func(event dom.Event) {

			target := event.Target

			if target.TagName == "INPUT" {

				if target.Id == "settings-backup" {

					value := target.Value.Get("value").String()
					view.Schema.Settings.Backup = value
					view.updateFooter(true)

				} else if target.Id == "settings-folder" {

					value := target.Value.Get("value").String()
					view.Schema.Settings.Folder = value
					view.updateFooter(true)

				} else if target.Id == "settings-port" {

					value    := target.Value.Get("value").String()
					num, err := strconv.ParseUint(value, 10, 16)

					if err == nil && num > 1024 && num < 65535 {
						view.Schema.Settings.Port = uint16(num)
						view.updateFooter(true)
					}

				} else {

					// Rename through changed inputs inside legend elements
					if target.ParentNode().TagName == "LEGEND" {

						fieldset := target.ParentNode().ParentNode()
						typ      := fieldset.GetAttribute("data-type")
						section  := fieldset.ParentNode().ParentNode()

						if typ == "identity" {

							orga_name     := section.GetAttribute("data-name")
							user_name_old := target.Value.Get("defaultValue").String()
							user_name_new := target.Value.Get("value").String()

							organization, ok := view.Schema.Settings.Organizations[orga_name]

							if ok == true {

								identity := organization.Identities[user_name_old]
								identity.Name = user_name_new

								delete(view.Schema.Settings.Organizations[orga_name].Identities, user_name_old)
								view.Schema.Settings.Organizations[orga_name].Identities[user_name_new] = identity

								view.Render()

							}

						} else if typ == "remote" {

							orga_name       := section.GetAttribute("data-name")
							remote_name_old := target.Value.Get("defaultValue").String()
							remote_name_new := target.Value.Get("value").String()

							organization, ok := view.Schema.Settings.Organizations[orga_name]

							if ok == true {

								remote := organization.Remotes[remote_name_old]
								remote.Name = remote_name_new

								delete(view.Schema.Settings.Organizations[orga_name].Remotes, remote_name_old)
								view.Schema.Settings.Organizations[orga_name].Remotes[remote_name_new] = remote

								view.Render()

							}

						}

					} else {

						fieldset := target.ParentNode().ParentNode()
						typ      := fieldset.GetAttribute("data-type")

						if typ == "identity" {

							// TODO: Update Settings Property
							fmt.Println("TODO: Change Settings property")
							fmt.Println(target, target.Value.Get("value").String())

						} else if typ == "remote" {

							// TODO: Update Settings Property
							fmt.Println("TODO: Change Settings property")
							fmt.Println(target, target.Value.Get("value").String())

						}

					}

				}

			}

		}))

	}

}

func (view Settings) Enter() bool {

	schema, err := actions.ReadSettings()

	if err == nil {
		view.Schema.Settings = schema.Settings
		view.Main.Storage.Write("settings", schema)
	}

	view.Render()

	return true

}

func (view Settings) Leave() bool {
	return true
}

func (view Settings) Render() {

	main   := view.GetElement("main")
	backup := view.GetElement("settings-backup")
	folder := view.GetElement("settings-folder")
	port   := view.GetElement("settings-port")

	if backup != nil {
		backup.SetAttribute("value", view.Schema.Settings.Backup)
	}

	if folder != nil {
		folder.SetAttribute("value", view.Schema.Settings.Folder)
	}

	if port != nil {
		port.SetAttribute("value", strconv.FormatUint(uint64(view.Schema.Settings.Port), 10))
	}

	if main != nil {

		sections := main.QuerySelectorAll("section[data-name]")

		for s := 0; s < len(sections); s++ {
			sections[s].Remove()
		}

		html_organizations := ""

		for name, organization := range view.Schema.Settings.Organizations {
			html_organizations += view.renderOrganizationSection(name, organization)
		}

		main.InsertAdjacentHTML("beforeend", html_organizations)

	}

}

func (view Settings) renderOrganizationSection(name string, organization structs.OrganizationSettings) string {

	html := ""

	if name != "" {

		html += "<section data-type=\"organization\" data-name=\"" + name + "\">"
		html += "<h2>" + organization.Name + "</h2>"
		html += "<button data-action=\"remove-organization\"></button>"
		html += "<article data-type=\"identities\">"
		html += "<h3>Identities</h3>"

		identities := make([]string, 0)

		for name, _ := range organization.Identities {
			identities = append(identities, name)
		}

		sort.Strings(identities)

		for _, name := range identities {
			html += view.renderIdentityFieldset(name, organization.Identities[name])
		}

		html += "</article>"
		html += "<article data-type=\"remotes\">"
		html += "<h3>Remotes</h3>"

		remotes := make([]string, 0)

		for name, _ := range organization.Remotes {
			remotes = append(remotes, name)
		}

		sort.Strings(remotes)

		for _, name := range remotes {
			html += view.renderRemoteFieldset(name, organization.Remotes[name])
		}

		html += "</article>"
		html += "<footer>"
		html += "<div></div>"
		html += "<div>"
		html += "<button class=\"primary\" data-action=\"add-identity\">Identity</button>"
		html += "<button class=\"primary\" data-action=\"add-remote\">Remote</button>"
		html += "</div>"
		html += "</footer>"
		html += "</section>"

	}

	return html

}

func (view Settings) renderIdentityFieldset(name string, identity structs.IdentitySettings) string {

	if name == "" {
		fieldset_identifier++
		name = "identity" + strconv.Itoa(fieldset_identifier)
	}

	html := ""
	html += "<fieldset data-type=\"identity\" data-name=\"" + name + "\">"
	html += "<legend>"

	if len(identity.Name) > 0 {
		html += "<input type=\"text\" placeholder=\"john_doe\" value=\"" + identity.Name + "\" size=\"" + strconv.Itoa(len(identity.Name)) + "\"/>"
	} else {
		html += "<input type=\"text\" placeholder=\"john_doe\" value=\"" + identity.Name + "\"/>"
	}

	html += "</legend>"
	html += "<button data-action=\"remove-identity\"></button>"
	html += "<div>"
	html += "<label for=\"identities-" + name + "-sshkey\" data-type=\"key\">SSH Key</label>"
	html += "<input id=\"identities-" + name + "-sshkey\" type=\"text\" placeholder=\"~/.ssh/id_rsa\" value=\"" + identity.SSHKey + "\"/>"
	html += "</div>"
	html += "<div>"
	html += "<label for=\"identities-" + name + "-git-user-name\" data-type=\"name\">User Name</label>"
	html += "<input id=\"identities-" + name + "-git-user-name\" type=\"text\" placeholder=\"John Doe\" value=\"" + identity.Git.User.Name + "\"/>"
	html += "</div>"
	html += "<div>"
	html += "<label for=\"identities-" + name + "-git-user-email\" data-type=\"email\">User Email</label>"
	html += "<input id=\"identities-" + name + "-git-user-email\" type=\"text\" placeholder=\"john.doe@example.com\" value=\"" + identity.Git.User.Email + "\"/>"
	html += "</div>"
	html += "</fieldset>"

	return html

}

func (view Settings) renderRemoteFieldset(name string, remote structs.RemoteSettings) string {

	if name == "" {
		fieldset_identifier++
		name = "new" + strconv.Itoa(fieldset_identifier)
	}

	html := ""
	html += "<fieldset data-type=\"remote\">"
	html += "<legend>"
	html += "<input type=\"text\" value=\"" + remote.Name + "\" size=\"" + strconv.Itoa(len(remote.Name)) + "\"/>"
	html += "</legend>"
	html += "<button data-action=\"remove-remote\"></button>"

	html += "<div>"
	html += "<label for=\"remotes-" + name + "-url\" data-type=\"url\">URL</label>"
	html += "<input id=\"remotes-" + name + "-url\" type=\"text\" placeholder=\"git@github.com:/{orga}/{repo}.git\" value=\"" + remote.URL + "\"/>"
	html += "</div>"

	html += "<div>"
	html += "<label for=\"remotes-" + name + "-type\" data-type=\"remote-type\">Type</label>"
	html += "<div>"

	remote_types := []string{
		"bitbucket",
		"github",
		"gitlab",
		"gitea",
		"gogs",
		"git",
	}

	for _, typ := range remote_types {

		html += "<input name=\"remotes-" + name + "-type\" type=\"radio\" data-remote=\"" + typ + "\" title=\"" + typ + "\" value=\"" + typ + "\""

		if typ == remote.Type {
			html += " checked"
		}

		html += "/>"

	}

	html += "</div>"
	html += "</div>"
	html += "</fieldset>"

	return html

}

func (view Settings) renderDialog() {

	dialog := view.GetElement("dialog")

	if dialog != nil {

		inputs := dialog.QuerySelectorAll("input")

		for i := 0; i < len(inputs); i++ {

			input := inputs[i]
			typ   := input.GetAttribute("type")

			if typ == "checkbox" {
				input.Value.Set("checked", false)
			} else {
				input.Value.Set("value", "")
			}

		}

	}

}

func (view Settings) updateFooter(changed bool) {

	footer := view.GetElement("footer")

	if footer != nil {

		cancel := footer.QuerySelector("button[data-action=\"cancel\"]")
		save   := footer.QuerySelector("button[data-action=\"save\"]")

		if changed == true {
			cancel.RemoveAttribute("disabled")
			save.RemoveAttribute("disabled")
		} else {
			cancel.SetAttribute("disabled", "")
			save.SetAttribute("disabled", "")
		}

	}

}
