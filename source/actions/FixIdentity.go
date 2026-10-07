package actions

import "git-evac/structs"
import "errors"

func FixIdentity(profile *structs.Profile, owner_name string, repo_name string) error {

	repository := profile.GetRepository(owner_name, repo_name)

	if repository == nil {
		return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" does not exist")
	}

	settings_owner := profile.Settings.GetOwner(owner_name)

	if settings_owner == nil {
		return errors.New("Owner \"" + owner_name + "\" has no settings")
	}

	identity := settings_owner.EffectiveIdentity(repo_name)

	if identity == nil {
		return nil
	}

	return repository.ApplyIdentity(identity)

}
