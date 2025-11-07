package structs

import "os"

func (profile *Profile) RefreshRepositories() {

	stat, err := os.Stat(profile.Settings.Folder)

	if err == nil && stat.IsDir() {

		profile.Console.Group("Refresh Repositories")

		info_owners, err_owners := os.ReadDir(profile.Settings.Folder)

		if err_owners == nil {

			for _, info_owner := range info_owners {

				if info_owner.IsDir() == true {

					info_repositories, err_repositories := os.ReadDir(profile.Settings.Folder + "/" + info_owner.Name())

					if err_repositories == nil {

						for _, info_repository := range info_repositories {

							if info_repository.IsDir() == true {

								stat, err := os.Stat(profile.Settings.Folder + "/" + info_owner.Name() + "/" + info_repository.Name() + "/.git")

								if err == nil && stat.IsDir() == true {

									owner_name := info_owner.Name()
									repository_name := info_repository.Name()

									if profile.HasRepositoryOwner(owner_name) == false {
										profile.AddRepositoryOwner(owner_name, profile.Settings.Folder + "/" + owner_name)
									}

									if profile.HasRepository(owner_name, repository_name) == false {
										owner := profile.GetRepositoryOwner(owner_name)
										owner.AddRepository(repository_name)
										profile.Console.Log("> " + owner_name + "/" + repository_name)
									}

								}

							}

						}

					}

				}

			}

		}

		profile.Console.GroupEnd("Refresh Repositories")

	} else {
		profile.Console.Warn("No Repositories in Folder \"" + profile.Settings.Folder + "\"")
	}

	for _, owner := range profile.Repositories {

		for _, repo := range owner.Repositories {
			repo.Status()
		}

	}

}

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
