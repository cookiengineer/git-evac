package api

import "git-evac/console"
import "git-evac/server/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"
import "os"
import "os/exec"

func Backup(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet || request.Method == http.MethodPost {

		param1 := request.PathValue("owner")
		param2 := request.PathValue("repository")

		if profile.HasOwner(param1) {

			owner := profile.Owners[param1]

			if owner.HasRepository(param2) {

				repo := owner.Repositories[param2]

				if repo.Status() {

					backup := profile.Settings.Backup

					if _, err := os.Stat(backup + "/" + owner.Name); os.IsNotExist(err) {
						os.MkdirAll(backup + "/" + owner.Name, 0755)
					}

					stat, err0 := os.Stat(backup + "/" + owner.Name)

					if err0 == nil && stat.IsDir() {

						cmd := exec.Command(
							"tar",
							"-czvf",
							backup + "/" + owner.Name + "/" + repo.Name + ".tar.gz",
							repo.Name,
						)
						cmd.Dir = owner.Folder

						_, err1 := cmd.Output()

						if err1 == nil {

							console.Log("> api.Backup(\"" + param1 + "\",\"" + param2 + "\")")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusOK)

							payload, _ := json.MarshalIndent(schemas.Repository{
								Repository: *repo,
							}, "", "\t")
							response.Write(payload)

						} else {

							console.Error("> api.Backup(\"" + param1 + "\",\"" + param2 + "\")")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusInternalServerError)
							response.Write([]byte("{}"))

						}

					} else {

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusConflict)
						response.Write([]byte("{}"))

					}

				} else {

					console.Error("> api.Backup(\"" + param1 + "\",\"" + param2 + "\")")

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
			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte("{}"))

		}

	} else {

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{}"))

	}

}
