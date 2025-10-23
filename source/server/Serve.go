package server

import "git-evac/structs"
import "net/http"
import "strconv"

func Serve(profile *structs.Profile) bool {

	var result bool

	profile.Console.Group("Serve()")
	profile.Console.Log("Listening on http://localhost:" + strconv.FormatUint(uint64(profile.Settings.Port), 10))

	err1 := http.ListenAndServe(":"+strconv.FormatUint(uint64(profile.Settings.Port), 10), nil)

	if err1 == nil {
		result = true
	}

	profile.Console.GroupEnd("Serve()")

	return result

}
