package structs

import "git-evac/types"
import "encoding/json"
import "testing"

func TestSettingsRoundTrip(t *testing.T) {

	settings := NewSettings(t.TempDir(), t.TempDir(), 3000)

	owner := NewSettingsOwner("cookiengineer")

	identity := types.NewIdentity("work")
	if identity.SetSSHKey("/home/user/.ssh/id_work") == false {
		t.Fatal("ssh key invalid")
	}
	if identity.SetGitUserName("Cookie Engineer") == false {
		t.Fatal("name invalid")
	}
	if identity.SetGitUserEmail("cookie@example.com") == false {
		t.Fatal("email invalid")
	}
	owner.SetIdentity(identity)
	owner.DefaultIdentity = "work"

	origin := types.NewRemote("origin", "ssh://git@homeserver:3001/{{owner}}/{{repo}}.git")
	github := types.NewRemote("github", "ssh://git@github.com/{{owner}}/{{repo}}.git")
	if origin.IsValidSchema() == false {
		t.Fatal("origin schema invalid")
	}
	if github.IsValidSchema() == false {
		t.Fatal("github schema invalid")
	}
	if github.ToURL("cookiengineer", "git-evac") != "ssh://git@github.com/cookiengineer/git-evac.git" {
		t.Fatal("ToURL failed: " + github.ToURL("cookiengineer", "git-evac"))
	}
	owner.SetRemote(origin)
	owner.SetRemote(github)

	pinned := NewSettingsRepository("private-thing")
	pinned.SetRemote(nil)
	pinned.SetIdentity(nil)
	// remove github for this private repo
	pinned.Remotes["github"] = nil
	owner.SetRepository(pinned)

	settings.Owners["cookiengineer"] = owner

	if settings.IsValid() == false {
		t.Fatal("settings invalid")
	}

	payload, err := json.Marshal(settings)
	if err != nil {
		t.Fatal(err)
	}

	decoded := NewSettings("", "", 0)
	err = json.Unmarshal(payload, decoded)
	if err != nil {
		t.Fatal(err)
	}

	decoded_owner := decoded.GetOwner("cookiengineer")
	if decoded_owner == nil {
		t.Fatal("owner lost")
	}
	if decoded_owner.DefaultIdentity != "work" {
		t.Fatal("default identity lost")
	}
	if decoded_owner.GetIdentity("work") == nil {
		t.Fatal("identity lost")
	}

	effective := decoded_owner.Effective("private-thing")
	if effective == nil {
		t.Fatal("effective nil")
	}
	if _, ok := effective.SnapshotRemotes()["github"]; ok == true {
		t.Fatal("pinned nil removal did not remove github")
	}
	if _, ok := effective.SnapshotRemotes()["origin"]; ok == false {
		t.Fatal("origin missing from effective")
	}

	normal := decoded_owner.Effective("public-thing")
	if _, ok := normal.SnapshotRemotes()["github"]; ok == false {
		t.Fatal("github missing from non-pinned effective")
	}

}
