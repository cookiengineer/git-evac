package actions

import "git-evac/structs"
import "errors"
import "os/exec"
import "sort"
import "strings"

func Push(profile *structs.Profile, owner_name string, repo_name string) error {

	if profile.HasRepositoryOwner(owner_name) {

		owner := profile.GetRepositoryOwner(owner_name)

		if owner != nil && owner.HasRepository(repo_name) {

			repository := owner.GetRepository(repo_name)
			messages := make(map[string]string)

			for remote, _ := range repository.Remotes {

				cmd := exec.Command("git", "push", remote, repository.CurrentBranch)
				folder := repository.Folder

				if strings.HasSuffix(folder, "/.git") {
					folder = folder[0:len(folder)-5]
				}

				cmd.Dir = folder

				msg0, err0 := cmd.Output()

				if err0 == nil {
					// Do Nothing
				} else {
					messages[remote] = strings.TrimSpace(string(msg0))
				}

			}

			if len(messages) == 0 {
				return nil
			} else {

				remotes := make([]string, 0)

				for remote, _ := range messages {
					remotes = append(remotes, remote)
				}

				sort.Strings(remotes)

				return errors.New("Repository \"" + owner_name + "/" + repo_name + "\"'s remotes \"" + strings.Join(remotes, ",") + " unreachable")

			}

		} else {
			return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" does not exist")
		}

	} else {
		return errors.New("Owner \"" + owner_name + "\" does not exist")
	}

}
