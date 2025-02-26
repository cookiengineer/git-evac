package api

import "git-evac/console"
import "git-evac/server/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func RepositoryStatus(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		param1 := request.PathValue("owner")
		param2 := request.PathValue("repository")

		if profile.HasOwner(param1) {

			owner := profile.Owners[param1]

			if owner.HasRepository(param2) {

				repo := owner.Repositories[param2]

				if repo.Status() {
					console.Log("/api/repositories/status/" + param1 + "/" + param2)
				}

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)

				payload, _ := json.MarshalIndent(schemas.Repository{
					Repository: *repo,
				}, "", "\t")
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
