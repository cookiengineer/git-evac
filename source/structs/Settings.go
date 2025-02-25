package structs

type Settings struct {
	User          string                    `json:"user"`
	Folder        string                    `json:"folder"`
	Port          uint16                    `json:"port"`
	Users         map[string]RemoteSettings `json:"users"`
	Organizations map[string]RemoteSettings `json:"organizations"`
}

type RemoteSettings struct {
	Name    string             `json:"name"`
	Remotes map[string]*Remote `json:"remotes"`
}
