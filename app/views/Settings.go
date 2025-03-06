package views

import "fmt"

import "gooey"
import "gooey/app"
import "gooey/dom"
import "git-evac-app/actions"
// import app_schemas "git-evac-app/schemas"
import "git-evac/server/schemas"
import "git-evac/structs"
// import "slices"
// import "sort"
import "strconv"
import "strings"

var fieldset_identifier int = 0

type Settings struct {
	Main *app.Main `json:"main"`
	app.BaseView
}

func NewSettings(main *app.Main) Settings {

	var view Settings

	view.Main     = main
	view.Elements = make(map[string]*dom.Element)

	view.SetElement("table", gooey.Document.QuerySelector("main table"))

	view.SetElement("settings-backup", gooey.Document.QuerySelector("main input#settings-backup"))
	view.SetElement("settings-folder", gooey.Document.QuerySelector("main input#settings-folder"))
	view.SetElement("settings-port",   gooey.Document.QuerySelector("main input#settings-port"))

	view.SetElement("articles-identities", gooey.Document.QuerySelector("main article#settings-identities"))
	view.SetElement("articles-remotes",    gooey.Document.QuerySelector("main article#settings-remotes"))

	view.SetElement("footer", gooey.Document.QuerySelector("body > footer"))

	view.Init()

	return view

}

func (view Settings) Init() {

	footer     := view.GetElement("footer")
	identities := view.GetElement("articles-identities")
	remotes    := view.GetElement("articles-remotes")

	if footer != nil {

		footer.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

			target := event.Target
			tagname := target.TagName

			if tagname == "BUTTON" {

				action := target.GetAttribute("data-action")

				if action == "cancel" {

					// TODO: Read Settings from backend, render again?
					// TODO: Revert changes

				} else if action == "save" {

					// TODO: Save the settings from to backend

				}

				fmt.Println("TODO: " + action)

			}

		}))

	}

	if identities != nil {

		identities.AddEventListener("change", dom.ToEventListener(func(event dom.Event) {

			target := event.Target
			tagname := target.TagName

			if tagname == "LEGEND" {

				// TODO: If an identity's name has changed
				// - delete old one from settings
				// - create new one from fieldset inputs

				fmt.Println("change event!")
				fmt.Println(event)

			} else if tagname == "INPUT" || tagname == "TEXTAREA" {

				// TODO: Find the legend, get the name, and change the property

				fmt.Println("change event!")
				fmt.Println(event)

			}

		}))

	}

	if remotes != nil {

		remotes.AddEventListener("change", dom.ToEventListener(func(event dom.Event) {

			// TODO: Same as above

		}))

	}

}

func (view Settings) Enter() bool {

	schema, err := actions.ReadSettings()

	if err == nil {
		view.Main.Storage.Write("settings", schema)
	}

	view.Render()

	return true

}

func (view Settings) Leave() bool {
	return true
}

func (view Settings) Render() {

	schema := schemas.Settings{}

	view.Main.Storage.Read("settings", &schema)



	// TODO: Remove This
	identity := structs.IdentitySettings{
		Name: "cookiengineer",
		SSHKey: "/home/cookiengineer/.ssh/id_rsa",
	}
	identity.Git.Core.SSHCommand = "ssh -i \"/home/cookiengineer/.ssh/id_rsa\""
	identity.Git.User.Name = "Cookie Engineer"
	identity.Git.User.Email = "cookiengineer@protonmail.ch"
	schema.Settings.Identities["cookiengineer"] = identity

	// TODO: End of removal



	backup := view.GetElement("settings-backup")
	folder := view.GetElement("settings-folder")
	port   := view.GetElement("settings-port")

	if backup != nil {
		backup.SetAttribute("value", schema.Settings.Backup)
	}

	if folder != nil {
		folder.SetAttribute("value", schema.Settings.Folder)
	}

	if port != nil {
		port.SetAttribute("value", strconv.FormatUint(uint64(schema.Settings.Port), 10))
	}

	identities := view.GetElement("articles-identities")

	if identities != nil {

		html_identities := "<h3 data-type=\"identities\">Identities</h3>"

		for name, identity := range schema.Settings.Identities {
			html_identities += view.renderIdentityFieldset(name, identity)
		}

		html_identities += "<footer>"
		html_identities += "<div></div>"
		html_identities += "<div><button class=\"primary\" data-action=\"create\">Create</button></div>"
		html_identities += "</footer>"

		identities.SetInnerHTML(html_identities)

	}

	remotes := view.GetElement("articles-remotes")

	if remotes != nil {

		html_remotes := "<h3 data-type=\"remotes\">Remotes</h3>"

		for name, remote := range schema.Settings.Remotes {
			html_remotes += view.renderRemoteFieldset(name, remote)
		}

		html_remotes += "<footer>"
		html_remotes += "<div></div>"
		html_remotes += "<div><button class=\"primary\" data-action=\"create\">Create</button></div>"
		html_remotes += "</footer>"

		remotes.SetInnerHTML(html_remotes)

	}

}

func (view Settings) renderIdentityFieldset(name string, identity structs.IdentitySettings) string {

	if name == "" {
		fieldset_identifier++
		name = "new" + strconv.Itoa(fieldset_identifier)
	}

	html := ""
	html += "<fieldset>"
	html += "<legend data-type=\"identity\"><input type=\"text\" placeholder=\"john_doe\" value=\"" + identity.Name + "\" size=\"" + strconv.Itoa(len(identity.Name)) + "\"/></legend>"
	html += "<div>"
	html += "<label for=\"identities-" + name + "-sshkey\" data-type=\"key\">SSH Key</label>"
	html += "<input id=\"identities-" + name + "-sshkey\" type=\"text\" placeholder=\"~/.ssh/id_rsa\" value=\"" + identity.SSHKey + "\"/>"
	html += "</div>"
	html += "<div>"
	html += "<label for=\"identities-" + name + "-git-user-name\" data-type=\"name\">Git User Name</label>"
	html += "<input id=\"identities-" + name + "-git-user-name\" type=\"text\" placeholder=\"John Doe\" value=\"" + identity.Git.User.Name + "\"/>"
	html += "</div>"
	html += "<div>"
	html += "<label for=\"identities-" + name + "-git-user-email\" data-type=\"email\">Git User Email</label>"
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
	html += "<fieldset>"
	html += "<legend data-type=\"remote\" contenteditable>" + remote.Name + "</legend>"

	html += "<div>"
	html += "<label for=\"remotes-" + name + "-url\" data-type=\"url\">URL</label>"
	html += "<input id=\"remotes-" + name + "-url\" type=\"text\" placeholder=\"git@github.com:/{orga}/{repo}.git\" value=\"" + remote.URL + "\"/>"
	html += "</div>"

	html += "<div>"
	html += "<label for=\"remotes-" + name + "-type\" data-type=\"remote-type\">Type</label>"
	html += "<div>"

	remote_types := []string{
		"github",
		"gitlab",
		"gitea",
		"gogs",
		"git",
	}

	for _, typ := range remote_types {

		html += "<input id=\"remotes-" + name + "-type\" name=\"remotes-" + name + "-type\" type=\"radio\" data-remote=\"" + typ + "\" title=\"" + typ + "\" value=\"" + typ + "\""

		if typ == remote.Type {
			html += " checked"
		}

		html += "/>"

	}

	// html += "<input id=\"remotes-" + name + "-type\" name=\"remotes-" + name + "-type\" type=\"radio\" data-remote=\"github\" title=\"github\" value=\"github\"/>"
	// html += "<input id=\"remotes-" + name + "-type\" name=\"remotes-" + name + "-type\" type=\"radio\" data-remote=\"gitlab\" title=\"gitlab\" value=\"gitlab\"/>"
	// html += "<input id=\"remotes-" + name + "-type\" name=\"remotes-" + name + "-type\" type=\"radio\" data-remote=\"gitea\" title=\"gitea\" value=\"gitea\"/>"
	// html += "<input id=\"remotes-" + name + "-type\" name=\"remotes-" + name + "-type\" type=\"radio\" data-remote=\"gogs\" title=\"gogs\" value=\"gogs\"/>"
	// html += "<input id=\"remotes-" + name + "-type\" name=\"remotes-" + name + "-type\" type=\"radio\" data-remote=\"git\" title=\"git\" value=\"git\"/>"

	html += "</div>"
	html += "</div>"

	html += "<div>"
	html += "<label for=\"remotes-" + name + "-owners\" data-type=\"organization\">Owners <abbr title=\"Users or Organizations mirrored on this remote\">?</abbr></label>"
	html += "<textarea id=\"remotes-" + name + "-owners\">"

	if len(remote.Owners) > 0 {
		html += strings.Join(remote.Owners, "\n")
	}

	html += "</textarea>"
	html += "</div>"

	html += "</fieldset>"

	return html

}

func (view Settings) renderFooter() {

	// TODO: Render Cancel and Save button when settings have changed

}
