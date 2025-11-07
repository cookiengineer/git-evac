package controllers

import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/interfaces"
import ui_components "github.com/cookiengineer/gooey/components/ui"
import "git-evac/schemas"
import app_actions "git-evac-app/actions"
import app_components "git-evac-app/components"
import app_views "git-evac-app/views"
import "strconv"

type Backups struct {
	Main    *app.Main          `json:"main"`
	View    *app_views.Backups `json:"view"`
	Schemas struct {
		Backups      *schemas.Backups      `json:"backups"`
		Repositories *schemas.Repositories `json:"repositories"`
	} `json:"schemas"`
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

				mapped := make(map[string]bool)

				for _, owner := range controller.Schemas.Backups.Owners {

					for _, backup := range owner.Backups {
						mapped[owner.Name + "/" + backup.Name] = true
					}

				}

				for _, owner := range controller.Schemas.Repositories.Owners {

					for _, repository := range owner.Repositories {
						mapped[owner.Name + "/" + repository.Name] = true
					}

				}

				label, ok0 := components.UnwrapComponent[*ui_components.Label](footer.Query("footer > label"))

				if ok0 == true {
					label.SetLabel("Selected " + strconv.Itoa(len(attributes)) + " of " + strconv.Itoa(len(mapped)) + " Items")
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

		schema_backups, err_backups := app_actions.Backups()
		schema_repositories, err_repositories := app_actions.Repositories()

		if err_backups == nil && err_repositories == nil {

			controller.Schemas.Backups = schema_backups
			controller.Schemas.Repositories = schema_repositories

			controller.Main.Storage.Write("backups", schema_backups)
			controller.Main.Storage.Write("repositories", schema_repositories)

			table, ok1 := components.UnwrapComponent[*app_components.BackupsTable](controller.View.Query("section > table[data-name=\"backups\"]"))

			if ok1 == true && (len(controller.Schemas.Backups.Owners) > 0 || len(controller.Schemas.Repositories.Owners) > 0) {

				table.Reset()
				table.SetSchema(controller.Schemas.Backups, controller.Schemas.Repositories)

			}

			footer := controller.Main.Footer

			if footer != nil {

				mapped := make(map[string]bool)

				for _, backup_owner := range controller.Schemas.Backups.Owners {

					for _, backup := range backup_owner.Backups {
						mapped[backup_owner.Name + "/" + backup.Name] = true
					}

				}

				for _, repository_owner := range controller.Schemas.Repositories.Owners {

					for _, repository := range repository_owner.Repositories {
						mapped[repository_owner.Name + "/" + repository.Name] = false
					}

				}

				label, ok0 := components.UnwrapComponent[*ui_components.Label](footer.Query("footer > label"))

				if ok0 == true {
					label.SetLabel("Selected 0 of " + strconv.Itoa(len(mapped)) + " Items")
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

			dialog.SetTitle("Restore " + strconv.Itoa(len(actions_restore)) + " Backups")
			dialog.SetContent(interfaces.Component(&table))
			dialog.Show()

		}

	}

}
