package routes

import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func Backups(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		payload, _ := json.MarshalIndent(schemas.Backups{
			Owners: profile.Backups,
		}, "", "\t")

		profile.Console.Log("> " + request.Method + " /api/backups: " + http.StatusText(http.StatusOK))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(payload)

	} else {

		profile.Console.Error("> " + request.Method + " /api/backups: " + http.StatusText(http.StatusMethodNotAllowed))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("[]"))

	}

}
