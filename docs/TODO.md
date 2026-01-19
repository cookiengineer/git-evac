
# git-evac

## Stories

- [ ] Implement a View that compares remotes and their remote URLs according
      to the settings schema. This View should offer to use the `FixRemote`
      step if the remotes don't match the schema.
- [ ] Implement a `FixRemote` action for remotes which are not matching the
      settings schema.

## Backend

### types

- [ ] Implement `types/Identity.go` method `IsValid()`

### server

- [ ] Implement `server/DispatchRoutes.go` route `POST /api/commit`
- [ ] Implement `server/DispatchRoutes.go` route `GET /api/diff`
- [ ] Implement `services/github/FetchRepositories.go`
- [ ] Implement `services/gitlab/FetchRepositories.go`
- [ ] Implement `services/gitea/FetchRepositories.go`

### server/routes

- [ ] Implement `routes.Commit()`
- [ ] Implement `routes.Diff()`

### structs

- [ ] `Profile.Refresh()` needs to create new RepositoryOwner instances if there are new ones
- [ ] `Profile.Refresh()` needs to remove owners if they were deleted
- [ ] `Profile.Refresh()` needs to remove repos if they were deleted

### actions

- [ ] Implement `actions/Commit.go`
- [ ] Implement `actions/Diff.go`
- [ ] Implement `actions/FixRemotes.go`


## Frontend

### app/components

- [ ] Implement `RemotesTable` Component

### app/controllers/Settings

- [ ] Implement `public/settings.html` Controller and View
- [ ] Change Remote Properties (URL, Type)
- [ ] Change Identity Properties (SSH Key, User Name, User Email)
- [ ] Save Settings

### app/structs

- [ ] SchedulerTable Dialog needs to refresh `repo.Status()` of Table after actions are done.
      This is already inside the fetch response, but not written to the storage.

