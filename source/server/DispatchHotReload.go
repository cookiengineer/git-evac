package server

import "git-evac/console"
import "git-evac/structs"
import "bytes"
import "net/http"
import "os"
import "os/exec"
import "strings"

func DispatchHotReload(profile *structs.Profile) bool {

	var result bool = false

	fs := http.FS(*profile.Filesystem)
	fsrv := http.FileServer(fs)

	http.HandleFunc("/wasm_exec.js", func(response http.ResponseWriter, request *http.Request) {

		cwd, err0 := os.Getwd()

		if err0 == nil && strings.HasSuffix(cwd, "/source") {

			cmd1 := exec.Command("go", "env", "GOROOT")
			tmp1, err1 := cmd1.Output()

			if err1 == nil {

				goroot := strings.TrimSpace(string(tmp1))
				root := cwd[0:len(cwd)-7]
				exec_source := goroot + "/lib/wasm/wasm_exec.js"
				exec_output := root + "/source/public/wasm_exec.js"

				stat1, err1 := os.Stat(exec_source)

				if err1 == nil && !stat1.IsDir() {

					bytes, err2 := os.ReadFile(exec_source)

					if err2 == nil {

						err3 := os.WriteFile(exec_output, bytes, 0666)

						if err3 == nil {
							console.Log("> Renew /wasm_exec.js: success")
						} else {
							console.Error("> Renew /wasm_exec.js: failure")
							console.Error(err3.Error())
						}

					} else {
						console.Error("> Renew /wasm_exec.js: failure")
						console.Error(err2.Error())
					}

				} else {
					console.Error("> Renew /wasm_exec.js: failure")
					console.Error(err1.Error())
				}

			} else {
				console.Error("> Renew /wasm_exec.js: failure")
				console.Error(err0.Error())
			}

		}

		fsrv.ServeHTTP(response, request)

	})

	http.HandleFunc("/main.wasm", func(response http.ResponseWriter, request *http.Request) {

		cwd, err0 := os.Getwd()

		if err0 == nil && strings.HasSuffix(cwd, "/source") {

			root := cwd[0:len(cwd)-7]
			go_source := root + "/app/main.go"
			go_output := root + "/source/public/main.wasm"

			var stdout1 bytes.Buffer
			var stderr1 bytes.Buffer

			cmd1 := exec.Command(
				"env",
				"CGO_ENABLED=0",
				"GOOS=js",
				"GOARCH=wasm",
				"go",
				"build",
				"-o",
				go_output,
				go_source,
			)
			cmd1.Dir = root + "/app"

			cmd1.Stdout = &stdout1
			cmd1.Stderr = &stderr1

			err1 := cmd1.Run()

			if err1 == nil {

				console.Log("> Rebuild /main.wasm: success")

			} else {

				console.Error("> Rebuild /main.wasm: failure")

				lines := strings.Split(strings.TrimSpace(string(stderr1.Bytes())), "\n")

				for l := 0; l < len(lines); l++ {
					console.Error(lines[l])
				}

			}

		}

		fsrv.ServeHTTP(response, request)

	})

	return result

}
