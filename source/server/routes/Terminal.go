package routes

import "git-evac/actions"
import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"

func Terminal(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		owner := request.PathValue("owner")
		repository := request.PathValue("repository")

		err := actions.Terminal(profile, owner, repository)

		if err == nil {

			repo := profile.GetRepository(owner, repository)
			repo.Status()

			profile.Console.Log("> " + request.Method + " /api/terminal/" + owner + "/" + repository + ": " + http.StatusText(http.StatusOK))

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)

			payload, _ := json.MarshalIndent(schemas.Repository{
				Repository: *repo,
			}, "", "\t")
			response.Write(payload)

		} else {

			profile.Console.Error("> " + request.Method + " /api/terminal/" + owner + "/" + repository + ": " + http.StatusText(http.StatusInternalServerError))
			profile.Console.Error("> " + err.Error())

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("{}"))

		}

	} else {

		profile.Console.Error("> " + request.Method + " /api/terminal: " + http.StatusText(http.StatusMethodNotAllowed))

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{}"))

	}

}
