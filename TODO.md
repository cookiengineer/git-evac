
# TODO


# App

- [ ] Implement a `Client` that allows to use fetch asynchronously, using events and callbacks
- [ ] `Client.Request(api, func(data any, err error))` method

- [ ] Use `gooey.Fetch()` API
- [ ] Implement `api/Restore.go` to use tar.gz
- [ ] Implement `api/Diff.go`
- [ ] Implement `api/Pull.go`
- [ ] Implement `api/Commit.go`


# Backend

- [ ] Implement `git.GlobalConfig.Parse()`
- [ ] Implement `git.LocalConfig.Parse()`
- [ ] Read `~/Software/git-evac.json` if it exists


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

