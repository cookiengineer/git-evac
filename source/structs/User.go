package structs

type User struct {
	Name         string                 `json:"name"`
	Folder       string                 `json:"folder"`
	Repositories map[string]*Repository `json:"repositories"`
	PublicKeys   struct {
		GPG []byte `json:"gpg"`
		SSH []byte `json:"ssh"`
	} `json:"public_keys"`
}

func NewUser(name string, folder string) User {

	var user User

	user.Name = name
	user.Folder = folder
	user.Repositories = make(map[string]*Repository)

	return user

}

func (user *User) GetRepository(name string) *Repository {

	var result *Repository = nil

	tmp, ok := user.Repositories[name]

	if ok == true {
		result = tmp
	} else {

		repo := NewRepository(name, user.Folder + "/" + name + "/.git")
		user.Repositories[name] = &repo
		result = user.Repositories[name]

	}

	return result

}
