//go:build wasm

package types

import "strings"

func (repo *Repository) AddRemote(owner_name string, repo_name string, schema Remote) bool {

	remote_name := schema.Name
	remote_url  := schema.URL
	remote_url   = strings.ReplaceAll(remote_url, "{{owner}}", owner_name)
	remote_url   = strings.ReplaceAll(remote_url, "{{repo}}", repo_name)

	repo.Remotes[remote_name] = NewRemote(remote_name, remote_url)

	return true

}

func (repo *Repository) RemoveRemote(remote_name string) bool {

	_, ok := repo.Remotes[remote_name]

	if ok == true {

		delete(repo.Remotes, remote_name)

		return true

	}

	return false

}

func (repo *Repository) Init() bool {
	return false
}

func (repo *Repository) Status() bool {
	return false
}
