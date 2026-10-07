package actions

import "git-evac/structs"
import "errors"

func FixRemotes(profile *structs.Profile, owner_name string, repo_name string) error {

	repository := profile.GetRepository(owner_name, repo_name)

	if repository == nil {
		return errors.New("Repository \"" + owner_name + "/" + repo_name + "\" does not exist")
	}

	settings_owner := profile.Settings.GetOwner(owner_name)

	if settings_owner == nil {
		return errors.New("Owner \"" + owner_name + "\" has no settings")
	}

	effective := settings_owner.Effective(repo_name)

	if effective == nil {
		return errors.New("Effective settings for \"" + owner_name + "/" + repo_name + "\" are invalid")
	}

	cloned := repository.IsCloned()

	for name, remote := range effective.SnapshotRemotes() {

		if remote == nil {
			continue
		}

		ok := repository.AddRemote(owner_name, repo_name, remote)

		if cloned == true && ok == false {
			return errors.New("Failed to add remote \"" + name + "\" to \"" + owner_name + "/" + repo_name + "\"")
		}

	}

	// Remotes that exist on the owner level but are missing from the effective
	// settings were explicitly removed by a per-repository override. Remove
	// them so that e.g. a private repository does not keep a public remote.
	owner_remotes := settings_owner.SnapshotRemotes()
	effective_remotes := effective.SnapshotRemotes()

	for name := range owner_remotes {

		if _, ok := effective_remotes[name]; ok == false {
			repository.RemoveRemote(name)
		}

	}

	return nil

}
