
# Backend

## cmds/git-evac{-debug}

- [ ] Implement parser for `$PWD/git-evac.json` into `profile.Settings`
- [ ] Implement fallback to default of parsing `$HOME/Software/<owner>/<repo>` folders

## server/api

- [ ] Migrate `api.Clone()` to `actions.Clone()` and `routes.Clone()`
- [ ] Implement `actions.Commit()` and `routes.Commit()`
- [ ] Implement `actions.Diff()` and `routes.Diff()`

## structs/Profile

- [ ] `Profile.Refresh()` needs to create new RepositoryOwner instances if there are new ones
- [ ] `Profile.Refresh()` needs to remove owners if they were deleted
- [ ] `Profile.Refresh()` needs to remove repos if they were deleted

## actions

- [ ] Implement `actions/Clone.go`
- [ ] Implement `actions/Commit.go`
- [ ] Implement `actions/Diff.go`
- [ ] Implement `actions/FixRemotes.go`


# Frontend

## app/components

- [ ] Implement BackupsTable Component

## app/controllers/Backups

- [ ] Implement `components.BackupsTable`
- [ ] Implement Dialog with `components.SchedulerTable`
- [ ] Implement Backup Action
- [ ] Implement Restore Action

## app/controllers/Settings

- [ ] Change Remote Properties (URL, Type)
- [ ] Change Identity Properties (SSH Key, User Name, User Email)
- [ ] Save Settings

