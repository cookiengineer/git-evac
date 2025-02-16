package server

import "git-evac/server/api"
import "git-evac/structs"
import "net/http"
import "strconv"

func Serve(profile *structs.Profile) bool {

	var result bool = false

	fsrv := http.FileServer(http.FS(*profile.Filesystem))
	http.Handle("/", fsrv)

	http.HandleFunc("/api/index", func(response http.ResponseWriter, request *http.Request) {
		api.Index(profile, request, response)
	})

	http.HandleFunc("/api/status/{organization}/{repository}", func(response http.ResponseWriter, request *http.Request) {
		api.Status(profile, request, response)
	})

	err1 := http.ListenAndServe(":"+strconv.FormatUint(uint64(profile.Settings.Port), 10), nil)

	if err1 == nil {
		result = true
	}

	return result

}
