package structs

type Profile struct {
	Users         map[string]*User         `json:"users"`
	Organizations map[string]*Organization `json:"organizations"`
	Remotes       map[string][]*Remote     `json:"remotes"`
	Settings      *Settings                `json:"settings"`
}
