package actions

import "git-evac/schemas"

func FixRemotes(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("PATCH", "/api/fixremotes", owner, repository)
}
