package structs

import "fmt"
import "os"
import "os/exec"
import "strings"

type Repository struct {
	Name          string             `json:"name"`
	Folder        string             `json:"folder"` // /path/to/.git
	Branches      map[string]string  `json:"branches"`
	Remotes       map[string]*Remote `json:"remotes"`
	IsPublic      bool               `json:"is_public"`
	CurrentBranch string             `json:"current_branch"`
	CurrentRemote string             `json:"current_remote"`
}

func NewRepository(name string, folder string) Repository {

	var repo Repository

	repo.Name = name
	repo.Folder = folder
	repo.Branches = make(map[string]string)
	repo.Remotes = make(map[string]*Remote)

	repo.CurrentBranch = "master"
	repo.CurrentRemote = "origin"

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
			fmt.Println("status!", repo.Folder)
			fmt.Println(string(buffer1))
		}

	}

	return result

}
