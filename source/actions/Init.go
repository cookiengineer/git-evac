package actions

import "git-evac/routes"
import "git-evac/structs"
import "net/http"
import "strconv"

func Init(profile *structs.Profile) bool {

	profile.Init()

	result := false

	fsrv := http.FileServer(http.FS(*profile.Filesystem))
	http.Handle("/", fsrv)

	http.HandleFunc("/api/index", func(response http.ResponseWriter, request *http.Request) {
		routes.Index(profile, request, response)
	})

	http.HandleFunc("/api/status/{organization}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		routes.Status(profile, request, response)
	})

	// http.HandleFunc("/api/config/{orga}/{repo}", func(response http.ResponseWriter, request *http.Request) {

	// 	// TODO: GET config
	// 	// TODO: PATCH config
	// 	// TODO: POST config

	// })

	// http.HandleFunc("/api/fetch/{orga}/{repo}", func(response http.ResponseWriter, request *http.Request) {

	// 	// TODO: Execute fetch action, update database

	// })

	// http.HandleFunc("/api/push/{orga}/{repo}/{remote}", func(response http.ResponseWriter, request *http.Request) {

	// 	// TODO: Execute push action, update database

	// })

	err1 := http.ListenAndServe(":"+strconv.FormatUint(uint64(profile.Settings.Port), 10), nil)

	if err1 == nil {
		result = true
	}

	return result

}
