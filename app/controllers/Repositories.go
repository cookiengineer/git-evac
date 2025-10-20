//go:build wasm

package controllers

// import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/interfaces"
import ui_components "github.com/cookiengineer/gooey/components/ui"
import "git-evac/schemas"
import app_actions "git-evac-app/actions"
import app_components "git-evac-app/components"
import app_views "git-evac-app/views"
import "strconv"
import "fmt"

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

						schedule_table, ok2 := components.UnwrapComponent[*app_components.ScheduleTable](dialog.Query("dialog > table[data-name=\"schedule\"]"))

						fmt.Println(schedule_table, ok2)

						if ok2 == true {

							schedule_table.Start("any")

							fmt.Println("Start all goroutines now")
							// TODO: Start all tasks asynchronously and call schedule_table.Finish(repository, &response)

						}

					} else if action == "cancel" {

						schedule_table, ok2 := components.UnwrapComponent[*app_components.ScheduleTable](dialog.Query("dialog > table[data-name=\"schedule\"]"))

						fmt.Println(schedule_table, ok2)

						if ok2 == true {
							schedule_table.Reset()
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
				table.SetSchema(controller.Schema)
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

			// TODO: Other actions

			dialog.SetTitle("Clone " + strconv.Itoa(len(actions_clone)) + " Repositories")

		} else if len(actions_fix) > 0 {

			// TODO: Other actions

			dialog.SetTitle("Fix " + strconv.Itoa(len(actions_fix)) + " Repositories")

		} else if len(actions_commit) > 0 {

			// TODO: Other actions

			dialog.SetTitle("Commit " + strconv.Itoa(len(actions_commit)) + " Repositories")

		} else if len(actions_pull) > 0 {

			table := app_components.NewScheduleTable("schedule", actions_pull)

			dialog.SetTitle("Pull " + strconv.Itoa(len(actions_pull)) + " Repositories")
			dialog.SetContent(interfaces.Component(&table))
			dialog.Show()

		} else if len(actions_push) > 0 {

			// TODO: Other actions

			dialog.SetTitle("Push " + strconv.Itoa(len(actions_push)) + " Repositories")

		}

	}

}
