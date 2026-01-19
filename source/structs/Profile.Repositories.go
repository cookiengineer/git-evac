package structs

import "git-evac/types"
import services_github "git-evac/services/github"
// import services_gitlab "git-evac/services/gitlab"
import services_gogs "git-evac/services/gogs"
import "os"

func (profile *Profile) RefreshLocalRepositories() {

	stat, err := os.Stat(profile.Settings.Folder)

	if err == nil && stat.IsDir() {

		profile.Console.Group("Refresh Local Repositories")

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

										profile.Console.Log("> Add " + owner_name + "/" + repository_name)
										owner := profile.GetRepositoryOwner(owner_name)
										owner.AddRepository(repository_name)

									}

								}

							}

						}

					}

				}

			}

		}

		profile.Console.GroupEnd("Refresh Local Repositories")

	} else {
		profile.Console.Warn("No Repositories in Folder \"" + profile.Settings.Folder + "\"")
	}

}

func (profile *Profile) RefreshServiceRepositories() {

	stat, err := os.Stat(profile.Settings.Folder)

	if err == nil && stat.IsDir() {

		profile.Console.Group("Refresh Service Repositories")

		info_owners, err_owners := os.ReadDir(profile.Settings.Folder)

		if err_owners == nil {

			for _, info_owner := range info_owners {

				if info_owner.IsDir() == true {

					owner_name := info_owner.Name()

					_, ok1 := profile.Settings.Owners[owner_name]

					if ok1 == true {

						for remote_name, service := range profile.Settings.Owners[owner_name].Services {

							switch service.Type {
							case "github":

								remote_repositories := services_github.FetchRepositories(service.URL, owner_name, service.Token, profile.Settings.Folder + "/" + owner_name)

								for _, repository := range remote_repositories {

									repository_name := repository.Name

									if profile.HasRepository(owner_name, repository_name) == false {

										profile.Console.Log("> Init " + owner_name + "/" + repository_name)

										owner := profile.GetRepositoryOwner(owner_name)
										owner.AddRepository(repository_name)

										remote, ok2 := profile.Settings.Owners[owner_name].Remotes[remote_name]
										repo := owner.GetRepository(repository_name)

										if repo != nil && ok2 == true {

											// Use remote as schema
											repo.AddRemote(owner_name, repository_name, types.Remote{
												Name: remote.Name,
												URL:  remote.URL,
											})

										}

									}

								}

							case "gitlab":

								// TODO: Support gitlab API

							case "gitea":

								// TODO: Support gitea API

							case "gogs":

								remote_repositories := services_gogs.FetchRepositories(service.URL, owner_name, service.Token, profile.Settings.Folder + "/" + owner_name)

								for _, repository := range remote_repositories {

									repository_name := repository.Name

									if profile.HasRepository(owner_name, repository_name) == false {

										profile.Console.Log("> Init " + owner_name + "/" + repository_name)

										owner := profile.GetRepositoryOwner(owner_name)
										owner.AddRepository(repository_name)

										remote, ok2 := profile.Settings.Owners[owner_name].Remotes[remote_name]
										repo := owner.GetRepository(repository_name)

										if repo != nil && ok2 == true {

											// Use remote as schema
											repo.AddRemote(owner_name, repository_name, types.Remote{
												Name: remote.Name,
												URL:  remote.URL,
											})

										}

									}

								}

							}

						}

					}

				}

			}

		}

		profile.Console.GroupEnd("Refresh Service Repositories")

	} else {
		profile.Console.Warn("No Repositories in Folder \"" + profile.Settings.Folder + "\"")
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

	var result *RepositoryOwner = nil

	owner, ok := profile.Repositories[owner_name]

	if ok == true {
		result = owner
	}

	return result

}

func (profile *Profile) GetRepository(owner_name string, repo_name string) *types.Repository {

	var result *types.Repository = nil

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
