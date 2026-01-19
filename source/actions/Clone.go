package actions

import "git-evac/structs"
import "bytes"
import "errors"
import "os"
import "os/exec"
import "path/filepath"
import "strings"

func Clone(profile *structs.Profile, owner_name string, repo_name string) error {

	if profile.HasRepository(owner_name, repo_name) {

		repository := profile.GetRepository(owner_name, repo_name)

		if repository != nil {

			stat, err0 := os.Stat(repository.Folder)

			if err0 == nil && stat.IsDir() {

				// Repository exists, pull instead
				return Pull(profile, owner_name, repo_name)

			} else if os.IsNotExist(err0) {

				folder := repository.Folder

				if strings.HasSuffix(folder, "/.git") {
					folder = folder[0:len(folder)-5]
				}

				parent := filepath.Dir(folder)

				// Repository does not exist, clone from origin
				origin, ok := repository.Remotes["origin"]

				if ok == true {

					var stdout_clone bytes.Buffer
					var stderr_clone bytes.Buffer

					cmd_clone := exec.Command(
						"git",
						"clone",
						"--single-branch",
						origin.URL,
						"./" + repository.Name,
					)
					cmd_clone.Dir = parent
					cmd_clone.Stdout = &stdout_clone
					cmd_clone.Stderr = &stderr_clone

					err_clone := cmd_clone.Run()

					if err_clone == nil {
						return nil
					} else {
						return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" failed to clone from origin remote")
					}

				} else {
					return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" has no origin remote")
				}

			} else {
				return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" does not exist")
			}

		} else {
			return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" does not exist")
		}

	} else {
		return errors.New("Owner \"" + owner_name + "\" does not exist")
	}

}
