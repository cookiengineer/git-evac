package structs

type Database struct {
	// TODO: Database Struct
}

func NewDatabase() Database {

	var database Database

	return database

}

func (database *Database) Initialize(folder string) *Profile {

	var profile Profile

	settings := Settings{
		Folder: folder,
	}

	// TODO: Walk folders
	// TODO: Walk organizations
	// TODO: Walk users

	profile.Settings = &settings

	return &profile

}
