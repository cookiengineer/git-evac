package api

import "git-evac/console"
// import "git-evac/server/schemas"
import "git-evac/structs"
// import "encoding/json"
import "net/http"
// import "os/exec"

func Restore(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodPatch {

		param1 := request.PathValue("owner")
		param2 := request.PathValue("repository")

		if profile.HasOwner(param1) {

			owner := profile.Owners[param1]

			if owner.HasRepository(param2) {

				console.Error("> api.Restore(\"" + param1 + "\",\"" + param2 + "\")")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusConflict)
				response.Write([]byte("{}"))

				// TODO: Move repository to .bak?

			} else {

				// TODO: Extract backup inside owner folder

			}

		} else {

			// TODO: Create owner folder first
			// TODO: Extract backup inside owner folder

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
