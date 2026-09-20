package remotes

import "git-evac/structs"
import "strings"

func GuessOrigin(owner *structs.RepositoryOwner, repository string) string {

	var result string

	candidates := make(map[string]int)

	for _, repo := range owner.SnapshotRepositories() {

		remote := repo.GetRemote("origin")

		if remote != nil {

			url := remote.GetURL()

			if strings.HasPrefix(url, "git@bitbucket.org") && strings.HasSuffix(url, "/" + repo.GetName() + ".git") {

				candidate := url[0:strings.Index(url, "/" + repo.GetName() + ".git")] + "/" + repository + ".git"

				rank, ok := candidates[candidate]

				if ok == true {
					candidates[candidate] = rank + 1
				} else {
					candidates[candidate] = 1
				}

			} else if strings.HasPrefix(url, "https://gitlab.com") && strings.HasSuffix(url, "/" + repo.GetName() + ".git") {

				candidate := url[0:strings.Index(url, "/" + repo.GetName() + ".git")] + "/" + repository + ".git"

				rank, ok := candidates[candidate]

				if ok == true {
					candidates[candidate] = rank + 1
				} else {
					candidates[candidate] = 1
				}

			} else if strings.HasPrefix(url, "git@gitlab.com") && strings.HasSuffix(url, "/" + repo.GetName() + ".git") {

				candidate := url[0:strings.Index(url, "/" + repo.GetName() + ".git")] + "/" + repository + ".git"

				rank, ok := candidates[candidate]

				if ok == true {
					candidates[candidate] = rank + 1
				} else {
					candidates[candidate] = 1
				}

			} else if strings.HasPrefix(url, "https://github.com") && strings.HasSuffix(url, "/" + repo.GetName() + ".git") {

				candidate := url[0:strings.Index(url, "/" + repo.GetName() + ".git")] + "/" + repository + ".git"

				rank, ok := candidates[candidate]

				if ok == true {
					candidates[candidate] = rank + 1
				} else {
					candidates[candidate] = 1
				}

			} else if strings.HasPrefix(url, "git@github.com") && strings.HasSuffix(url, "/" + repo.GetName() + ".git") {

				candidate := url[0:strings.Index(url, "/" + repo.GetName() + ".git")] + "/" + repository + ".git"

				rank, ok := candidates[candidate]

				if ok == true {
					candidates[candidate] = rank + 1
				} else {
					candidates[candidate] = 1
				}

			} else if strings.HasSuffix(url, "/" + repo.GetName() + ".git") {

				candidate := url[0:strings.Index(url, "/" + repo.GetName() + ".git")] + "/" + repository + ".git"

				rank, ok := candidates[candidate]

				if ok == true {
					candidates[candidate] = rank + 1
				} else {
					candidates[candidate] = 1
				}

			}

		}

	}

	current_rank := 0
	current_candidate := ""

	for candidate, rank := range candidates {

		if rank > current_rank {
			current_candidate = candidate
			current_rank = rank
		}

	}

	if current_candidate != "" {
		result = current_candidate
	}

	return result

}
