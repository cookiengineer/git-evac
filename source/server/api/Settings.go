package api

import "git-evac/console"
import "git-evac/server/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func Settings(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		payload, err := json.MarshalIndent(schemas.Settings{
			Settings: profile.Settings,
		}, "", "\t")

		if err == nil {

			console.Log("> api.Settings()")

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			response.Write(payload)

		} else {

			console.Error("> api.Settings()")

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("[]"))

		}

	} else {

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("[]"))

	}

}
