package structs

type Organization struct {
	Name         string                 `json:"name"`
	Repositories map[string]*Repository `json:"repositories"`
	Remotes      map[string]*Remote     `json:"remotes"`
}
