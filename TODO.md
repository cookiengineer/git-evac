
# Backend

## cmds/git-evac{-debug}

- [ ] Implement parser for `$PWD/git-evac.json` into `profile.Settings`
- [ ] Implement fallback to default of parsing `$HOME/Software/<owner>/<repo>` folders

## server/api

- [ ] Migrate `api.Clone()` to `actions.Clone()` and `routes.Clone()`
- [ ] Migrate `api.Pull()` to `actions.Pull()` and `routes.Pull()`
- [ ] Implement `actions.Commit()` and `routes.Commit()`
- [ ] Implement `actions.Diff()` and `routes.Diff()`

## structs/Profile

- [ ] `Profile.Refresh()` needs to create new RepositoryOwner instances if there are new ones
- [ ] `Profile.Refresh()` needs to remove owners if they were deleted
- [ ] `Profile.Refresh()` needs to remove repos if they were deleted



# Frontend

## app/actions

- [ ] Implement `actions/Diff.go` when Schema is ready
- [ ] Implement `actions/FixRemotes.go` when Workflow and Schema are ready

## app/controllers/Repositories

- [ ] Implement goroutine for Clone Action
- [ ] Implement goroutine for Fix Action
- [ ] Implement goroutine for Commit Action
- [ ] Implement goroutine for Pull Action
- [ ] Implement goroutine for Push Action

## app/controllers/Backups

- [ ] Implement Backup Action
- [ ] Implement Restore Action
- [ ] Implement `components.BackupsTable`
- [ ] Implement Dialog with `components.SchedulerTable`

## app/controllers/Settings

- [ ] Change Remote Properties (URL, Type)
- [ ] Change Identity Properties (SSH Key, User Name, User Email)
- [ ] Save Settings

