package types

import utils_strings "git-evac/utils/strings"
import "encoding/json"
import "strings"
import "sync"

type Service struct {
	mutex sync.RWMutex
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

func (service *Service) MarshalJSON() ([]byte, error) {

	service.mutex.RLock()
	defer service.mutex.RUnlock()

	type Alias Service

	return json.Marshal((*Alias)(service))

}

func (service *Service) GetName() string {

	var result string

	service.mutex.RLock()
	result = service.Name
	service.mutex.RUnlock()

	return result

}

func (service *Service) GetURL() string {

	var result string

	service.mutex.RLock()
	result = service.URL
	service.mutex.RUnlock()

	return result

}

func (service *Service) GetType() string {

	var result string

	service.mutex.RLock()
	result = service.Type
	service.mutex.RUnlock()

	return result

}

func (service *Service) GetToken() string {

	var result string

	service.mutex.RLock()
	result = service.Token
	service.mutex.RUnlock()

	return result

}

func (service *Service) IsValid() bool {

	service.mutex.RLock()
	defer service.mutex.RUnlock()

	if utils_strings.IsName(service.Name) {

		// TODO: Valid URL Schemas

		return true

	}

	return false

}

func (service *Service) SetURL(value string) bool {

	var result bool

	service.mutex.Lock()
	defer service.mutex.Unlock()

	if strings.HasPrefix(value, "https://codeberg.org") {

		service.Type = "forgejo"
		service.URL  = "https://codeberg.org"
		result = true

	} else if strings.HasPrefix(value, "https://api.github.com") || strings.HasPrefix(value, "https://github.com") {

		service.Type = "github"
		service.URL  = "https://api.github.com"
		result = true

	} else if strings.HasPrefix(value, "https://gitlab.com") {

		service.Type = "gitlab"
		service.URL  = "https://gitlab.com"
		result = true

	} else {

		if strings.HasPrefix(value, "https://") || strings.HasPrefix(value, "http://") {

			if strings.HasSuffix(value, "/") {
				service.URL = strings.TrimSpace(value[0:len(value) - 1])
				result = true
			} else {
				service.URL = strings.TrimSpace(value)
				result = true
			}

		}

	}

	return result

}

func (service *Service) SetToken(value string) bool {

	var result bool

	service.mutex.Lock()
	defer service.mutex.Unlock()

	if strings.TrimSpace(value) != "" {
		service.Token = strings.TrimSpace(value)
	}

	return result

}

func (service *Service) SetType(value string) bool {

	var result bool

	service.mutex.Lock()
	defer service.mutex.Unlock()

	switch value {
	case "forgejo":
		service.Type = value
		result = true
	case "github":
		service.Type = value
		result = true
	case "gitlab":
		service.Type = value
		result = true
	case "gitea":
		service.Type = value
		result = true
	case "gogs":
		service.Type = value
		result = true
	}

	return result

}
