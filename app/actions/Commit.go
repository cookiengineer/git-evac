package actions

import "git-evac/server/schemas"

func Commit(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("/api/commit", owner, repository)
}
