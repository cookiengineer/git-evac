
# TODO

Table like this:

actions can be:

- if local changes: commit
- if no local changes: pull, push
- if no github/gitlab/homeserver remote, then fix remotes

| organization    | repository | remotes                  | branch | status        | actions  |
| tholian-network | endpoint   | github gitlab homeserver | master | local changes | [commit] |


Profile Settings have to have users and organizations map to remotes,
meaning `map[string]*[]Remote` or something like that.


# TODO

- [ ] Implement `git.GlobalConfig.Parse()`
- [ ] Implement `git.LocalConfig.Parse()`

- [ ] Read `structs.Profile` from `~/Software/git-evac.json` if it exists already

- [ ] Implement Settings UI to configure multiple git/gogs/github/gitlab servers
- [ ] Implement `structs.Server`


# Synchronize View

- Select Repositories
- Select Remotes
- Synchronize Progress


# Backup View

- Select Repositories
- Select Drives/Folders
- Backup Progress


# Restore View

- Select Drives/Folders
- Select Repositories
- Restore Progress
