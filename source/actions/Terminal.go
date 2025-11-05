package actions

import "git-evac/structs"
import "errors"
import "os/exec"
import "strings"

func Terminal(profile *structs.Profile, owner_name string, repo_name string) error {

	if profile.HasRepositoryOwner(owner_name) {

		owner := profile.GetRepositoryOwner(owner_name)

		if owner != nil && owner.HasRepository(repo_name) {

			repository := owner.GetRepository(repo_name)

			cmd := exec.Command("kitty")
			folder := repository.Folder

			if strings.HasSuffix(folder, "/.git") {
				folder = folder[0:len(folder)-5]
			}

			cmd.Dir = folder

			_, err0 := cmd.Output()

			if err0 == nil {
				return nil
			} else {
				return err0
			}

		} else {
			return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" does not exist")
		}

	} else {
		return errors.New("Owner \"" + owner_name + "\" does not exist")
	}

}
