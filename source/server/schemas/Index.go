package schemas

import "git-evac/structs"

type Index struct {
	Users         map[string]*structs.User         `json:"users"`
	Organizations map[string]*structs.Organization `json:"organizations"`
}
