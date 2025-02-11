package structs

type User struct {
	Name       string                 `json:"name"`
	Repos      map[string]*Repository `json:"repos"`
	Remotes    map[string]*Remote     `json:"remotes"`
	PublicKeys struct {
		GPG []byte `json:"gpg"`
		SSH []byte `json:"ssh"`
	} `json:"public_keys"`
}
