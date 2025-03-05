package actions

import "git-evac/server/schemas"

func Status(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("GET", "/api/status", owner, repository)
}
