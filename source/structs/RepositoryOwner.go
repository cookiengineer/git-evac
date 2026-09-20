package structs

import "git-evac/types"
import "encoding/json"
import "sync"

type RepositoryOwner struct {
	mutex        sync.RWMutex
	Name         string                       `json:"name"`
	Folder       string                       `json:"folder"`
	Repositories map[string]*types.Repository `json:"repositories"`
}

func NewRepositoryOwner(name string, folder string) *RepositoryOwner {

	var owner RepositoryOwner

	owner.Name = name
	owner.Folder = folder
	owner.Repositories = make(map[string]*types.Repository)

	return &owner

}

func (owner *RepositoryOwner) MarshalJSON() ([]byte, error) {

	owner.mutex.RLock()
	defer owner.mutex.RUnlock()

	type Alias RepositoryOwner

	return json.Marshal((*Alias)(owner))

}

func (owner *RepositoryOwner) AddRepository(name string) bool {

	var result bool

	owner.mutex.Lock()

	_, ok := owner.Repositories[name]

	if ok == false {

		owner.Repositories[name] = types.NewRepository(name, owner.Folder + "/" + name + "/.git")
		result = true

	}

	owner.mutex.Unlock()

	return result

}

func (owner *RepositoryOwner) GetRepository(name string) *types.Repository {

	var result *types.Repository = nil

	owner.mutex.RLock()

	tmp, ok := owner.Repositories[name]

	if ok == true {
		result = tmp
	}

	owner.mutex.RUnlock()

	return result

}

func (owner *RepositoryOwner) HasRepository(name string) bool {

	var result bool

	owner.mutex.RLock()

	_, ok := owner.Repositories[name]

	if ok == true {
		result = true
	}

	owner.mutex.RUnlock()

	return result

}

func (owner *RepositoryOwner) RemoveRepository(name string) bool {

	var result bool

	owner.mutex.Lock()

	_, ok := owner.Repositories[name]

	if ok == true {
		delete(owner.Repositories, name)
		result = true
	}

	owner.mutex.Unlock()

	return result

}

func (owner *RepositoryOwner) SnapshotRepositories() map[string]*types.Repository {

	owner.mutex.RLock()
	defer owner.mutex.RUnlock()

	result := make(map[string]*types.Repository, len(owner.Repositories))

	for name, repository := range owner.Repositories {
		result[name] = repository
	}

	return result

}
