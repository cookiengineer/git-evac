package server

import "git-evac/server/api"
import "git-evac/structs"
import "net/http"
import "io"

func Dispatch(profile *structs.Profile) bool {

	var result bool = false

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

	// Canonical to /api/terminal
	http.HandleFunc("/api/fix/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		api.Terminal(profile, request, response)
	})

	http.HandleFunc("/api/terminal/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		api.Terminal(profile, request, response)
	})

	http.HandleFunc("/api/clone/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		api.Clone(profile, request, response)
	})

	// TODO: Diff API
	// http.HandleFunc("/api/diff/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	api.Diff(profile, request, response)
	// })

	// TODO: Commit API
	// http.HandleFunc("/api/commit/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	api.Commit(profile, request, response)
	// })

	// TODO: Pull API
	// http.HandleFunc("/api/pull/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	api.Pull(profile, request, response)
	// })

	http.HandleFunc("/api/push/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		api.Push(profile, request, response)
	})

	http.HandleFunc("/api/index", func(response http.ResponseWriter, request *http.Request) {

		for _, owner := range profile.Owners {

			for _, repo := range owner.Repositories {
				repo.Status()
			}

		}

		api.Index(profile, request, response)

	})

	http.HandleFunc("/api/status/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		api.Status(profile, request, response)
	})

	http.HandleFunc("/api/settings", func(response http.ResponseWriter, request *http.Request) {
		api.Settings(profile, request, response)
	})

	return result

}
