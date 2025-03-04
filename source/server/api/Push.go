package api

import "git-evac/console"
import "git-evac/server/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"
import "os/exec"
import "strings"

func Push(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		param1 := request.PathValue("owner")
		param2 := request.PathValue("repository")

		if profile.HasOwner(param1) {

			owner := profile.Owners[param1]

			if owner.HasRepository(param2) {

				repo := owner.Repositories[param2]
				repo.Status()

				errors := make(map[string]string)

				for name, _ := range repo.Remotes {

					cmd := exec.Command("git", "push", name, repo.CurrentBranch)
					cmd.Dir = repo.Folder[0:len(repo.Folder)-5]

					output, err := cmd.Output()

					if err == nil {
						// Do Nothing
					} else {
						errors[name] = strings.TrimSpace(string(output))
					}

				}

				if len(errors) == 0 {

					console.Log("> api.Push(\"" + param1 + "\",\"" + param2 + "\")")

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusOK)

					payload, _ := json.MarshalIndent(schemas.Repository{
						Repository: *repo,
					}, "", "\t")
					response.Write(payload)

				} else {

					console.Error("> api.Push(\"" + param1 + "\",\"" + param2 + "\")")

					for remote, message := range errors {
						console.Error("-> git push " + remote + " failed:")
						console.Error(message)
					}

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusInternalServerError)
					response.Write([]byte("{}"))

				}

			} else {

				console.Error("> api.Push(\"" + param1 + "\",\"" + param2 + "\")")

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
