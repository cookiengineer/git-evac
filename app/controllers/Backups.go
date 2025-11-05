package controllers

// import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac-app/actions"
import app_components "git-evac-app/components"
import app_views "git-evac-app/views"

type Backups struct {
	Main   *app.Main          `json:"main"`
	Schema any                `json:"schema"`
	View   *app_views.Backups `json:"view"`
}

func NewBackups(main *app.Main, view interfaces.View) *Backups {

	var controller Backups

	controller.Main = main
	controller.View = view.(*app_views.Backups)

	dialog := controller.Main.Dialog
	footer := controller.Main.Footer

	table, ok1 := components.UnwrapComponent[*app_components.BackupsTable](controller.View.Query("section > table[data-name=\"backups\"]"))

	if dialog != nil && footer != nil && table != nil && ok1 == true {

		dialog.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

			if event == "action" {

				action, ok := attributes["action"].(string)

				if ok == true {

					if action == "confirm" {

						scheduler_table, ok2 := components.UnwrapComponent[*app_components.SchedulerTable](dialog.Query("dialog > table[data-name=\"scheduler\"]"))

						if ok2 == true {

							go func() {
								scheduler_table.Start()
							}()

						}

					} else if action == "cancel" {

						scheduler_table, ok2 := components.UnwrapComponent[*app_components.SchedulerTable](dialog.Query("dialog > table[data-name=\"scheduler\"]"))

						if ok2 == true {

							go func() {
								scheduler_table.Stop()
								scheduler_table.Reset()
							}()

						}

						dialog.Hide()

					}

				}

			}

		}, false))

		footer.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

			if event == "action" {

				action, ok := attributes["action"].(string)

				if ok == true {

					selected := table.Selected()

					if len(selected) > 0 {

						filtered := make(map[string]string)

						for repository, available_action := range selected {

							if action == "backup" && available_action == "backup" {
								filtered[repository] = "backup"
							} else if action == "backup" && available_action == "backup-or-restore" {
								filtered[repository] = "backup"
							} else if action == "restore" && available_action == "restore" {
								filtered[repository] = "restore"
							} else if action == "restore" && available_action == "backup-or-restore" {
								filtered[repository] = "restore"
							}

						}

						controller.showDialog(filtered)

					}

				}

			}

		}, false))

		table.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

			if event == "action" {

				filtered := make(map[string]string)

				for repository, raw_action := range attributes {

					action, ok := raw_action.(string)

					if ok == true {
						filtered[repository] = action
					}

				}

				if len(filtered) == 1 {
					controller.showDialog(filtered)
				}

			}

		}, false))

		table.Component.AddEventListener("select", components.ToEventListener(func(event string, attributes map[string]any) {

			if event == "select" {

				actions_backup := make([]string, 0)
				actions_restore := make([]string, 0)

				for repository, raw_action := range attributes {

					action, ok := raw_action.(string)

					if ok == true {

						if action == "backup" {
							actions_backup = append(actions_backup, repository)
						} else if action == "restore" {
							actions_restore = append(actions_restore, repository)
						} else if action == "backup-or-restore" {
							actions_backup = append(actions_backup, repository)
							actions_restore = append(actions_restore, repository)
						}

					}

				}

				length_all := 0

				for _, owner := range controller.Schema.Owners {
					length_all += len(owner.Repositories)
				}

				label, ok0 := components.UnwrapComponent[*ui_components.Label](footer.Query("footer > label"))

				if ok0 == true {
					label.SetLabel("Selected " + strconv.Itoa(len(attributes)) + " of " + strconv.Itoa(length_all) + " Repositories")
				}

				buttons_backup, ok1 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"backup\"]"))
				buttons_restore, ok2 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"restore\"]"))

				if ok1 == true {

					buttons_backup.SetLabel("Backup " + strconv.Itoa(len(actions_backup)))

					if len(actions_backup) > 0 {
						buttons_backup.Enable()
					} else {
						buttons_backup.Disable()
					}

				}

				if ok2 == true {

					buttons_restore.SetLabel("Restore " + strconv.Itoa(len(actions_restore)))

					if len(actions_restore) > 0 {
						buttons_restore.Enable()
					} else {
						buttons_restore.Disable()
					}

				}

			}

		}, false))

	}

	controller.Update()

	return &controller

}

func (controller *Backups) Name() string {
	return "backups"
}

func (controller *Backups) Update() {

	if controller.Main != nil {

		schema, err := actions.Index()

		if err == nil {

			controller.Schema = schema
			controller.Main.Storage.Write("repositories", schema)

			table, ok1 := components.UnwrapComponent[*app_components.BackupsTable](controller.View.Query("section > table[data-name=\"backups\"]"))

			if len(controller.Schema.Owners) > 0 && ok1 == true {

				table.Reset()
				table.SetSchema(controller.Schema)

			}

			footer := controller.Main.Footer

			if footer != nil {

				length := 0

				for _, owner := range controller.Schema.Owners {
					length += len(owner.Repositories)
				}

				label, ok0 := components.UnwrapComponent[*ui_components.Label](footer.Query("footer > label"))

				if ok0 == true {
					label.SetLabel("Selected 0 of " + strconv.Itoa(length) + " Repositories")
				}

				buttons_backup, ok1 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"backup\"]"))
				buttons_restore, ok2 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"restore\"]"))

				if ok1 == true {
					buttons_backup.SetLabel("Backup 0")
					buttons_backup.Disable()
				}

				if ok2 == true {
					buttons_restore.SetLabel("Restore 0")
					buttons_restore.Disable()
				}

			}

		}

		controller.Render()

	}

}

func (controller *Backups) Render() {
	controller.View.Render()
}

func (controller *Backups) showDialog(selected map[string]string) {

	dialog := controller.Main.Dialog

	if dialog != nil {

		actions_backup := make(map[string]string)
		actions_restore := make(map[string]string)

		for repository, action := range selected {

			if action == "backup" {
				actions_backup[repository] = "backup"
			} else if action == "restore" {
				actions_restore[repository] = "restore"
			}

		}

		if len(actions_backup) > 0 {

			table := app_components.NewSchedulerTable("scheduler", actions_backup)

			dialog.SetTitle("Backup " + strconv.Itoa(len(actions_backup)) + " Repositories")
			dialog.SetContent(interfaces.Component(&table))
			dialog.Show()

		} else if len(actions_restore) > 0 {

			table := app_components.NewSchedulerTable("scheduler", actions_restore)

			dialog.SetTitle("Restore " + strconv.Itoa(len(actions_restore)) + " Repositories")
			dialog.SetContent(interfaces.Component(&table))
			dialog.Show()

		}

	}

}
