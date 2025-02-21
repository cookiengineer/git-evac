package main

import "app/client/api"
import "gooey"
import "gooey/timers"
import "app/components"
import "app/storage"
import "time"

func main() {

	main := gooey.Document.QuerySelector("main")
	dialog := gooey.Document.QuerySelector("dialog")

	timers.SetTimeout(func() {

		components.InitDialog()
		// TODO: components.InitHeader()
		components.InitTable()
		components.InitFooter()

	}, 0)

	timers.SetTimeout(func() {

		index, err := api.Index()

		if err == nil {
			storage.Index = index
		}

		components.RenderTable(storage.Index)

	}, 500)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
