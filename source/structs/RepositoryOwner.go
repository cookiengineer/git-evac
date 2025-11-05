package structs

type RepositoryOwner struct {
	Name         string                 `json:"name"`
	Folder       string                 `json:"folder"`
	Repositories map[string]*Repository `json:"repositories"`
}

func NewRepositoryOwner(name string, folder string) RepositoryOwner {

	var owner RepositoryOwner

	owner.Name = name
	owner.Folder = folder
	owner.Repositories = make(map[string]*Repository)

	return owner

}

func (owner *RepositoryOwner) AddRepository(name string) bool {

	_, ok := owner.Repositories[name]

	if ok == false {

		repo := NewRepository(name, owner.Folder + "/" + name + "/.git")
		owner.Repositories[name] = &repo

		return true

	}

	return false

}

func (owner *RepositoryOwner) GetRepository(name string) *Repository {

	var result *Repository

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

