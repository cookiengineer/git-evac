package routes

import "git-evac/actions"
import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func Clone(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodPost {

		owner := request.PathValue("owner")
		repository := request.PathValue("repository")

		err := actions.Clone(profile, owner, repository)

		if err == nil {

			repo := profile.GetRepository(owner, repository)
			repo.Status()

			profile.Console.Log("> " + request.Method + " /api/clone/" + owner + "/" + repository + ": " + http.StatusText(http.StatusOK))

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)

			payload, _ := json.MarshalIndent(schemas.Repository{
				Repository: *repo,
			}, "", "\t")
			response.Write(payload)

		} else {

			profile.Console.Error("> " + request.Method + " /api/clone/" + owner + "/" + repository + ": " + http.StatusText(http.StatusInternalServerError))
			profile.Console.Error("> " + err.Error())

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("{}"))

		}

	} else {

		profile.Console.Error("> " + request.Method + " /api/clone: " + http.StatusText(http.StatusMethodNotAllowed))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{}"))

	}

}
