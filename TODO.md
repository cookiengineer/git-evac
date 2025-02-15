
# TODO

Overview Actions can be:

- if local changes: "commit"
- if no local changes: "pull" or "push"
- if no github/gitlab/homeserver remote: "fix remotes"


# TODO

- [ ] Implement `git.GlobalConfig.Parse()`
- [ ] Implement `git.LocalConfig.Parse()`

- [ ] Read `Profile.Settings` from `~/Software/git-evac.json` if it exists already

- [ ] Implement Settings UI to configure multiple git/gogs/github/gitlab servers
- [ ] Implement `structs.Server` to reflect GitHub / GitLab and Gogs/Gitea APIs


# Manage View

- [ ] Select Repositories

- [ ] Commit Action
- [ ] Dialog: Open Terminal, then call `/api/status/...`

- [ ] Pull Action
- [ ] Dialog: Select Remotes
- [ ] Dialog: Pull Progress

- [ ] Push Action
- [ ] Dialog: Select Remotes
- [ ] Dialog: Push Progress


# Backup View

- [ ] Select Repositories
- [ ] Backup Action
- [ ] Dialog: Select Drives/Folders
- [ ] Dialog: Backup Progress


# Restore View

- [ ] Select Repositories
- [ ] Restore Action
- [ ] Dialog: Select Drives/Folders _or_ Remote
- [ ] Dialog: Restore Progress

