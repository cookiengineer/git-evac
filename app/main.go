package main

import "gooey"
import "gooey/console"
import "gooey/timers"
import "time"

func main() {

	main := gooey.Document.QuerySelector("main")

	timers.SetTimeout(func() {
		console.Inspect(main)
	}, 1000)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
