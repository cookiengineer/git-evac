package server

import "git-evac/server/routes"
import "git-evac/structs"
import "net/http"

func DispatchRoutes(profile *structs.Profile) bool {

	var result bool

	// GET /api/index
	http.HandleFunc("/api/index", func(response http.ResponseWriter, request *http.Request) {
		profile.Refresh()
		routes.Index(profile, request, response)
	})

	// GET /api/backup || POST /api/backup
	http.HandleFunc("/api/backup/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		routes.Backup(profile, request, response)
	})

	// TODO: GET /api/clone
	// http.HandleFunc("/api/clone/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	routes.Clone(profile, request, response)
	// })

	// TODO: POST /api/commit
	// http.HandleFunc("/api/commit/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	routes.Commit(profile, request, response)
	// })

	// TODO: GET /api/diff
	// http.HandleFunc("/api/diff/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	routes.Diff(profile, request, response)
	// })

	// PATCH /api/restore
	http.HandleFunc("/api/restore/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		routes.Restore(profile, request, response)
	})

	// GET /api/fix is canonical to GET /api/terminal
	http.HandleFunc("/api/fix/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		routes.Terminal(profile, request, response)
	})

	// TODO: PATCH /api/pull
	// http.HandleFunc("/api/pull/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
	// 	routes.Pull(profile, request, response)
	// })

	// GET /api/push
	http.HandleFunc("/api/push/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		routes.Push(profile, request, response)
	})

	// GET /api/status
	http.HandleFunc("/api/status/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		routes.Status(profile, request, response)
	})

	// GET /api/terminal
	http.HandleFunc("/api/terminal/{owner}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		routes.Terminal(profile, request, response)
	})

	return result

}
