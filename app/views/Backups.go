package views

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/app"
import "git-evac-app/actions"
// import app_schemas "git-evac-app/schemas"
// import "git-evac/server/schemas"
// import "git-evac/structs"
// import "slices"
// import "sort"
// import "strconv"
// import "strings"

type Backups struct {
	Main *app.Main `json:"main"`
	app.BaseView
}

func NewBackups(main *app.Main) Backups {

	var view Backups

	view.Main     = main
	view.Elements = make(map[string]*dom.Element)

	view.SetElement("table", dom.Document.QuerySelector("main table"))
	view.SetElement("dialog", dom.Document.QuerySelector("dialog"))
	view.SetElement("footer", dom.Document.QuerySelector("footer"))

	view.Init()

	return view

}

func (view Backups) Init() {

	dialog := view.GetElement("dialog")
	footer := view.GetElement("footer")
	table := view.GetElement("table")

	if table != nil {

		// TODO

	}

	if dialog != nil {

		// TODO

	}

	if footer != nil {

		// TODO

	}

}

func (view Backups) Enter() bool {

	schema, err := actions.Index()

	if err == nil {
		view.Main.Storage.Write("repositories", schema)
	}

	view.Render()

	return true

}

func (view Backups) Leave() bool {
	return true
}

func (view Backups) Render() {

	// TODO

}

