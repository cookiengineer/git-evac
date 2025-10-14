package api

import "git-evac/console"
import "git-evac/schemas"
import "git-evac/structs"
import "encoding/json"
import "net/http"
import "os"
import "os/exec"

func Restore(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodPatch {

		param1 := request.PathValue("owner")
		param2 := request.PathValue("repository")

		if profile.HasOwner(param1) {

			owner := profile.Owners[param1]

			if owner.HasRepository(param2) {

				backup := profile.Settings.Backup
				folder := profile.Settings.Folder

				stat0, err0 := os.Stat(backup + "/" + param1 + "/" + param2 + ".tar.gz")
				_, err1     := os.Stat(folder + "/" + param1 + "/" + param2 + ".bak")

				if err0 == nil && !stat0.IsDir() && os.IsNotExist(err1) {

					err2 := os.Rename(
						folder + "/" + param1 + "/" + param2,
						folder + "/" + param1 + "/" + param2 + ".bak",
					)

					if err2 == nil {

						cmd := exec.Command(
							"tar",
							"-xzvf",
							backup + "/" + owner.Name + "/" + param2 + ".tar.gz",
							param2,
						)
						cmd.Dir = owner.Folder

						_, err3 := cmd.Output()

						if err3 == nil {

							console.Log("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

							repo := owner.GetRepository(param2)
							repo.Status()

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusOK)

							payload, _ := json.MarshalIndent(schemas.Repository{
								Repository: *repo,
							}, "", "\t")
							response.Write(payload)

						} else {

							console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusInternalServerError)
							response.Write([]byte("{}"))

						}

					} else {

						console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusConflict)
						response.Write([]byte("{}"))

					}

				} else {

					console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusConflict)
					response.Write([]byte("{}"))

				}

			} else {

				backup := profile.Settings.Backup

				stat, err0 := os.Stat(backup + "/" + owner.Name + "/" + param2 + ".tar.gz")

				if err0 == nil && !stat.IsDir() {

					cmd := exec.Command(
						"tar",
						"-xzvf",
						backup + "/" + owner.Name + "/" + param2 + ".tar.gz",
						param2,
					)
					cmd.Dir = owner.Folder

					_, err1 := cmd.Output()

					if err1 == nil {

						repo := owner.GetRepository(param2)

						if repo != nil {

							repo.Status()

							console.Log("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusOK)

							payload, _ := json.MarshalIndent(schemas.Repository{
								Repository: *repo,
							}, "", "\t")
							response.Write(payload)

						} else {

							console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusInternalServerError)
							response.Write([]byte("{}"))

						}

					} else {

						console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusInternalServerError)
						response.Write([]byte("{}"))

					}

				} else {

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusNotFound)
					response.Write([]byte("{}"))

				}

			}

		} else {

			backup := profile.Settings.Backup
			folder := profile.Settings.Folder

			if _, err := os.Stat(folder + "/" + param1); os.IsNotExist(err) {
				os.MkdirAll(folder + "/" + param1, 0755)
			}

			stat, err0 := os.Stat(backup + "/" + param1 + "/" + param2 + ".tar.gz")

			if err0 == nil && !stat.IsDir() {

				cmd := exec.Command(
					"tar",
					"-xzvf",
					backup + "/" + param1 + "/" + param2 + ".tar.gz",
					param2,
				)
				cmd.Dir = folder + "/" + param1

				_, err1 := cmd.Output()

				if err1 == nil {

					owner := profile.GetOwner(param1, folder + "/" + param1)

					if owner != nil {

						repo := owner.GetRepository(param2)

						if repo != nil {

							repo.Status()

							console.Log("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusOK)

							payload, _ := json.MarshalIndent(schemas.Repository{
								Repository: *repo,
							}, "", "\t")
							response.Write(payload)

						} else {

							console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusInternalServerError)
							response.Write([]byte("{}"))

						}

					} else {

						console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusInternalServerError)
						response.Write([]byte("{}"))

					}

				} else {

					console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusInternalServerError)
					response.Write([]byte("{}"))

				}

			} else {

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusNotFound)
				response.Write([]byte("{}"))

			}

		}

	} else {

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{}"))

	}

}
