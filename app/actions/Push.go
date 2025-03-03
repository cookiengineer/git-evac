package actions

import "git-evac/server/schemas"

func Push(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("/api/push", owner, repository)
}
