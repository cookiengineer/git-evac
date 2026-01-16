package types

import utils_strings "git-evac/utils/strings"

type Service struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Type  string `json:"type"`
	Token string `json:"token"`
}

func NewService(name string) Service {

	var settings Service

	settings.Name = name
	settings.URL = ""
	settings.Token = ""
	settings.Type = "git"

	return settings

}

func (settings *Service) IsValid() bool {

	if utils_strings.IsName(settings.Name) {

		// TODO: Valid URL Schemas

		return true

	}

	return false

}
