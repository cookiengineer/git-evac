package types

import "encoding/json"
import "sync"

type Repository struct {
	mutex            sync.RWMutex
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

func (repo *Repository) MarshalJSON() ([]byte, error) {

	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	type Alias Repository

	return json.Marshal((*Alias)(repo))

}

func (repo *Repository) GetName() string {

	var result string

	repo.mutex.RLock()
	result = repo.Name
	repo.mutex.RUnlock()

	return result

}

func (repo *Repository) GetFolder() string {

	var result string

	repo.mutex.RLock()
	result = repo.Folder
	repo.mutex.RUnlock()

	return result

}

func (repo *Repository) GetCurrentBranch() string {

	var result string

	repo.mutex.RLock()
	result = repo.CurrentBranch
	repo.mutex.RUnlock()

	return result

}

func (repo *Repository) GetCurrentRemote() string {

	var result string

	repo.mutex.RLock()
	result = repo.CurrentRemote
	repo.mutex.RUnlock()

	return result

}

func (repo *Repository) GetRemote(name string) *Remote {

	var result *Remote = nil

	repo.mutex.RLock()

	remote, ok := repo.Remotes[name]

	if ok == true {
		result = remote
	}

	repo.mutex.RUnlock()

	return result

}

func (repo *Repository) SnapshotRemotes() map[string]*Remote {

	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	result := make(map[string]*Remote, len(repo.Remotes))

	for name, remote := range repo.Remotes {
		result[name] = remote
	}

	return result

}

func (repo *Repository) NeedsClone() bool {

	var result bool

	repo.mutex.RLock()

	if len(repo.Branches) == 0 {
		result = true
	}

	repo.mutex.RUnlock()

	return result

}

func (repo *Repository) NeedsCommit() bool {

	var result bool

	repo.mutex.RLock()

	if repo.HasLocalChanges == true {
		result = true
	}

	repo.mutex.RUnlock()

	return result

}

func (repo *Repository) NeedsFix() bool {

	var result bool

	repo.mutex.RLock()

	// TODO: Check against schema, maybe with a NeedsRemoteFix(schema)?

	if repo.HasRemoteChanges == true {
		result = true
	}

	repo.mutex.RUnlock()

	return result

}
