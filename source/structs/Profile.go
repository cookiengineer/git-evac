package structs

type Profile struct {
	Users         map[string]*User         `json:"users"`
	Organizations map[string]*Organization `json:"organizations"`
	Settings      *Settings                `json:"settings"`
}
