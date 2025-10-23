package main

import "git-evac/public"
import "git-evac/server"
import "git-evac/structs"
import "git-evac/webview"
import "io/fs"
import "os"
import "os/signal"
import os_user "os/user"
import "strconv"
import "strings"
import "syscall"
import "time"

func main() {

	var backup string = ""
	var folder string = ""
	var port uint16 = 3000

	console := structs.NewConsole(os.Stdout, os.Stderr, 0)

	if len(os.Args) >= 2 {

		parameters := os.Args[1:]

		for p := 0; p < len(parameters); p++ {

			parameter := parameters[p]

			if strings.HasPrefix(parameter, "--backup=") {

				tmp := strings.TrimSpace(parameter[9:])

				if strings.HasPrefix(tmp, "\"") && strings.HasSuffix(tmp, "\"") {
					tmp = strings.TrimSpace(tmp[1 : len(tmp)-1])
				} else {
					tmp = strings.TrimSpace(tmp)
				}

				if strings.HasPrefix(tmp, "~/") {

					user, err := os_user.Current()

					if err == nil {
						tmp = user.HomeDir + "/" + tmp[2:]
					}

				} else if strings.Contains(tmp, "~") {
					console.Error("Malformed Backup Parameter: " + tmp)
					tmp = ""
				}

				if tmp != "" {

					stat, err := os.Stat(tmp)

					if err == nil && stat.IsDir() {
						backup = tmp
					}

				}

			} else if strings.HasPrefix(parameter, "--folder=") {

				tmp := strings.TrimSpace(parameter[9:])

				if strings.HasPrefix(tmp, "\"") && strings.HasSuffix(tmp, "\"") {
					tmp = strings.TrimSpace(tmp[1 : len(tmp)-1])
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

	user, err := os_user.Current()

	if err == nil {

		if backup == "" {
			backup = user.HomeDir + "/Backup"
		}

		if folder == "" {
			folder = user.HomeDir + "/Software"
		}

	}

	if backup != "" && folder != "" {

		filesystem, _ := fs.Sub(public.FS, ".")
		profile := structs.NewProfile(console, backup, folder, port)
		profile.Filesystem = &filesystem

		console.Clear("")
		console.Group("git-evac: Command-Line Arguments")
		console.Log("> Backup: " + backup)
		console.Log("> Folder: " + folder)
		console.Log("> Port:   " + strconv.FormatUint(uint64(port), 10))
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

			profile.Init()

			server.Dispatch(profile)
			server.DispatchRoutes(profile)

			result := server.Serve(profile)

			if result == false {
				console.Error("Port " + strconv.FormatUint(uint64(port), 10) + " is probably already in use?")
			}

			done <- result

		}()

		go func() {

			time.Sleep(1 * time.Second)

			console.Log("Opening WebView...")

			view := webview.New(true)
			view.SetTitle("Git Evac")
			view.SetSize(800, 600, webview.HintNone)
			view.Navigate("http://localhost:" + strconv.FormatUint(uint64(port), 10) + "/index.html")

			view.Run()

			done <- true

		}()

		select {
		case <-done:
			console.Log("The WebView or Server has been closed, exiting...")
		case <-signal_channel:
			console.Log("Received OS signal, exiting...")
		}

		// give webview time to cleanup
		time.Sleep(250 * time.Millisecond)

	}

}
