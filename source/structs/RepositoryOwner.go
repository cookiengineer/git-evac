package structs

import "git-evac/types"

type RepositoryOwner struct {
	Name         string                       `json:"name"`
	Folder       string                       `json:"folder"`
	Repositories map[string]*types.Repository `json:"repositories"`
}

func NewRepositoryOwner(name string, folder string) RepositoryOwner {

	var owner RepositoryOwner

	owner.Name = name
	owner.Folder = folder
	owner.Repositories = make(map[string]*types.Repository)

	return owner

}

func (owner *RepositoryOwner) AddRepository(name string) bool {

	_, ok := owner.Repositories[name]

	if ok == false {

		owner.Repositories[name] = types.NewRepository(name, owner.Folder + "/" + name + "/.git")

		return true

	}

	return false

}

func (owner *RepositoryOwner) GetRepository(name string) *types.Repository {

	var result *types.Repository = nil

	tmp, ok := owner.Repositories[name]

	if ok == true {
		result = tmp
	}

	return result

}

func (owner *RepositoryOwner) HasRepository(name string) bool {

	var result bool

	_, ok := owner.Repositories[name]

	if ok == true {
		result = true
	}

	return result

}

