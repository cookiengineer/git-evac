package git

type ConfigDifftool struct {
	Name   string `json:"name"`
	Cmd    string `json:"cmd"`
	Prompt bool   `json:"prompt"`
}

