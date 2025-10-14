package api

import "git-evac/console"
import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"
import "os/exec"

func Terminal(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		param1 := request.PathValue("owner")
		param2 := request.PathValue("repository")

		if profile.HasOwner(param1) {

			owner := profile.Owners[param1]

			if owner.HasRepository(param2) {

				repo := owner.Repositories[param2]

				cmd := exec.Command("kitty")
				cmd.Dir = repo.Folder[0:len(repo.Folder)-5]

				_, err := cmd.Output()

				if err == nil {

					repo.Status()

					console.Log("> api.Terminal(\"" + param1 + "\",\"" + param2 + "\")")

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusOK)

					payload, _ := json.MarshalIndent(schemas.Repository{
						Repository: *repo,
					}, "", "\t")
					response.Write(payload)

				} else {

					console.Error("> api.Terminal(\"" + param1 + "\",\"" + param2 + "\")")

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusInternalServerError)
					response.Write([]byte("{}"))

				}

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

