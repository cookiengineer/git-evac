package structs

func (profile *Profile) AddRepositoryOwner(owner_name string, owner_folder string) bool {

	_, ok := profile.Repositories[owner_name]

	if ok == false {

		owner := NewRepositoryOwner(owner_name, owner_folder)
		profile.Repositories[owner_name] = &owner

		return true

	}

	return false

}

func (profile *Profile) GetRepositoryOwner(owner_name string) *RepositoryOwner {

	var result *RepositoryOwner

	owner, ok := profile.Repositories[owner_name]

	if ok == true {
		result = owner
	}

	return result

}

func (profile *Profile) GetRepository(owner_name string, repo_name string) *Repository {

	var result *Repository

	owner, ok1 := profile.Repositories[owner_name]

	if ok1 == true {

		repository, ok2 := owner.Repositories[repo_name]

		if ok2 == true {
			result = repository
		}

	}

	return result

}

func (profile *Profile) HasRepositoryOwner(owner_name string) bool {

	var result bool

	_, ok := profile.Repositories[owner_name]

	if ok == true {
		result = true
	}

	return result

}

func (profile *Profile) HasRepository(owner_name string, repo_name string) bool {

	var result bool

	owner, ok1 := profile.Repositories[owner_name]

	if ok1 == true {

		_, ok2 := owner.Repositories[repo_name]

		if ok2 == true {
			result = true
		}

	}

	return result

}
