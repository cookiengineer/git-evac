package actions

import "git-evac/schemas"

func FixIdentity(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("PATCH", "/api/fixidentity", owner, repository)
}
