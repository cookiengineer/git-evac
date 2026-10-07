//go:build wasm

package controllers

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/location"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/interfaces"
import ui_components "github.com/cookiengineer/gooey/components/ui"
import "git-evac/schemas"
import "git-evac/structs"
import "git-evac/types"
import app_actions "git-evac-app/actions"
import app_views "git-evac-app/views"
import "strconv"
import "strings"

type Settings struct {
	Main   *app.Main           `json:"main"`
	View   *app_views.Settings `json:"view"`
	Schema *schemas.Settings   `json:"schema"`
}

func NewSettings(main *app.Main, view interfaces.View) *Settings {

	var controller Settings

	controller.Main = main
	controller.View = view.(*app_views.Settings)

	return &controller

}

func (controller *Settings) Enter() bool {

	if controller.Main.Header != nil {

		controller.Main.Header.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

			if event == "action" {

				action, ok := attributes["action"].(string)

				if ok == true && action == "refresh" {
					location.GetLocation().Reload()
				}

			}

		}, false))

	}

	if controller.Main.Footer != nil {

		controller.Main.Footer.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

			if event == "action" {

				action, ok := attributes["action"].(string)

				if ok == true {

					if action == "save" {
						controller.save()
					} else if action == "cancel" {
						location.GetLocation().Reload()
					} else if action == "add-owner" {
						controller.showOwnerDialog()
					}

				}

			}

		}, false))

	}

	if controller.Main.Dialog != nil {

		controller.Main.Dialog.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

			if event == "action" {

				action, ok := attributes["action"].(string)

				if ok == true {

					if action == "confirm" {
						controller.confirmOwnerDialog()
					} else if action == "cancel" || action == "close" {
						controller.hideOwnerDialog()
					}

				}

			}

		}, false))

	}

	if controller.View.Element != nil {

		controller.View.Element.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {

			if event.Target == nil {
				return
			}

			action := event.Target.GetAttribute("data-action")

			if action == "" {
				return
			}

			if action == "add-owner" {
				controller.showOwnerDialog()
			} else if action == "remove-owner" {
				controller.removeOwner(event.Target)
			} else if action == "add-identity" {
				controller.addIdentity(event.Target)
			} else if action == "remove-identity" {
				controller.removeEntry(event.Target, "identities")
			} else if action == "add-remote" {
				controller.addRemote(event.Target)
			} else if action == "remove-remote" {
				controller.removeEntry(event.Target, "remotes")
			} else if action == "add-service" {
				controller.addService(event.Target)
			} else if action == "remove-service" {
				controller.removeEntry(event.Target, "services")
			}

		}))

	}

	go controller.Update()

	return true

}

func (controller *Settings) Leave() bool {

	if controller.Main.Header != nil {
		controller.Main.Header.Component.RemoveEventListener("action", nil)
	}

	if controller.Main.Footer != nil {
		controller.Main.Footer.Component.RemoveEventListener("action", nil)
	}

	if controller.Main.Dialog != nil {
		controller.Main.Dialog.Component.RemoveEventListener("action", nil)
	}

	if controller.View.Element != nil {
		controller.View.Element.RemoveEventListener("click", nil)
	}

	return true

}

func (controller *Settings) Name() string {
	return "settings"
}

func (controller *Settings) Update() {

	if controller.Main != nil {

		schema, err := app_actions.ReadSettings()

		if err == nil && schema != nil {

			controller.Schema = schema
			controller.Main.Storage.Write("settings", schema)
			controller.View.SetSchema(schema)
			controller.setButtons(true)

		}

		controller.Render()

	}

}

func (controller *Settings) Render() {
	controller.View.Render()
}

func (controller *Settings) readInto() {

	if controller.Schema != nil {
		controller.View.ApplyToSchema(controller.Schema)
	}

}

func (controller *Settings) save() {

	if controller.Schema == nil {
		return
	}

	controller.readInto()

	schema, err := app_actions.SaveSettings(*controller.Schema)

	if err == nil && schema != nil {
		controller.Schema = schema
		controller.Main.Storage.Write("settings", schema)
		controller.View.SetSchema(schema)
	}

	controller.setButtons(true)

}

func (controller *Settings) setButtons(enabled bool) {

	if controller.Main.Footer == nil {
		return
	}

	save_button, ok1 := components.UnwrapComponent[*ui_components.Button](controller.Main.Footer.Query("footer > button[data-action=\"save\"]"))
	cancel_button, ok2 := components.UnwrapComponent[*ui_components.Button](controller.Main.Footer.Query("footer > button[data-action=\"cancel\"]"))

	if ok1 == true {
		if enabled == true {
			save_button.Enable()
		} else {
			save_button.Disable()
		}
	}

	if ok2 == true {
		if enabled == true {
			cancel_button.Enable()
		} else {
			cancel_button.Disable()
		}
	}

}

func (controller *Settings) showOwnerDialog() {

	if controller.Main.Dialog == nil {
		return
	}

	name_input := dom.GetDocument().QuerySelector("#organization-name")

	if name_input != nil && name_input.Value != nil {
		name_input.Value.Set("value", "")
	}

	controller.Main.Dialog.SetTitle("Create Organization Settings")
	controller.Main.Dialog.Enable()
	controller.Main.Dialog.Show()

}

func (controller *Settings) hideOwnerDialog() {

	if controller.Main.Dialog != nil {
		controller.Main.Dialog.Disable()
		controller.Main.Dialog.Hide()
	}

}

func (controller *Settings) confirmOwnerDialog() {

	if controller.Schema == nil {
		return
	}

	name_input := dom.GetDocument().QuerySelector("#organization-name")

	name := ""

	if name_input != nil && name_input.Value != nil {
		value := name_input.Value.Get("value")
		if !value.IsNull() && !value.IsUndefined() {
			name = strings.TrimSpace(value.String())
		}
	}

	if name != "" {

		controller.readInto()

		owners := controller.Schema.Settings.GetOwners()

		if _, ok := owners[name]; ok == false {
			owners[name] = structs.NewSettingsOwner(name)
			controller.Schema.Settings.SetOwners(owners)
			controller.View.SetSchema(controller.Schema)
		}

	}

	controller.hideOwnerDialog()

}

func (controller *Settings) removeOwner(target *dom.Element) {

	article := target.QueryParent("article")

	if article == nil || controller.Schema == nil {
		return
	}

	name := article.GetAttribute("data-name")

	if name == "" {
		return
	}

	controller.readInto()

	owners := controller.Schema.Settings.GetOwners()

	delete(owners, name)

	controller.Schema.Settings.SetOwners(owners)
	controller.View.SetSchema(controller.Schema)

}

func (controller *Settings) addIdentity(target *dom.Element) {

	owner_name := controller.ownerNameFromTarget(target)

	if owner_name == "" || controller.Schema == nil {
		return
	}

	controller.readInto()

	owner := controller.Schema.Settings.GetOwner(owner_name)

	if owner == nil {
		return
	}

	name := "identity"

	for i := 1; ; i++ {

		candidate := "identity-" + strconv.Itoa(i)

		if owner.GetIdentity(candidate) == nil {
			name = candidate
			break
		}

	}

	owner.SetIdentity(types.NewIdentity(name))
	controller.View.SetSchema(controller.Schema)

}

func (controller *Settings) addRemote(target *dom.Element) {

	owner_name := controller.ownerNameFromTarget(target)

	if owner_name == "" || controller.Schema == nil {
		return
	}

	controller.readInto()

	owner := controller.Schema.Settings.GetOwner(owner_name)

	if owner == nil {
		return
	}

	name := "remote"

	for i := 1; ; i++ {

		candidate := "remote-" + strconv.Itoa(i)

		if owner.GetRemote(candidate) == nil {
			name = candidate
			break
		}

	}

	owner.SetRemote(types.NewRemote(name, ""))

	controller.View.SetSchema(controller.Schema)

}

func (controller *Settings) addService(target *dom.Element) {

	owner_name := controller.ownerNameFromTarget(target)

	if owner_name == "" || controller.Schema == nil {
		return
	}

	controller.readInto()

	owner := controller.Schema.Settings.GetOwner(owner_name)

	if owner == nil {
		return
	}

	name := "service"

	for i := 1; ; i++ {

		candidate := "service-" + strconv.Itoa(i)

		if owner.GetService(candidate) == nil {
			name = candidate
			break
		}

	}

	owner.SetService(types.NewService(name))

	controller.View.SetSchema(controller.Schema)

}

func (controller *Settings) removeEntry(target *dom.Element, kind string) {

	owner_name := controller.ownerNameFromTarget(target)

	if owner_name == "" || controller.Schema == nil {
		return
	}

	fieldset := target.QueryParent("fieldset")

	if fieldset == nil {
		return
	}

	name := fieldset.GetAttribute("data-name")

	if name == "" {
		return
	}

	controller.readInto()

	owner := controller.Schema.Settings.GetOwner(owner_name)

	if owner == nil {
		return
	}

	if kind == "identities" {
		owner.RemoveIdentity(name)
	} else if kind == "remotes" {
		owner.RemoveRemote(name)
	} else if kind == "services" {
		owner.RemoveService(name)
	}

	controller.View.SetSchema(controller.Schema)

}

func (controller *Settings) ownerNameFromTarget(target *dom.Element) string {

	article := target.QueryParent("article")

	if article == nil {
		return ""
	}

	return article.GetAttribute("data-name")

}
