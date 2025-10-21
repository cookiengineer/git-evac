package actions

import "git-evac/structs"
import "errors"
import "os"
import "os/exec"

func Restore(profile *structs.Profile, owner_name string, repo_name string) error {

	if profile.HasOwner(owner_name) {

		owner := profile.GetOwner(owner_name, profile.Settings.Folder + "/" + owner_name)

		if owner.HasRepository(repo_name) {

			repository := owner.GetRepository(repo_name)

			stat0, err0 := os.Stat(profile.Settings.Backup + "/" + owner.Name + "/" + repository.Name + ".tar.gz")
			_, err1 := os.Stat(profile.Settings.Folder + "/" + owner.Name + "/" + repository.Name + ".bak")

			if err0 == nil && !stat0.IsDir() && os.IsNotExist(err1) {

				err2 := os.Rename(
					profile.Settings.Folder + "/" + owner.Name + "/" + repository.Name,
					profile.Settings.Folder + "/" + owner.Name + "/" + repository.Name + ".bak",
				)

				if err2 == nil {

					cmd := exec.Command(
						"tar",
						"-xzvf",
						profile.Settings.Backup + "/" + owner.Name + "/" + repository.Name + ".tar.gz",
						repository.Name,
					)
					cmd.Dir = owner.Folder

					_, err3 := cmd.Output()

					if err3 == nil {
						return nil
					} else {
						return err3
					}

				} else {
					return err2
				}

			} else {

				if err0 != nil || (err0 == nil && stat0.IsDir()) {
					return errors.New("Repository Backup \"" + owner_name + "/" + repo_name + ".tar.gz\" does not exist")
				} else {
					return errors.New("Repository Backup \"" + owner_name + "/" + repo_name + ".bak\" already exists")
				}

			}

		} else {

			stat0, err0 := os.Stat(profile.Settings.Backup + "/" + owner_name + "/" + repo_name + ".tar.gz")

			if err0 == nil && !stat0.IsDir() {

				cmd := exec.Command(
					"tar",
					"-xzvf",
					profile.Settings.Backup + "/" + owner.Name + "/" + repo_name + ".tar.gz",
					repo_name,
				)
				cmd.Dir = owner.Folder

				_, err3 := cmd.Output()

				if err3 == nil {
					return nil
				} else {
					return err3
				}

			} else {
				return errors.New("Repository Backup \"" + owner_name + "/" + repo_name + ".tar.gz\" does not exist")
			}

		}

	} else {
		return errors.New("Owner \"" + owner_name + "\" does not exist")
	}

}
