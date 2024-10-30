package structs

type Profile struct {
	Users map[string]*Account `json:"users"`
	Orgas map[string]*Account `json:"orgas"`
}
