package main

import "git-evac/server"
import "git-evac/structs"
import "encoding/json"
import "os"
import os_user "os/user"
import "strconv"
import "strings"

func main() {

	var config string = ""

	user, _ := os_user.Current()
	console := structs.NewConsole(os.Stderr, os.Stdout, 0)

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

		fsys := os.DirFS("public")

		profile := structs.NewProfile(console, settings)
		profile.Filesystem = &fsys

		console.Clear("")
		console.Group("git-evac-debug: Command-Line Arguments")

		if valid_config == true {
			console.Log("> Config: " + config)
		} else {
			console.Warn("> Invalid Config: " + config)
		}

		console.GroupEnd("git-evac-debug")

		profile.Refresh()

		server.Dispatch(profile)
		server.DispatchRoutes(profile)
		server.DispatchHotReload(profile)

		if server.Serve(profile) == false {
			console.Error("Port " + strconv.FormatUint(uint64(profile.Settings.Port), 10) + " is already in use.")
			os.Exit(1)
		} else {
			os.Exit(0)
		}

	}

}
