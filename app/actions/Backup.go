package actions

import "git-evac/schemas"

func Backup(owner string, repository string) (*schemas.Repository, error) {
	return fetchAPI("PATCH", "/api/backup", owner, repository)
}
