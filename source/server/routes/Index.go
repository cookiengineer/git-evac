package routes

import "git-evac/console"
import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func Index(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		payload, _ := json.MarshalIndent(schemas.Repositories{
			Owners: profile.Owners,
		}, "", "\t")

		console.Log("> " + request.Method + " /api/index: " + http.StatusText(http.StatusOK))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(payload)

	} else {

		console.Error("> " + request.Method + " /api/index: " + http.StatusText(http.StatusMethodNotAllowed))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("[]"))

	}

}
