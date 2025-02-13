package structs

type Settings struct {
	User    string   `json:"user"`
	Folder  string   `json:"folder"`
	Port    uint16   `json:"port"`

	// TODO: Configure Gogs instances and their Remote schemas via UI
	// Servers []Server `json:"servers"`
}

