
# TODO


# Backend

- [ ] `Profile.Refresh()` needs to create new RepositoryOwner instances if there are new ones
- [ ] `Profile.Refresh()` needs to remove owners if they were deleted
- [ ] `Profile.Refresh()` needs to remove repos if they were deleted

- [ ] Finish Implementation of `server/api/Restore.go`
- [ ] Finish Implementation of `server/api/Pull.go`


- [ ] Parse `~/Software/git-evac.json` as `Profile` if it exists, do that in `main.go`
- [ ] Implement `git.GlobalConfig.Parse()`
- [ ] Implement `git.LocalConfig.Parse()`

- [ ] Implement `api/Diff.go`
- [ ] Implement `api/Commit.go`


# App

- [ ] Implement `actions/Diff.go` when Schema is ready
- [ ] Implement `actions/FixRemotes.go` when Workflow and Schema are ready

## Repositories View

- [ ] Fix Action
- [ ] Dialog: Fix Remotes workflow
- [ ] Dialog: Fix Terminal workflow

- [ ] Diff Action
- [ ] Dialog: Show Diff of Files
- [ ] Dialog: Allow to check/select committed files
- [ ] Dialog: Show label for "Commit x files, delete y files" depending on selection

- [ ] Commit Action
- [ ] Dialog: Commit Progress

- [ ] Pull Action
- [ ] Dialog: Pull Progress


## Backups View

- [ ] Select Repositories
- [ ] Backup Action
- [ ] Dialog: Backup Progress

- [ ] Select Repositories
- [ ] Restore Action
- [ ] Dialog: Restore Progress

## Settings View

- [ ] Change Remote Properties (URL, Type)
- [ ] Change Identity Properties (SSH Key, User Name, User Email)
- [ ] Save Settings

