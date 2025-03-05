
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


# Views

## Manage View

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


## Backup View

- [ ] Select Repositories
- [ ] Backup Action
- [ ] Dialog: Select Drives/Folders
- [ ] Dialog: Backup Progress


## Restore View

- [ ] Select Repositories
- [ ] Restore Action
- [ ] Dialog: Select Drives/Folders _or_ Remote
- [ ] Dialog: Restore Progress


## Settings View

- [ ] UI for `Folder` Setting
- [ ] UI for `Port` Setting
- [ ] UI for `Identities` Settings
- [ ] UI fro `Remotes` Settings

