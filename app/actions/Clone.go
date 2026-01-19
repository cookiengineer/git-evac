package actions

import "git-evac/schemas"

func Clone(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("POST", "/api/clone", owner, repository)
}
