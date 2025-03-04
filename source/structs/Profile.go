package structs

import "git-evac/console"
import "io/fs"
import "os"

type Profile struct {
	Owners     map[string]*RepositoryOwner `json:"owners"`
	Settings   Settings                    `json:"settings"`
	Filesystem *fs.FS
}

func NewProfile(backup string, folder string, port uint16) *Profile {

	var profile Profile

	profile.Owners = make(map[string]*RepositoryOwner)
	profile.Filesystem = nil

	profile.Settings.Backup = backup
	profile.Settings.Folder = folder
	profile.Settings.Port = port

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

									// TODO: Read and parse .git/config
									// TODO: Use the user.name or user.email setting to detect identity

									owner := profile.GetOwner(entry1.Name(), root + "/" + entry1.Name())
									repo := owner.GetRepository(entry2.Name())

									if owner != nil && repo != nil {
										console.Log("> Discovered " + owner.Name + "/" + repo.Name)
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

func (profile *Profile) GetOwner(name string, folder string) *RepositoryOwner {

	var result *RepositoryOwner = nil

	tmp, ok := profile.Owners[name]

	if ok == true {
		result = tmp
	} else {

		owner := NewRepositoryOwner(name, folder)
		profile.Owners[name] = &owner
		result = profile.Owners[name]

	}

	return result

}

func (profile *Profile) HasOwner(name string) bool {

	var result bool = false

	_, ok := profile.Owners[name]

	if ok == true {
		result = true
	}

	return result

}

