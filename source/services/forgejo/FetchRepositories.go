package forgejo

import "git-evac/types"
import "fmt"
import "net/http"
import "strings"

func FetchRepositories(api_url string, owner string, token string, parent_folder string) []*types.Repository {

	result := make([]*types.Repository, 0)

	orgas_url := fmt.Sprintf("%s/api/v1/orgs/%s/repos?limit=0&sort=updated&order=desc", api_url, owner)
	users_url := fmt.Sprintf("%s/api/v1/users/%s/repos?limit=0&sort=updated&order=desc", api_url, owner)

	orga_repositories, orga_status, orga_err := fetchAPI(orgas_url, token)

	if orga_err == nil && orga_status == http.StatusOK {

		for _, remote_repository := range orga_repositories {

			if remote_repository.Name != "" && strings.HasPrefix(remote_repository.FullName, owner + "/") {

				repository := types.NewRepository(remote_repository.Name, parent_folder + "/" + remote_repository.Name + "/.git")
				result = append(result, repository)

			}

		}

	} else {

		user_repositories, user_status, user_err := fetchAPI(users_url, token)

		if user_err == nil && user_status == http.StatusOK {

			for _, remote_repository := range user_repositories {

				if remote_repository.Name != "" && strings.HasPrefix(remote_repository.FullName, owner + "/") {

					repository := types.NewRepository(remote_repository.Name, parent_folder + "/" + remote_repository.Name + "/.git")
					result = append(result, repository)

				}

			}

		}

	}

	return result

}

