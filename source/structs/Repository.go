package structs

import "git-evac/git"
import "os"
import "os/exec"
import "slices"
import "strings"

type Repository struct {
	Name             string             `json:"name"`
	Folder           string             `json:"folder"` // /path/to/.git
	Branches         []string           `json:"branches"`
	Remotes          map[string]*Remote `json:"remotes"`
	CurrentBranch    string             `json:"current_branch"`
	CurrentRemote    string             `json:"current_remote"`
	IsPublic         bool               `json:"is_public"`
	HasLocalChanges  bool               `json:"has_local_changes"`
	HasRemoteChanges bool               `json:"has_remote_changes"`
	Identity         string             `json:"identity"`
}

func NewRepository(name string, folder string) Repository {

	var repo Repository

	repo.Name = name
	repo.Folder = folder
	repo.Branches = make([]string, 0)
	repo.Remotes = make(map[string]*Remote)

	repo.CurrentBranch = "master"
	repo.CurrentRemote = "origin"
	repo.Identity = ""

	repo.Status()

	return repo

}

func (repo *Repository) Status() bool {

	var result bool = false

	stat, err0 := os.Stat(repo.Folder)

	if err0 == nil && stat.IsDir() && strings.HasSuffix(repo.Folder, "/.git") {

		cmd1 := exec.Command("git", "status", "--branch", "--short")
		cmd1.Dir = repo.Folder[0:len(repo.Folder)-5]

		buffer1, err1 := cmd1.Output()

		if err1 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer1)), "\n")

			if len(lines) > 1 {

				if strings.HasPrefix(lines[0], "## ") {

					tmp := strings.Split(lines[0][3:], "...")

					if len(tmp) == 2 {

						repo.CurrentBranch = tmp[0]

						if strings.Contains(tmp[1], "/") {
							repo.CurrentRemote = tmp[1][0:strings.Index(tmp[1], "/")]
						}

					}

					result = true

				}

				for l := 1; l < len(lines); l++ {

					line := lines[l]

					if strings.HasPrefix(line, "??") {
						// Untracked
					} else if strings.HasPrefix(line, "!!") {
						// Ignored
					} else if len(line) > 3 {

						// X: index, Y: worktree
						status_x := git.ToStatus(line[0:1])
						status_y := git.ToStatus(line[1:2])

						if status_x != git.StatusUnchanged && status_x != git.StatusUntracked && status_x != git.StatusIgnored {
							repo.HasLocalChanges = true
						}

						if status_y != git.StatusUnchanged && status_y != git.StatusUntracked && status_y != git.StatusIgnored {
							repo.HasRemoteChanges = true
						}

					}

				}

			} else if len(lines) == 1 {

				if strings.HasPrefix(lines[0], "## ") {

					repo.HasRemoteChanges = false
					repo.HasLocalChanges  = false

					tmp := strings.Split(lines[0][3:], "...")

					if len(tmp) == 2 {

						repo.CurrentBranch = tmp[0]

						if strings.Contains(tmp[1], "/") {
							repo.CurrentRemote = tmp[1][0:strings.Index(tmp[1], "/")]
						}

					}

					result = true

				}

			}

		}

		cmd2 := exec.Command("git", "branch", "--all")
		cmd2.Dir = repo.Folder[0:len(repo.Folder)-5]

		buffer2, err2 := cmd2.Output()

		if err2 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer2)), "\n")

			if len(lines) > 1 {

				for l := 0; l < len(lines); l++ {

					line := strings.TrimSpace(lines[l])

					if strings.HasPrefix(line, "* ") {

						branch := strings.TrimSpace(line[2:])

						if branch != "" {

							repo.CurrentBranch = branch

							if slices.Contains(repo.Branches, branch) == false {
								repo.Branches = append(repo.Branches, branch)
							}

						}

					} else if strings.HasPrefix(line, "remotes/") {

						branch := line[strings.LastIndex(line, "/")+1:]

						if branch != "" {

							if slices.Contains(repo.Branches, branch) == false {
								repo.Branches = append(repo.Branches, branch)
							}

						}

					}

				}

			} else if len(lines) == 1 {

				if strings.HasPrefix(lines[0], "* ") {

					branch := strings.TrimSpace(lines[0][2:])

					if branch != "" {

						repo.CurrentBranch = branch

						if slices.Contains(repo.Branches, branch) == false {
							repo.Branches = append(repo.Branches, branch)
						}

					}

				}

			}

		}

		cmd3 := exec.Command("git", "remote", "--verbose")
		cmd3.Dir = repo.Folder[0:len(repo.Folder)-5]

		buffer3, err3 := cmd3.Output()

		if err3 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer3)), "\n")

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if strings.Contains(line, "\t") && strings.HasSuffix(line, " (fetch)") {

					tmp := strings.Split(line[0:len(line)-8], "\t")

					if len(tmp) == 2 {

						name := strings.TrimSpace(tmp[0])
						url := strings.TrimSpace(tmp[1])

						_, ok := repo.Remotes[name]

						if ok == true {
							repo.Remotes[name].URL = url
						} else {

							remote := NewRemote(name, url)
							repo.Remotes[name] = &remote

						}

					}

				}

			}

		}

	}

	return result

}
