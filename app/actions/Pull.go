package actions

import "git-evac/schemas"

func Pull(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("PATCH", "/api/pull", owner, repository)
}
