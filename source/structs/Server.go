package structs

import "git-evac/api"

type Server struct {
	Type    api.Type `json:"type"`
	Token   string   `json:"token"`
	URL     string   `json:"url"`
	Remotes struct {
		GIT   string `json:"git"`
		HTTP  string `json:"http"`
		HTTPS string `json:"https"`
		SSH   string `json:"ssh"`
	} `json:"remotes"`
}
