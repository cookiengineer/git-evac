package api

import "git-evac/console"
import "git-evac/schemas"
import "git-evac/structs"
import utils_remotes "git-evac/utils/remotes"
import "bytes"
import "encoding/json"
import "net/http"
import "os"
import "os/exec"
import "strings"

func Clone(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		param1 := request.PathValue("owner")
		param2 := request.PathValue("repository")

		if profile.HasOwner(param1) {

			owner := profile.Owners[param1]

			if owner.HasRepository(param2) == false {

				remote := utils_remotes.GuessOrigin(owner, param2)
				_, err0 := os.Stat(owner.Folder + "/" + param2)

				if remote != "" && os.IsNotExist(err0) {

					folder := owner.Folder + "/" + param2

					var stdout1 bytes.Buffer
					var stderr1 bytes.Buffer

					cmd1 := exec.Command(
						"git",
						"clone",
						remote,
						folder,
					)
					cmd1.Dir = owner.Folder

					cmd1.Stdout = &stdout1
					cmd1.Stderr = &stderr1

					err1 := cmd1.Run()

					if err1 == nil {

						console.Log("> api.Clone(\"" + param1 + "\",\"" + param2 + "\") from \"" + remote + "\"")

						repo := owner.GetRepository(param2)

						if repo != nil {

							repo.Status()

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusOK)

							payload, _ := json.MarshalIndent(schemas.Repository{
								Repository: *repo,
							}, "", "\t")
							response.Write(payload)

						} else {

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusInternalServerError)
							response.Write([]byte("{}"))

						}

					} else {

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusFailedDependency)
						response.Write([]byte("{}"))

						console.Error("> api.Clone(\"" + param1 + "\",\"" + param2 + "\") from \"" + remote + "\"")

						lines := strings.Split(strings.TrimSpace(string(stderr1.Bytes())), "\n")

						for l := 0; l < len(lines); l++ {
							console.Error(lines[l])
						}

					}

				} else {

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusConflict)
					response.Write([]byte("{}"))

				}

			} else {

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusConflict)
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
