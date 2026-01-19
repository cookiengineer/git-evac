package types

import utils_strings "git-evac/utils/strings"

type Service struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Type  string `json:"type"`
	Token string `json:"token"`
}

func NewService(name string) *Service {

	var service Service

	service.Name = name
	service.URL = ""
	service.Token = ""
	service.Type = "git"

	return &service

}

func (service *Service) IsValid() bool {

	if utils_strings.IsName(service.Name) {

		// TODO: Valid URL Schemas

		return true

	}

	return false

}
