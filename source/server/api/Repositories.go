package api

import "git-evac/server/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func Repositories(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		payload, err := json.MarshalIndent(schemas.Repositories{
			Owners: profile.Owners,
		}, "", "\t")

		if err == nil {

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			response.Write(payload)

		} else {

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
