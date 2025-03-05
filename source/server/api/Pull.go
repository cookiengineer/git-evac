package api

import "git-evac/console"
// import "git-evac/server/schemas"
import "git-evac/structs"
// import "encoding/json"
import "net/http"
import "os/exec"

func Pull(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		param1 := request.PathValue("owner")
		param2 := request.PathValue("repository")

		if profile.HasOwner(param1) {

			owner := profile.Owners[param1]

			if owner.HasRepository(param2) {

				repo := owner.Repositories[param2]

				cmd1 := exec.Command("git", "fetch", "--all")
				cmd1.Dir = repo.Folder[0:len(repo.Folder)-5]

				_, err1 := cmd1.Output()

				if err1 == nil {

					// TODO
					console.Warn("TODO: Implement git diff")
					console.Warn("TODO: Implement git merge")

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

