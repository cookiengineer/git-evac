package structs

type Account struct {
	Name    string                 `json:"name"`
	Repos   map[string]*Repository `json:"repos"`
	Service string                 `json:"service"`
}
