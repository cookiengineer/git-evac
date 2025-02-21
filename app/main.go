package main

import "fmt"

import "app/client/api"
import "gooey"
import "gooey/timers"
import "app/components"
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
			components.RenderTable(index)
		}

	}, 500)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
