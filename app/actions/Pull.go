package actions

import "git-evac/server/schemas"

func Pull(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("PATCH", "/api/pull", owner, repository)
}
