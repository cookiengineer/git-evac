package main

import webview "github.com/webview/webview_go"
import "git-evac/actions"
import "git-evac/console"
import "embed"
import "io/fs"
import "os"

//go:embed public/*
var embedded_filesystem embed.FS

func main() {

	mode := ""

	if len(os.Args) == 2 {

		if os.Args[1] == "dev" || os.Args[1] == "development" {
			mode = "development"
		} else {
			mode = "production"
		}

	} else {
		mode = "production"
	}

	if mode == "development" {

		fsys := os.DirFS("public")

		console.Clear()
		console.Group("git-evac: Command-Line Arguments")
		console.Inspect(struct {
			Mode string
		}{
			Mode: mode,
		})
		console.GroupEnd("")

		console.Log("Listening on http://localhost:13337")
		actions.Serve(fsys, 13337)

	} else if mode == "production" {

		fsys, _ := fs.Sub(embedded_filesystem, "public")

		console.Clear()
		console.Group("git-evac: Command-Line Arguments")
		console.Inspect(struct {
			Mode string
		}{
			Mode: mode,
		})
		console.GroupEnd("")

		console.Log("Listening on http://localhost:13337")

		go func() {

			console.Log("Opening WebView...")

			view := webview.New(true)
			view.SetTitle("Git Evac")
			view.SetSize(800, 600, webview.HintNone)
			view.Navigate("http://localhost:13337/index.html")
			view.Run()
			// defer view.Destroy()

		}()

		actions.Serve(fsys, 13337)

	}

}
