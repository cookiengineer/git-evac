package controllers

// import "github.com/cookiengineer/gooey/bindings/dom"
// import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
// import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/interfaces"
// import "git-evac-app/actions"

type Settings struct {
	Main   *app.Main `json:"main"`
	Schema any       `json:"schema"`
	View   *app.View `json:"view"`
}

func NewSettings(main *app.Main, view interfaces.View) *Settings {

	var controller Settings

	controller.Main = main
	controller.View = view.(*app.View)

	return &controller

}

func (controller *Settings) Enter() bool {

	// TODO: Add Event Listeners

	go controller.Update()

	return true

}

func (controller *Settings) Leave() bool {

	// TODO: Remove Event Listeners

	return true

}

func (controller *Settings) Name() string {
	return "settings"
}

func (controller *Settings) Update() {

	if controller.Main != nil {

		// TODO: ReadSettings()

		controller.Render()

	}

}

func (controller *Settings) Render() {
	controller.View.Render()
}
