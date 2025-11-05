package routes

import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func Repositories(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		payload, _ := json.MarshalIndent(schemas.Repositories{
			Owners: profile.Repositories,
		}, "", "\t")

		profile.Console.Log("> " + request.Method + " /api/repositories: " + http.StatusText(http.StatusOK))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(payload)

	} else {

		profile.Console.Error("> " + request.Method + " /api/repositories: " + http.StatusText(http.StatusMethodNotAllowed))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("[]"))

	}

}
