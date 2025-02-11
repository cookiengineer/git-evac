package actions

import "git-evac/handlers"
import "git-evac/routes"
import "git-evac/structs"
import "io/fs"
import "net/http"
import "strconv"

func Serve(filesystem fs.FS, port int) bool {

	database := structs.NewDatabase()
	result := false

	fsrv := http.FileServer(http.FS(filesystem))
	http.Handle("/", fsrv)

	http.HandleFunc("/api/index", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {
			routes.Index(nil, &database, request, response)
		} else {
			handlers.RespondWith(request, response, http.StatusMethodNotAllowed)
		}

	})

	http.HandleFunc("/api/status/{orga}/{repo}", func(response http.ResponseWriter, request *http.Request) {

		// if request.Method == http.MethodGet {

		// 	config := database.GetConfig(
		// 		request.PathValue("orga"),
		// 		request.PathValue("repo"),
		// 	)

		// 	if config != nil {
		// 		routes.Status(config, &database, request, response)
		// 	} else {
		// 		handlers.RespondWith(request, response, http.StatusNotFound)
		// 	}

		// } else {
		// 	handlers.RespondWith(request, response, http.StatusMethodNotAllowed)
		// }

	})

	http.HandleFunc("/api/config/{orga}/{repo}", func(response http.ResponseWriter, request *http.Request) {

		// TODO: GET config
		// TODO: PATCH config
		// TODO: POST config

	})

	http.HandleFunc("/api/fetch/{orga}/{repo}", func(response http.ResponseWriter, request *http.Request) {

		// TODO: Execute fetch action, update database

	})

	http.HandleFunc("/api/push/{orga}/{repo}/{remote}", func(response http.ResponseWriter, request *http.Request) {

		// TODO: Execute push action, update database

	})

	err1 := http.ListenAndServe(":"+strconv.Itoa(port), nil)

	if err1 == nil {
		result = true
	}

	return result

}
