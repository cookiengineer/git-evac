package server

import "git-evac/server/api"
import "git-evac/structs"
import "net/http"
import "io"
import "strconv"

func Serve(profile *structs.Profile) bool {

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

	http.HandleFunc("/api/terminal/open/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		api.TerminalOpen(profile, request, response)
	})

	// http.HandleFunc("/api/git/clone/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	api.GitClone(profile, request, response)
	// })

	// http.HandleFunc("/api/git/commit/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	api.GitCommit(profile, request, response)
	// })

	// http.HandleFunc("/api/git/pull/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	api.GitPull(profile, request, response)
	// })

	// http.HandleFunc("/api/git/push/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	api.GitPush(profile, request, response)
	// })

	http.HandleFunc("/api/repositories", func(response http.ResponseWriter, request *http.Request) {

		for _, owner := range profile.Owners {

			for _, repo := range owner.Repositories {
				repo.Status()
			}

		}

		api.Repositories(profile, request, response)

	})

	http.HandleFunc("/api/repositories/status/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		api.RepositoryStatus(profile, request, response)
	})

	http.HandleFunc("/api/settings", func(response http.ResponseWriter, request *http.Request) {
		api.Settings(profile, request, response)
	})

	err1 := http.ListenAndServe(":"+strconv.FormatUint(uint64(profile.Settings.Port), 10), nil)

	if err1 == nil {
		result = true
	}

	return result

}
