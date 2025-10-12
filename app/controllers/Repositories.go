package controllers

// import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/data"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/server/schemas"
import app_actions "git-evac-app/actions"
import app_components "git-evac-app/components"
import app_views "git-evac-app/views"
import "sort"
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

	table, ok1 := components.UnwrapComponent[*app_components.RepositoriesTable](controller.View.Query("section > table[data-name=\"repositories\"]"))

	if table != nil && ok1 == true {

		table.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

			if event == "action" {

				action, ok := attributes["action"].(string)

				if ok == true {

					fmt.Println("Table action event:", action)

					if action == "clone" {

						// TODO

					} else if action == "fix" {

						// TODO

					} else if action == "commit" {

						// TODO

					} else if action == "pull" {

						// TODO

					} else if action == "push" {

						// TODO

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

				dataset := data.NewDataset(0)

				for _, owner := range controller.Schema.Owners {

					for _, repository := range owner.Repositories {

						actions := make([]string, 0)
						branches := make([]string, 0)
						remotes := make([]string, 0)

						for _, branch_name := range repository.Branches {
							branches = append(branches, branch_name)
						}

						for remote_name, _ := range repository.Remotes {
							remotes = append(remotes, remote_name)
						}

						if repository.HasRemoteChanges == true {
							actions = append(actions, "fix")
						} else if repository.HasLocalChanges == true {
							actions = append(actions, "commit")
						} else {
							actions = append(actions, "pull")
							actions = append(actions, "push")
						}

						sort.Strings(actions)
						sort.Strings(branches)
						sort.Strings(remotes)

						dataset.Add(data.Data(map[string]any{
							"repository": owner.Name + "/" + repository.Name,
							"remotes":    remotes,
							"branches":   branches,
							"actions":    actions,
						}))

						table.SetDataset(dataset)
						table.SortBy("repository")

					}

				}

			}

		}

		controller.Render()

	}

}

func (controller *Repositories) Render() {
	controller.View.Render()
}
