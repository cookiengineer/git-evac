package actions

import "git-evac/structs"
import "errors"
import "os"
import "os/exec"

func Backup(profile *structs.Profile, owner_name string, repo_name string) error {

	if profile.HasRepositoryOwner(owner_name) {

		owner := profile.GetRepositoryOwner(owner_name)

		if owner != nil && owner.HasRepository(repo_name) {

			repository := owner.GetRepository(repo_name)

			repository.Status()

			if _, err := os.Stat(profile.Settings.Backup + "/" + owner.Name); os.IsNotExist(err) {
				os.MkdirAll(profile.Settings.Backup + "/" + owner.Name, 0755)
			}

			stat0, err0 := os.Stat(profile.Settings.Backup + "/" + owner.Name)

			if err0 == nil && stat0.IsDir() {

				cmd := exec.Command(
					"tar",
					"-czvf",
					profile.Settings.Backup + "/" + owner.Name + "/" + repository.Name + ".tar.gz",
					repository.Name,
				)
				cmd.Dir = owner.Folder

				_, err1 := cmd.Output()

				if err1 == nil {
					return nil
				} else {
					return err1
				}

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
