package actions

import "git-evac/server/schemas"

func Fix(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("GET", "/api/terminal", owner, repository)
}
