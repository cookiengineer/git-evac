package actions

import "git-evac/structs"
import "bytes"
import "errors"
import "os/exec"
import "sort"
import "strings"

func Pull(profile *structs.Profile, owner_name string, repo_name string) error {

	if profile.HasOwner(owner_name) {

		owner := profile.GetOwner(owner_name, profile.Settings.Folder + "/" + owner_name)

		if owner.HasRepository(repo_name) {

			repository := owner.GetRepository(repo_name)
			folder := repository.Folder

			if strings.HasSuffix(folder, "/.git") {
				folder = folder[0:len(folder)-5]
			}

			// 1. git fetch --all
			var stdout_fetch bytes.Buffer
			var stderr_fetch bytes.Buffer

			cmd_fetch := exec.Command("git", "fetch", "--all")
			cmd_fetch.Dir = folder
			cmd_fetch.Stdout = &stdout_fetch
			cmd_fetch.Stderr = &stderr_fetch

			err_fetch := cmd_fetch.Run()

			if err_fetch == nil {

				// 2. git diff --name-status master..origin/master
				var stdout_diff bytes.Buffer
				var stderr_diff bytes.Buffer

				cmd_diff := exec.Command(
					"git",
					"diff",
					"--name-status",
					repository.CurrentBranch + "..origin/" + repository.CurrentBranch,
				)
				cmd_diff.Dir = folder
				cmd_diff.Stdout = &stdout_diff
				cmd_diff.Stderr = &stderr_diff

				err_diff := cmd_diff.Run()

				if err_diff == nil {

					diff_status := strings.TrimSpace(stdout_diff.String())

					if diff_status == "" {
						// No remote changes, no merge necessary
						return nil
					} else {

						// Remote changes could look like this:
						// M       CONFLICT.txt
						// TODO: M might be an indicator of potential merge conflict

						// 3. git merge --no-edit origin/master
						var stdout_merge bytes.Buffer
						var stderr_merge bytes.Buffer

						cmd_merge := exec.Command(
							"git",
							"merge",
							"--no-edit",
							"origin/" + repository.CurrentBranch,
						)
						cmd_merge.Dir = folder
						cmd_merge.Stdout = &stdout_merge
						cmd_merge.Stderr = &stderr_merge

						err_merge := cmd_merge.Run()

						if err_merge == nil {
							// Remote changes, no local changes, merge successful
							return nil
						} else {

							conflicts := make(map[string]bool)
							lines := strings.Split(strings.TrimSpace(stdout_merge.String()), "\n")

							for _, line := range lines {

								if strings.HasPrefix(line, "CONFLICT") && strings.Contains(line, "Merge conflict in ") {
									filepath := strings.TrimSpace(line[strings.Index(line, "Merge conflict in ")+18:])
									conflicts[filepath] = true
								}

							}

							if len(conflicts) == 0 {
								return nil
							} else {

								files := make([]string, 0)

								for file, _ := range conflicts {
									files = append(files, file)
								}

								sort.Strings(files)

								return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" failed to merge changes for files \"" + strings.Join(files, "\", \"") + "\"")

							}

						}

					}

				} else {
					return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" failed to diff against origin remote")
				}

			} else {
				return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" failed to fetch all remotes")
			}

		} else {
			return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" does not exist")
		}

	} else {
		return errors.New("Owner \"" + owner_name + "\" does not exist")
	}

}
