package routes

import "git-evac/structs"
import "encoding/json"
import "net/http"

func Index(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		payload, err := json.MarshalIndent(struct {
			Users         map[string]*structs.User         `json:"users"`
			Organizations map[string]*structs.Organization `json:"organizations"`
		}{
			Users:         profile.Users,
			Organizations: profile.Organizations,
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
