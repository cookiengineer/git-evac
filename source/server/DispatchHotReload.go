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

		if request.Method == http.MethodGet {

			cwd, err0 := os.Getwd()

			if err0 == nil && strings.HasSuffix(cwd, "/source") {

				cmd1 := exec.Command("go", "env", "GOROOT")
				tmp1, err1 := cmd1.Output()

				if err1 == nil {

					goroot := strings.TrimSpace(string(tmp1))
					root := cwd[0:len(cwd)-7]
					exec_source := goroot + "/lib/wasm/wasm_exec.js"
					exec_output := root + "/source/public/wasm_exec.js"

					stat2, err2 := os.Stat(exec_source)

					if err2 == nil && !stat2.IsDir() {

						bytes, err3 := os.ReadFile(exec_source)

						if err3 == nil {

							err4 := os.WriteFile(exec_output, bytes, 0666)

							if err4 == nil {

								console.Log("> " + request.Method + " /wasm_exec.js: " + http.StatusText(http.StatusOK))

								response.Header().Set("Content-Type", "application/javascript")
								response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
								response.Header().Set("Pragma", "no-cache")
								response.Header().Set("Expires", "0")

								fsrv.ServeHTTP(response, request)

							} else {

								console.Error("> " + request.Method + " /wasm_exec.js: " + http.StatusText(http.StatusInternalServerError))
								console.Error(err4.Error())

								response.Header().Set("Content-Type", "application/javascript")
								response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
								response.Header().Set("Pragma", "no-cache")
								response.Header().Set("Expires", "0")

								response.WriteHeader(http.StatusInternalServerError)
								response.Write([]byte("// " + err4.Error()))

							}

						} else {

							console.Error("> " + request.Method + " /wasm_exec.js: " + http.StatusText(http.StatusInternalServerError))
							console.Error(err3.Error())

							response.Header().Set("Content-Type", "application/javascript")
							response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
							response.Header().Set("Pragma", "no-cache")
							response.Header().Set("Expires", "0")

							response.WriteHeader(http.StatusInternalServerError)
							response.Write([]byte("// " + err3.Error()))

						}

					} else {

						console.Error("> " + request.Method + " /wasm_exec.js: " + http.StatusText(http.StatusNotFound))

						if err2 != nil {
							console.Error(err2.Error())
						}

						response.Header().Set("Content-Type", "application/javascript")
						response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
						response.Header().Set("Pragma", "no-cache")
						response.Header().Set("Expires", "0")

						response.WriteHeader(http.StatusNotFound)

						if err2 != nil {
							response.Write([]byte("// " + err2.Error()))
						} else {
							response.Write([]byte(""))
						}

					}

				} else {

					console.Error("> " + request.Method + " /wasm_exec.js: " + http.StatusText(http.StatusInternalServerError))

					if err1 != nil {
						console.Error(err1.Error())
					}

					response.Header().Set("Content-Type", "application/javascript")
					response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
					response.Header().Set("Pragma", "no-cache")
					response.Header().Set("Expires", "0")

					response.WriteHeader(http.StatusInternalServerError)
					response.Write([]byte(""))

				}

			} else {

				console.Log("> " + request.Method + " /wasm_exec.js: " + http.StatusText(http.StatusOK))

				response.Header().Set("Content-Type", "application/javascript")
				response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
				response.Header().Set("Pragma", "no-cache")
				response.Header().Set("Expires", "0")

				fsrv.ServeHTTP(response, request)

			}

		} else {

			console.Error("> " + request.Method + " /wasm_exec.js: " + http.StatusText(http.StatusMethodNotAllowed))

			response.Header().Set("Content-Type", "application/javascript")
			response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
			response.Header().Set("Pragma", "no-cache")
			response.Header().Set("Expires", "0")

			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("[]"))

		}

	})

	http.HandleFunc("/main.wasm", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

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

					console.Log("> " + request.Method + " /main.wasm: " + http.StatusText(http.StatusOK))

					response.Header().Set("Content-Type", "application/wasm")
					response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
					response.Header().Set("Pragma", "no-cache")
					response.Header().Set("Expires", "0")

					fsrv.ServeHTTP(response, request)

				} else {

					console.Error("> " + request.Method + " /main.wasm: " + http.StatusText(http.StatusInternalServerError))

					response.Header().Set("Content-Type", "application/wasm")
					response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
					response.Header().Set("Pragma", "no-cache")
					response.Header().Set("Expires", "0")

					payload := []byte("// " + err1.Error())

					lines := strings.Split(strings.TrimSpace(string(stderr1.Bytes())), "\n")

					for l := 0; l < len(lines); l++ {
						console.Error("> " + lines[l])
						payload = append(payload, []byte(lines[l])...)
					}

					response.WriteHeader(http.StatusInternalServerError)
					response.Write(payload)

				}

			} else {

				console.Log("> " + request.Method + " /main.wasm: " + http.StatusText(http.StatusOK))

				response.Header().Set("Content-Type", "application/wasm")
				response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
				response.Header().Set("Pragma", "no-cache")
				response.Header().Set("Expires", "0")

				fsrv.ServeHTTP(response, request)

			}

		} else {

			console.Error("> " + request.Method + " /main.wasm: " + http.StatusText(http.StatusMethodNotAllowed))

			response.Header().Set("Content-Type", "application/wasm")
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte(""))

		}

	})

	return result

}
