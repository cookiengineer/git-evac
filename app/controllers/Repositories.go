//go:build wasm

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

type Repositories struct {
	Main   *app.Main               `json:"main"`
	Schema *schemas.Repositories   `json:"schema"`
	View   *app_views.Repositories `json:"view"`
}

func NewRepositories(main *app.Main, view interfaces.View) *Repositories {

	var controller Repositories

	controller.Main = main
	controller.View = view.(*app_views.Repositories)

	dialog := controller.Main.Dialog
	footer := controller.Main.Footer
	table, ok1 := components.UnwrapComponent[*app_components.RepositoriesTable](controller.View.Query("section > table[data-name=\"repositories\"]"))

	if dialog != nil && footer != nil && table != nil && ok1 == true {

		dialog.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

			if event == "action" {

				action, ok := attributes["action"].(string)

				if ok == true {

					if action == "confirm" {

						schedule_table, ok2 := components.UnwrapComponent[*app_components.SchedulerTable](dialog.Query("dialog > table[data-name=\"schedule\"]"))

						if ok2 == true {

							go func() {
								schedule_table.Start()
							}()

						}

					} else if action == "cancel" {

						schedule_table, ok2 := components.UnwrapComponent[*app_components.SchedulerTable](dialog.Query("dialog > table[data-name=\"schedule\"]"))

						if ok2 == true {

							go func() {
								schedule_table.Stop()
								schedule_table.Reset()
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

							if action == "clone" && available_action == "clone" {
								filtered[repository] = "clone"
							} else if action == "fix" && available_action == "fix" {
								filtered[repository] = "fix"
							} else if action == "commit" && available_action == "commit" {
								filtered[repository] = "commit"
							} else if action == "pull" && available_action == "pull-or-push" {
								filtered[repository] = "pull"
							} else if action == "push" && available_action == "pull-or-push" {
								filtered[repository] = "push"
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

				actions_clone  := make([]string, 0)
				actions_fix    := make([]string, 0)
				actions_commit := make([]string, 0)
				actions_pull   := make([]string, 0)
				actions_push   := make([]string, 0)

				for repository, raw_action := range attributes {

					action, ok := raw_action.(string)

					if ok == true {

						if action == "clone" {
							actions_clone = append(actions_clone, repository)
						} else if action == "fix" {
							actions_fix = append(actions_fix, repository)
						} else if action == "commit" {
							actions_commit = append(actions_commit, repository)
						} else if action == "pull-or-push" {
							actions_pull = append(actions_pull, repository)
							actions_push = append(actions_push, repository)
						} else if action == "pull" {
							actions_pull = append(actions_pull, repository)
						} else if action == "push" {
							actions_push = append(actions_push, repository)
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

				buttons_clone, ok1 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"clone\"]"))
				buttons_fix, ok2 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"fix\"]"))
				buttons_commit, ok3 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"commit\"]"))
				buttons_pull, ok4 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"pull\"]"))
				buttons_push, ok5 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"push\"]"))

				if ok1 == true {

					buttons_clone.SetLabel("Clone " + strconv.Itoa(len(actions_clone)))

					if len(actions_clone) > 0 {
						buttons_clone.Enable()
					} else {
						buttons_clone.Disable()
					}

				}

				if ok2 == true {

					buttons_fix.SetLabel("Fix " + strconv.Itoa(len(actions_fix)))

					if len(actions_fix) > 0 {
						buttons_fix.Enable()
					} else {
						buttons_fix.Disable()
					}

				}

				if ok3 == true {

					buttons_commit.SetLabel("Commit " + strconv.Itoa(len(actions_commit)))

					if len(actions_commit) > 0 {
						buttons_commit.Enable()
					} else {
						buttons_commit.Disable()
					}

				}

				if ok4 == true {

					buttons_pull.SetLabel("Pull " + strconv.Itoa(len(actions_pull)))

					if len(actions_pull) > 0 {
						buttons_pull.Enable()
					} else {
						buttons_pull.Disable()
					}

				}

				if ok5 == true {

					buttons_push.SetLabel("Push " + strconv.Itoa(len(actions_push)))

					if len(actions_push) > 0 {
						buttons_push.Enable()
					} else {
						buttons_push.Disable()
					}

				}

			}

		}, false))

	}

	controller.Update()

	return &controller

}

func (controller *Repositories) Name() string {
	return "repositories"
}

func (controller *Repositories) Update() {

	if controller.Main != nil {

		schema, err := app_actions.Index()

		if err == nil {

			controller.Schema = schema
			controller.Main.Storage.Write("repositories", schema)

			table, ok1 := components.UnwrapComponent[*app_components.RepositoriesTable](controller.View.Query("section > table"))

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

				buttons_clone, ok1 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"clone\"]"))
				buttons_fix, ok2 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"fix\"]"))
				buttons_commit, ok3 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"commit\"]"))
				buttons_pull, ok4 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"pull\"]"))
				buttons_push, ok5 := components.UnwrapComponent[*ui_components.Button](footer.Query("footer > button[data-action=\"push\"]"))

				if ok1 == true {
					buttons_clone.SetLabel("Clone 0")
					buttons_clone.Disable()
				}

				if ok2 == true {
					buttons_fix.SetLabel("Fix 0")
					buttons_fix.Disable()
				}

				if ok3 == true {
					buttons_commit.SetLabel("Commit 0")
					buttons_commit.Disable()
				}

				if ok4 == true {
					buttons_pull.SetLabel("Pull 0")
					buttons_pull.Disable()
				}

				if ok5 == true {
					buttons_push.SetLabel("Push 0")
					buttons_push.Disable()
				}

			}

		}

		controller.Render()

	}

}

func (controller *Repositories) Render() {
	controller.View.Render()
}

func (controller *Repositories) showDialog(selected map[string]string) {

	dialog := controller.Main.Dialog

	if dialog != nil {

		actions_clone  := make(map[string]string)
		actions_fix    := make(map[string]string)
		actions_commit := make(map[string]string)
		actions_pull   := make(map[string]string)
		actions_push   := make(map[string]string)

		for repository, action := range selected {

			if action == "clone" {
				actions_clone[repository] = "clone"
			} else if action == "fix" {
				actions_fix[repository] = "fix"
			} else if action == "commit" {
				actions_commit[repository] = "commit"
			} else if action == "pull" {
				actions_pull[repository] = "pull"
			} else if action == "push" {
				actions_push[repository] = "push"
			}

		}

		if len(actions_clone) > 0 {

			table := app_components.NewSchedulerTable("scheduler", actions_clone)

			dialog.SetTitle("Clone " + strconv.Itoa(len(actions_clone)) + " Repositories")
			dialog.SetContent(interfaces.Component(&table))
			dialog.Show()

		} else if len(actions_fix) > 0 {

			table := app_components.NewSchedulerTable("scheduler", actions_fix)

			dialog.SetTitle("Fix " + strconv.Itoa(len(actions_fix)) + " Repositories")
			dialog.SetContent(interfaces.Component(&table))
			dialog.Show()

		} else if len(actions_commit) > 0 {

			table := app_components.NewSchedulerTable("scheduler", actions_commit)

			dialog.SetTitle("Commit " + strconv.Itoa(len(actions_commit)) + " Repositories")
			dialog.SetContent(interfaces.Component(&table))
			dialog.Show()

		} else if len(actions_pull) > 0 {

			table := app_components.NewSchedulerTable("scheduler", actions_pull)

			dialog.SetTitle("Pull " + strconv.Itoa(len(actions_pull)) + " Repositories")
			dialog.SetContent(interfaces.Component(&table))
			dialog.Show()

		} else if len(actions_push) > 0 {

			table := app_components.NewSchedulerTable("scheduler", actions_push)

			dialog.SetTitle("Push " + strconv.Itoa(len(actions_push)) + " Repositories")
			dialog.SetContent(interfaces.Component(&table))
			dialog.Show()

		}

	}

}
