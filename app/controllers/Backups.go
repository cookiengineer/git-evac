package controllers

// import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/interfaces"
import "git-evac-app/actions"
import "fmt"

type Backups struct {
	Main   *app.Main `json:"main"`
	Schema any       `json:"schema"`
	View   *app.View `json:"view"`
}

func NewBackups(main *app.Main, view interfaces.View) *Backups {

	var controller Backups

	controller.Main = main
	controller.View = view.(*app.View)
	// controller.View = app.NewView("backups", "Backups", "/backups.html")

	// controller.View.SetElement("table", dom.Document.QuerySelector("main table"))
	// controller.View.SetElement("dialog", dom.Document.QuerySelector("dialog"))
	// controller.View.SetElement("footer", dom.Document.QuerySelector("footer"))

	table, ok1 := components.Unwrap[*content.Table](controller.View.Query(""))

	if table != nil && ok1 == true {

		fmt.Println("table", table)

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
		}

		controller.Render()

	}

}

func (controller *Backups) Render() {
	controller.View.Render()
}
