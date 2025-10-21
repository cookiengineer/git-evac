package routes

import "git-evac/console"
import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func Status(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		owner := request.PathValue("owner")
		repository := request.PathValue("repository")

		repo := profile.GetRepository(owner, repository)

		if repo != nil {
			repo.Status()
		}

		console.Log("> " + request.Method + " /api/status/" + owner + "/" + repository + ": " + http.StatusText(http.StatusOK))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)

		payload, _ := json.MarshalIndent(schemas.Repository{
			Repository: *repo,
		}, "", "\t")
		response.Write(payload)

	} else {

		console.Error("> " + request.Method + " /api/status: " + http.StatusText(http.StatusMethodNotAllowed))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{}"))

	}

}

