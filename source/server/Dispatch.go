package server

import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "io"
import "net/http"
import "os"

func Dispatch(profile *structs.Profile) bool {

	var result bool

	fs := http.FS(*profile.Filesystem)
	fsrv := http.FileServer(fs)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		if request.URL.Path == "/" {

			response.Header().Set("Location", "/index.html")
			response.WriteHeader(http.StatusSeeOther)
			response.Write([]byte{})

		} else if request.URL.Path == "/index.html" {

			directives := []string{
				"default-src 'self' 'unsafe-eval' 'wasm-unsafe-eval'",
				"script-src 'self' 'unsafe-eval' 'wasm-unsafe-eval'",
				"script-src-elem 'self'",
				"worker-src 'self'",
				"frame-src * 'self'",
				"connect-src * 'self'",
			}

			// WebASM's JSON.parse/stringify requires wasm-unsafe-eval directive
			response.Header().Set("Access-Control-Allow-Origin", "*")

			for d := 0; d < len(directives); d++ {
				response.Header().Set("Content-Security-Policy", directives[d])
			}

			file, err := fs.Open("/index.html")

			if err == nil {

				buffer := make([]byte, 0)

				for {

					bytes := make([]byte, 1024)
					num, err := file.Read(bytes)

					if err == nil {
						buffer = append(buffer, bytes[0:num]...)
					} else if err == io.EOF {
						buffer = append(buffer, bytes[0:num]...)
						break
					}

				}

				response.Write(buffer)

			}

		} else {
			fsrv.ServeHTTP(response, request)
		}

	})

	http.HandleFunc("/FS.go", func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(""))
	})

	http.HandleFunc("/api/settings", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			payload, _ := json.MarshalIndent(schemas.Settings{
				Settings: profile.Settings,
			}, "", "\t")

			profile.Console.Log("> GET /api/settings:" + http.StatusText(http.StatusOK))

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			response.Write(payload)

		} else if request.Method == http.MethodPost {

			bytes0, err0 := io.ReadAll(request.Body)

			if err0 == nil {

				var schema schemas.Settings

				err1 := json.Unmarshal(bytes0, &schema)

				if err1 == nil && schema.IsValid() {

					profile.Settings.Backup = schema.Settings.Backup
					profile.Settings.Folder = schema.Settings.Folder
					profile.Settings.Port = schema.Settings.Port
					profile.Settings.Organizations = schema.Settings.Organizations

					stat2, err2 := os.Stat(profile.Settings.Folder)

					if err2 == nil && stat2.IsDir() {

						payload, _ := json.MarshalIndent(schemas.Settings{
							Settings: profile.Settings,
						}, "", "\t")

						err3 := os.WriteFile(profile.Settings.Folder+"/git-evac.json", payload, 0666)

						if err3 == nil {

							profile.Console.Log("> POST /api/settings: " + http.StatusText(http.StatusOK))

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusOK)
							response.Write(payload)

						} else {

							profile.Console.Error("> POST /api/settings: " + http.StatusText(http.StatusInternalServerError))
							profile.Console.Error("> " + err3.Error())

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusInternalServerError)
							response.Write([]byte("{}"))

						}

					} else {

						profile.Console.Error("> POST /api/settings: " + http.StatusText(http.StatusConflict))
						profile.Console.Error("> " + err2.Error())

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusConflict)
						response.Write([]byte("{}"))

					}

				} else {

					profile.Console.Error("> POST /api/settings: " + http.StatusText(http.StatusBadRequest))

					if err1 != nil {
						profile.Console.Error("> " + err1.Error())
					}

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte("{}"))

				}

			} else {

				profile.Console.Error("> POST /api/settings: " + http.StatusText(http.StatusBadRequest))

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("{}"))

			}

		} else {

			profile.Console.Error("> " + request.Method + " /api/settings: " + http.StatusText(http.StatusMethodNotAllowed))

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("[]"))

		}

	})

	return result

}
