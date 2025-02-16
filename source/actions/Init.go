package actions

import "git-evac/server"
import "git-evac/structs"

func Init(profile *structs.Profile) bool {

	var result bool = false

	profile.Init()

	if server.Serve(profile) == true {
		result = true
	}

	return result

}
