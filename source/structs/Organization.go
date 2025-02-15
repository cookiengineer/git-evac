package structs

type Organization struct {
	Name         string                 `json:"name"`
	Folder       string                 `json:"folder"`
	Repositories map[string]*Repository `json:"repositories"`
}

func NewOrganization(name string, folder string) Organization {

	var orga Organization

	orga.Name = name
	orga.Folder = folder
	orga.Repositories = make(map[string]*Repository)

	return orga

}

func (orga *Organization) GetRepository(name string) *Repository {

	var result *Repository = nil

	tmp, ok := orga.Repositories[name]

	if ok == true {
		result = tmp
	} else {

		repo := NewRepository(name, orga.Folder + "/" + name + "/.git")
		orga.Repositories[name] = &repo
		result = orga.Repositories[name]

	}

	return result

}

func (orga *Organization) HasRepository(name string) bool {

	var result bool = false

	_, ok := orga.Repositories[name]

	if ok == true {
		result = true
	}

	return result

}

