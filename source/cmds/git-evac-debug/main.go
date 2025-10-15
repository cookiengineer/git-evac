package main

import "git-evac/console"
import "git-evac/server"
import "git-evac/structs"
import "os"
import os_user "os/user"
import "strconv"
import "strings"

func main() {

	var folder string = ""
	var port uint16 = 3000

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

		fsys := os.DirFS("public")
		profile := structs.NewProfile("/tmp/backup", folder, port)
		profile.Filesystem = &fsys

		console.Clear()
		console.Group("git-evac-debug: Command-Line Arguments")
		console.Inspect(struct {
			Backup string
			Folder string
			Port   uint16
		}{
			Backup: "/tmp/backup",
			Folder: folder,
			Port:   port,
		})
		console.GroupEnd("")

		profile.Init()
		server.Dispatch(profile)
		server.DispatchHotReload(profile)

		if server.Serve(profile) == false {
			console.Error("Port " + strconv.FormatUint(uint64(port), 10) + " is already in use.")
			os.Exit(1)
		} else {
			os.Exit(0)
		}

	}

}
