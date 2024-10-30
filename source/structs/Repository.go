package structs

type Repository struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Service  string `json:"service"`
	IsPublic bool   `json:"is_public"`
}
