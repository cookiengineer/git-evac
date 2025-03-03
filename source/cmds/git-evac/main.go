package main

import "git-evac/console"
import "git-evac/public"
import "git-evac/server"
import "git-evac/structs"
import "git-evac/webview"
import "io/fs"
import "os"
import os_user "os/user"
import "strconv"
import "strings"
import "time"

func main() {

	var folder string = ""
	var port uint16 = 1234

	if len(os.Args) >= 2 {

		parameters := os.Args[1:]

		for p := 0; p < len(parameters); p++ {

			parameter := parameters[p]

			if strings.HasPrefix(parameter, "--folder=") {

				tmp := strings.TrimSpace(parameter[9:])

				if strings.HasPrefix(tmp, "\"") && strings.HasSuffix(tmp, "\"") {
					tmp = strings.TrimSpace(tmp[1:len(tmp)-1])
				} else {
					tmp = strings.TrimSpace(tmp)
				}

				if strings.HasPrefix(tmp, "~/") {

					user, err := os_user.Current()

					if err == nil {
						tmp = user.HomeDir + "/" + tmp[2:]
					}

				} else if strings.Contains(tmp, "~") {
					console.Error("Malformed Folder Parameter: " + tmp)
					tmp = ""
				}

				if tmp != "" {

					stat, err := os.Stat(tmp)

					if err == nil && stat.IsDir() {
						folder = tmp
					}

				}

			} else if strings.HasPrefix(parameter, "--port=") {

				tmp := strings.TrimSpace(parameter[7:])

				num, err := strconv.ParseUint(tmp, 10, 16)

				if err == nil && num > 0 && num < 65535 {
					port = uint16(num)
				}

			}

		}

	}

	if folder == "" {

		user, err := os_user.Current()

		if err == nil {
			folder = user.HomeDir + "/Software"
		}

	}

	if folder != "" {

		filesystem, _ := fs.Sub(public.FS, ".")
		profile := structs.NewProfile(folder, port)
		profile.Filesystem = &filesystem

		console.Clear()
		console.Group("git-evac: Command-Line Arguments")
		console.Inspect(struct {
			Folder string
			Port   uint16
		}{
			Folder: folder,
			Port:   port,
		})
		console.GroupEnd("")

		go func() {

			time.Sleep(1 * time.Second)

			console.Log("Opening WebView...")

			view := webview.New(true)
			view.SetTitle("Git Evac")
			view.SetSize(800, 600, webview.HintNone)
			view.Navigate("http://localhost:" + strconv.FormatUint(uint64(port), 10) + "/index.html")
			view.Run()
			// defer view.Destroy()

		}()

		profile.Init()
		server.Dispatch(profile)

		if server.Serve(profile) == false {
			console.Error("Port " + strconv.FormatUint(uint64(port), 10) + " is already in use.")
			os.Exit(1)
		} else {
			os.Exit(0)
		}

	}

}
