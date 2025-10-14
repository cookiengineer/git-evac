package actions

import "git-evac/schemas"

func Clone(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("GET", "/api/clone", owner, repository)
}
