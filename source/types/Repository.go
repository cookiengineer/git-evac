package types

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

func NewRepository(name string, folder string) *Repository {

	var repo Repository

	repo.Name = name
	repo.Folder = folder
	repo.Branches = make([]string, 0)
	repo.Remotes = make(map[string]*Remote)

	repo.CurrentBranch = "master"
	repo.CurrentRemote = "origin"
	repo.Identity = ""

	repo.Status()

	return &repo

}

func (repo *Repository) NeedsClone() bool {

	var result bool

	if len(repo.Branches) == 0 {
		result = true
	}

	return result

}

func (repo *Repository) NeedsCommit() bool {

	var result bool

	if repo.HasLocalChanges == true {
		result = true
	}

	return result

}

func (repo *Repository) NeedsFix() bool {

	var result bool

	// TODO: Check against schema, maybe with a NeedsRemoteFix(schema)?

	if repo.HasRemoteChanges == true {
		result = true
	}

	return result

}

