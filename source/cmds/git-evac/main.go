package main

import "git-evac/public"
import "git-evac/server"
import "git-evac/structs"
import "git-evac/webview"
import "encoding/json"
import "io/fs"
import "os"
import "os/signal"
import os_user "os/user"
import "strconv"
import "strings"
import "syscall"
import "time"

func main() {

	var config string = ""

	user, _ := os_user.Current()
	console := structs.NewConsole(os.Stdout, os.Stderr, 0)

	if len(os.Args) >= 2 {

		parameters := os.Args[1:]

		for _, parameter := range parameters {

			if strings.HasPrefix(parameter, "--config=") {

				tmp := strings.TrimSpace(parameter[9:])

				if strings.HasPrefix(tmp, "\"") && strings.HasSuffix(tmp, "\"") {
					tmp = strings.TrimSpace(tmp[1 : len(tmp)-1])
				} else {
					tmp = strings.TrimSpace(tmp)
				}

				if strings.HasPrefix(tmp, "~/") {
					tmp = user.HomeDir + "/" + tmp[2:]
				} else if strings.Contains(tmp, "~") {
					console.Error("Malformed Config Parameter: \"" + tmp + "\"")
					tmp = ""
				}

				if tmp != "" && strings.HasSuffix(tmp, ".json") {
					config = tmp
				}

			}

		}

	}

	if config == "" {
		config = user.HomeDir + "/.config/git-evac/git-evac.json"
	}

	if config != "" {

		settings := structs.NewSettings(user.HomeDir + "/Backup", user.HomeDir + "/Software", 3000)
		buffer1, err1 := os.ReadFile(config)

		valid_config := false

		if err1 == nil {

			err2 := json.Unmarshal(buffer1, settings)

			if err2 == nil {
				valid_config = true
			} else {
				console.Error(err2.Error())
			}

		}


		fsys, _ := fs.Sub(public.FS, ".")
		profile := structs.NewProfile(console, settings)
		profile.Filesystem = &fsys

		console.Clear("")
		console.Group("git-evac")

		if valid_config == true {
			console.Log("> Config: " + config)
		} else {
			console.Warn("> Invalid Config: " + config)
		}

		console.GroupEnd("git-evac")

		signal_channel := make(chan os.Signal, 1)
		signal.Notify(
			signal_channel,
			syscall.SIGINT,
			syscall.SIGKILL,
			syscall.SIGTERM,
		)

		done := make(chan bool, 2)

		go func() {

			profile.Refresh()

			server.Dispatch(profile)
			server.DispatchRoutes(profile)

			result := server.Serve(profile)

			if result == false {
				console.Error("Port " + strconv.FormatUint(uint64(profile.Settings.Port), 10) + " is probably already in use?")
			}

			done <- result

		}()

		go func() {

			time.Sleep(3 * time.Second)

			console.Log("Opening WebView...")

			view := webview.New(true)
			view.SetTitle("Git Evac")
			view.SetSize(800, 600, webview.HintNone)
			view.Navigate("http://localhost:" + strconv.FormatUint(uint64(profile.Settings.Port), 10) + "/index.html")

			view.Run()

			done <- true

		}()

		select {
		case <-done:
			console.Log("The WebView or Server has been closed, exiting...")
		case <-signal_channel:
			console.Log("Received OS signal, exiting...")
		}

		// Give WebView time to cleanup
		time.Sleep(250 * time.Millisecond)

	}

}
