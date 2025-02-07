package structs

type User struct {
	Name       string                 `json:"name"`
	Repos      map[string]*Repository `json:"repos"`
	Service    string                 `json:"service"`
	PublicKeys struct {
		GPG []byte `json:"gpg"`
		SSH []byte `json:"ssh"`
	} `json:"public_keys"`
}
