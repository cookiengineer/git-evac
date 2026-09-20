package structs

import "git-evac/types"
import services_forgejo "git-evac/services/forgejo"
import services_github "git-evac/services/github"
import services_gitlab "git-evac/services/gitlab"
import services_gitea "git-evac/services/gitea"
import services_gogs "git-evac/services/gogs"
import "os"

func (profile *Profile) RefreshLocalRepositories() {

	folder := profile.Settings.GetFolder()

	stat, err := os.Stat(folder)

	if err == nil && stat.IsDir() {

		profile.Console.Group("Refresh Local Repositories")

		info_owners, err_owners := os.ReadDir(folder)

		if err_owners == nil {

			owners_on_disk       := make(map[string]bool)
			repositories_on_disk := make(map[string]bool)

			for _, info_owner := range info_owners {

				if info_owner.IsDir() == true {

					owner_name := info_owner.Name()

					owners_on_disk[owner_name] = true

					if profile.HasRepositoryOwner(owner_name) == false {
						profile.AddRepositoryOwner(owner_name, folder + "/" + owner_name)
					}

					info_repositories, err_repositories := os.ReadDir(folder + "/" + owner_name)

					if err_repositories == nil {

						for _, info_repository := range info_repositories {

							if info_repository.IsDir() == true {

								stat, err := os.Stat(folder + "/" + owner_name + "/" + info_repository.Name() + "/.git")

								if err == nil && stat.IsDir() == true {

									repository_name := info_repository.Name()

									repositories_on_disk[owner_name + "/" + repository_name] = true

									if profile.HasRepository(owner_name, repository_name) == false {

										profile.Console.Log("> Add " + owner_name + "/" + repository_name)
										owner := profile.GetRepositoryOwner(owner_name)

										if owner != nil {
											owner.AddRepository(repository_name)
										}

									}

								}

							}

						}

					}

				}

			}

			for owner_name, owner := range profile.SnapshotRepositories() {

				if owners_on_disk[owner_name] == false {

					if profile.RemoveRepositoryOwner(owner_name) == true {
						profile.Console.Log("> Remove " + owner_name)
					}

				} else {

					for repository_name := range owner.SnapshotRepositories() {

						if repositories_on_disk[owner_name + "/" + repository_name] == false {

							if owner.RemoveRepository(repository_name) == true {
								profile.Console.Log("> Remove " + owner_name + "/" + repository_name)
							}

						}

					}

				}

			}

		}

		profile.Console.GroupEnd("Refresh Local Repositories")

	} else {
		profile.Console.Warn("No Repositories in Folder \"" + folder + "\"")
	}

}

func (profile *Profile) RefreshServiceRepositories() {

	folder := profile.Settings.GetFolder()

	stat, err := os.Stat(folder)

	if err == nil && stat.IsDir() {

		profile.Console.Group("Refresh Service Repositories")

		info_owners, err_owners := os.ReadDir(folder)

		if err_owners == nil {

			for _, info_owner := range info_owners {

				if info_owner.IsDir() == true {

					owner_name := info_owner.Name()

					settings_owner := profile.Settings.GetOwner(owner_name)

					if settings_owner != nil {

						for remote_name, service := range settings_owner.SnapshotServices() {

							remote_repositories := make([]*types.Repository, 0)

							switch service.GetType() {
							case "forgejo":
								remote_repositories = services_forgejo.FetchRepositories(service.GetURL(), owner_name, service.GetToken(), folder + "/" + owner_name)
							case "github":
								remote_repositories = services_github.FetchRepositories(service.GetURL(), owner_name, service.GetToken(), folder + "/" + owner_name)
							case "gitlab":
								remote_repositories = services_gitlab.FetchRepositories(service.GetURL(), owner_name, service.GetToken(), folder + "/" + owner_name)
							case "gitea":
								remote_repositories = services_gitea.FetchRepositories(service.GetURL(), owner_name, service.GetToken(), folder + "/" + owner_name)
							case "gogs":
								remote_repositories = services_gogs.FetchRepositories(service.GetURL(), owner_name, service.GetToken(), folder + "/" + owner_name)
							}

							if len(remote_repositories) > 0 {

								for _, repository := range remote_repositories {

									repository_name := repository.GetName()

									if profile.HasRepository(owner_name, repository_name) == false {

										profile.Console.Log("> Init " + owner_name + "/" + repository_name)

										owner := profile.GetRepositoryOwner(owner_name)

										if owner != nil {

											owner.AddRepository(repository_name)

											remote := settings_owner.GetRemote(remote_name)
											repo   := owner.GetRepository(repository_name)

											if repo != nil && remote != nil {

												// Use remote as schema
												repo.AddRemote(owner_name, repository_name, types.NewRemote(remote.GetName(), remote.GetURL()))

											}

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
		profile.Console.Warn("No Repositories in Folder \"" + folder + "\"")
	}

}

func (profile *Profile) AddRepositoryOwner(owner_name string, owner_folder string) bool {

	var result bool

	profile.mutex.Lock()

	_, ok := profile.Repositories[owner_name]

	if ok == false {

		profile.Repositories[owner_name] = NewRepositoryOwner(owner_name, owner_folder)
		result = true

	}

	profile.mutex.Unlock()

	return result

}

func (profile *Profile) RemoveRepositoryOwner(owner_name string) bool {

	var result bool

	if owner_name != "" {

		profile.mutex.Lock()

		_, ok := profile.Repositories[owner_name]

		if ok == true {
			delete(profile.Repositories, owner_name)
			result = true
		}

		profile.mutex.Unlock()

	}

	return result

}

func (profile *Profile) GetRepositoryOwner(owner_name string) *RepositoryOwner {

	var result *RepositoryOwner = nil

	profile.mutex.RLock()

	owner, ok := profile.Repositories[owner_name]

	if ok == true {
		result = owner
	}

	profile.mutex.RUnlock()

	return result

}

func (profile *Profile) GetRepository(owner_name string, repo_name string) *types.Repository {

	var result *types.Repository = nil

	owner := profile.GetRepositoryOwner(owner_name)

	if owner != nil {
		result = owner.GetRepository(repo_name)
	}

	return result

}

func (profile *Profile) RemoveRepository(owner_name string, repo_name string) bool {

	var result bool

	owner := profile.GetRepositoryOwner(owner_name)

	if owner != nil {
		result = owner.RemoveRepository(repo_name)
	}

	return result

}

func (profile *Profile) HasRepositoryOwner(owner_name string) bool {

	var result bool

	profile.mutex.RLock()

	_, ok := profile.Repositories[owner_name]

	if ok == true {
		result = true
	}

	profile.mutex.RUnlock()

	return result

}

func (profile *Profile) HasRepository(owner_name string, repo_name string) bool {

	var result bool

	owner := profile.GetRepositoryOwner(owner_name)

	if owner != nil {
		result = owner.HasRepository(repo_name)
	}

	return result

}
