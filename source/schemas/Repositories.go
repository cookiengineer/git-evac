package schemas

import "git-evac/structs"

type Repositories struct {
	Owners map[string]*structs.RepositoryOwner `json:"owners"`
}
