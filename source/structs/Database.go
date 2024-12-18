package structs

type Database struct {
	// TODO: Database Struct
}

func NewDatabase() Database {

	var database Database

	return database

}

func (database *Database) GetConfig(orga string, repo string) *Config {

	var config Config

	config.Orga = orga
	config.Repo = repo

	return &config

}
