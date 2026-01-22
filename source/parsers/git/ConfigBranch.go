package git

type ConfigBranch struct {
	Name   string `json:"name"`
	Remote string `json:"remote"`
	Merge  string `json:"merge"`
}

