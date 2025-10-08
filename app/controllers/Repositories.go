package controllers

// import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
import "git-evac/server/schemas"
import "git-evac-app/actions"
import "fmt"

type Repositories struct {
	Main   *app.Main             `json:"main"`
	Schema *schemas.Repositories `json:"schema"`
	View   *app.View             `json:"view"`
}

func NewRepositories(main *app.Main, view *app.View) *Repositories {

	var controller Repositories

	controller.Main = main
	controller.View = view

	table, ok1 := components.Unwrap[*content.Table](controller.View.Query(""))

	if table != nil && ok1 == true {

		fmt.Println("table", table)

	}

	controller.Update()

	return &controller

}

func (controller *Repositories) Name() string {
	return "repositories"
}

func (controller *Repositories) Update() {

	if controller.Main != nil {

		schema, err := actions.Index()

		if err == nil {
			controller.Schema = schema
			controller.Main.Storage.Write("repositories", schema)
		}

		controller.Render()

	}

}

func (controller *Repositories) Render() {
	controller.View.Render()
}
