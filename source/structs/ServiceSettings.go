package structs

import utils_strings "git-evac/utils/strings"

type ServiceSettings struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Type  string `json:"type"`
	Token string `json:"token"`
}

func NewServiceSettings(name string) ServiceSettings {

	var settings ServiceSettings

	settings.Name = name
	settings.URL = ""
	settings.Token = ""
	settings.Type = "git"

	return settings

}

func (settings *ServiceSettings) IsValid() bool {

	if utils_strings.IsName(settings.Name) {

		// TODO: Valid URL Schemas

		return true

	}

	return false

}
