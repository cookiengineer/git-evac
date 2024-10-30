package git

type ConfigRemote struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Fetch string `json:"fetch"`
}

