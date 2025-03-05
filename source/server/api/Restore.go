package api

import "git-evac/console"
// import "git-evac/server/schemas"
import "git-evac/structs"
// import "encoding/json"
import "net/http"
import "os"
// import "os/exec"

func Restore(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodPatch {

		param1 := request.PathValue("owner")
		param2 := request.PathValue("repository")

		if profile.HasOwner(param1) {

			owner := profile.Owners[param1]

			if owner.HasRepository(param2) {

				// TODO: Move repository to .bak?
				console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusConflict)
				response.Write([]byte("{}"))

			} else {

				backup := profile.Settings.Backup
				file   := backup + "/" + owner.Name + "/" + param2 + ".tar.gz"

				stat, err0 := os.Stat(file)

				if err0 == nil && !stat.IsDir() {

					// TODO: Extract backup inside owner folder

				} else {

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusNotFound)
					response.Write([]byte("{}"))

				}

			}

		} else {

			folder := profile.Settings.Folder
			backup := profile.Settings.Backup
			file   := backup + "/" + param1 + "/" + param2 + ".tar.gz"

			if _, err := os.Stat(folder + "/" + param1); os.IsNotExist(err) {
				os.MkdirAll(backup + "/" + param1, 0755)
			}

			stat, err0 := os.Stat(file)

			if err0 == nil && !stat.IsDir() {

				// TODO: Extract backup inside owner folder

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

	// TODO:
	// Extract archive with
	// tar -xzvf whatever.tar.gz

	// Alternatively:
	// mkdir /tmp/sandbox
	// tar -xzvf whatever.tar.gz /tmp/sandbox
	// mv /tmp/sandbox/whatever /path/to/owner/whatever

}
