package controllers

// import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/interfaces"
import "git-evac/schemas"
import app_actions "git-evac-app/actions"
import app_components "git-evac-app/components"
import app_views "git-evac-app/views"
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
				table.SetSchema(controller.Schema)
			}

		}

		controller.Render()

	}

}

func (controller *Repositories) Render() {
	controller.View.Render()
}
