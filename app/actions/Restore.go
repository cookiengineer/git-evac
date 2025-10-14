package actions

import "git-evac/schemas"

func Restore(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("PATCH", "/api/restore", owner, repository)
}
