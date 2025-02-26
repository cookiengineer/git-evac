package schemas

import "git-evac/structs"

type Index struct {
	Owners map[string]*structs.RepositoryOwner `json:"owners"`
}
