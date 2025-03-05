package views

import "gooey"
import "gooey/app"
import "gooey/dom"
import "git-evac-app/actions"
// import app_schemas "git-evac-app/schemas"
// import "git-evac/server/schemas"
// import "git-evac/structs"
// import "slices"
// import "sort"
// import "strconv"
// import "strings"

type Settings struct {
	Main *app.Main `json:"main"`
	app.BaseView
}

func NewSettings(main *app.Main) Settings {

	var view Settings

	view.Main     = main
	view.Elements = make(map[string]*dom.Element)

	view.SetElement("table", gooey.Document.QuerySelector("main table"))
	view.SetElement("footer", gooey.Document.QuerySelector("footer"))

	view.Init()

	return view

}

func (view Settings) Init() {

	footer := view.GetElement("footer")
	table := view.GetElement("table")

	if table != nil {

		// TODO

	}

	if footer != nil {

		// TODO

	}

}

func (view Settings) Enter() bool {

	schema, err := actions.ReadSettings()

	if err == nil {
		view.Main.Storage.Write("settings", schema)
	}

	view.Render()

	return true

}

func (view Settings) Leave() bool {
	return true
}

func (view Settings) Render() {

	// TODO

}

