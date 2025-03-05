package actions

import "git-evac/server/schemas"

func Clone(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("GET", "/api/clone", owner, repository)
}
