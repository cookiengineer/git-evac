package routes

import "git-evac/console"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func Status(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		param1 := request.PathValue("organization")
		param2 := request.PathValue("repository")

		if profile.HasUser(param1) {

			user := profile.Users[param1]

			if user.HasRepository(param2) {

				repo := user.Repositories[param2]

				if repo.Status() {
					console.Log("/api/status @" + param1 + "/" + param2)
				}

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)

				payload, _ := json.MarshalIndent(repo, "", "\t")
				response.Write(payload)

			} else {

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusNotFound)
				response.Write([]byte("{}"))

			}

		} else if profile.HasOrganization(param1) {

			orga := profile.Organizations[param1]

			if orga.HasRepository(param2) {

				repo := orga.Repositories[param2]

				if repo.Status() {
					console.Log("/api/status " + param1 + "/" + param2)
				}

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)

				payload, _ := json.MarshalIndent(repo, "", "\t")
				response.Write(payload)

			} else {

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusNotFound)
				response.Write([]byte("{}"))

			}

		} else {

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte("{}"))

		}

	} else {

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{}"))

	}

}
