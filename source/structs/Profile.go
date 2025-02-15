package structs

import "git-evac/console"
import "io/fs"
import "os"
import os_user "os/user"

type Profile struct {
	Users         map[string]*User         `json:"users"`
	Organizations map[string]*Organization `json:"organizations"`
	Settings      Settings                 `json:"settings"`
	Filesystem    *fs.FS
}

func NewProfile(folder string, port uint16) *Profile {

	var profile Profile

	profile.Users = make(map[string]*User)
	profile.Organizations = make(map[string]*Organization)
	profile.Filesystem = nil

	profile.Settings.Folder = folder
	profile.Settings.Port = port

	user, err := os_user.Current()
	
	if err == nil {
		profile.Settings.User = user.Username
	}

	return &profile

}

func (profile *Profile) Init() {

	stat, err0 := os.Stat(profile.Settings.Folder)

	if err0 == nil && stat.IsDir() {

		console.Group("Discover Repositories in \"" + profile.Settings.Folder + "\"")

		root := profile.Settings.Folder

		entries1, err1 := os.ReadDir(root)

		if err1 == nil {

			for _, entry1 := range entries1 {

				if entry1.IsDir() {

					entries2, err2 := os.ReadDir(root + "/" + entry1.Name())

					if err2 == nil {

						for _, entry2 := range entries2 {

							if entry2.IsDir() {

								stat, err3 := os.Stat(root + "/" + entry1.Name() + "/" + entry2.Name() + "/.git")

								if err3 == nil && stat.IsDir() {

									// TODO: Better user detection, maybe with git commits?

									if entry1.Name() == profile.Settings.User {

										user := profile.GetUser(entry1.Name(), root + "/" + entry1.Name())
										repo := user.GetRepository(entry2.Name())

										if user != nil && repo != nil {
											console.Log("> Discovered @" + user.Name + "/" + repo.Name)
										}

									} else {

										orga := profile.GetOrganization(entry1.Name(), root + "/" + entry1.Name())
										repo := orga.GetRepository(entry2.Name())

										if orga != nil && repo != nil {
											console.Log("> Discovered " + orga.Name + "/" + repo.Name)
										}

									}

								}

							}

						}

					}

				}

			}

		}

		console.GroupEnd("")

	}

}

func (profile *Profile) GetOrganization(name string, folder string) *Organization {

	var result *Organization = nil

	tmp, ok := profile.Organizations[name]

	if ok == true {
		result = tmp
	} else {

		orga := NewOrganization(name, folder)
		profile.Organizations[name] = &orga
		result = profile.Organizations[name]

	}

	return result

}

func (profile *Profile) GetUser(name string, folder string) *User {

	var result *User = nil

	tmp, ok := profile.Users[name]

	if ok == true {
		result = tmp
	} else {

		user := NewUser(name, folder)
		profile.Users[name] = &user
		result = profile.Users[name]

	}

	return result

}

func (profile *Profile) HasOrganization(name string) bool {

	var result bool = false

	_, ok := profile.Organizations[name]

	if ok == true {
		result = true
	}

	return result

}

func (profile *Profile) HasUser(name string) bool {

	var result bool = false

	_, ok := profile.Users[name]

	if ok == true {
		result = true
	}

	return result

}

