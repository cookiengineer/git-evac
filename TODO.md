
# TODO


# App

- [ ] Implement a `Client` that allows to use fetch asynchronously, using events and callbacks
- [ ] `Client.Request(api, func(data any, err error))` method

- [ ] Use `gooey.Fetch()` API
- [ ] Refactor `api/Repositories.go`
- [ ] Refactor `api/TerminalOpen.go`
- [ ] Implement `api/GitClone.go`
- [ ] Implement `api/GitPull.go`
- [ ] Implement `api/GitPush.go`
- [ ] Implement `api/GitCommit.go`


# Backend

- [ ] Implement `git.GlobalConfig.Parse()`
- [ ] Implement `git.LocalConfig.Parse()`
- [ ] Read `~/Software/git-evac.json` if it exists


# Views

## Manage View

- [ ] Fix Action
- [ ] Dialog: Fix Remotes workflow
- [ ] Dialog: Fix Terminal workflow

- [ ] Clone Action
- [ ] Dialog: Clone Progress

- [ ] Commit Action
- [ ] Dialog: Commit Progress

- [ ] Pull Action
- [ ] Dialog: Pull Progress

- [ ] Push Action
- [ ] Dialog: Push Progress


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

