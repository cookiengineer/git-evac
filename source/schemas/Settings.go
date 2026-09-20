package schemas

import "git-evac/structs"

type Settings struct {
	Settings *structs.Settings `json:"settings"`
}

func (schema *Settings) IsValid() bool {

	if schema.Settings == nil {
		return false
	}

	return schema.Settings.IsValid()

}
