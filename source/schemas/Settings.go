package schemas

import "git-evac/structs"

type Settings struct {
	Settings structs.Settings `json:"settings"`
}

func (schema *Settings) IsValid() bool {
	return schema.Settings.IsValid()
}
