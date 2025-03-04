package api

// import "git-evac/console"
// import "git-evac/server/schemas"
import "git-evac/structs"
// import "encoding/json"
import "net/http"
// import "os/exec"

func Restore(profile *structs.Profile, request *http.Request, response http.ResponseWriter) {

	// TODO:
	// Extract archive with
	// tar -xzvf whatever.tar.gz

	// Alternatively:
	// mkdir /tmp/sandbox
	// tar -xzvf whatever.tar.gz /tmp/sandbox
	// mv /tmp/sandbox/whatever /path/to/owner/whatever

}
