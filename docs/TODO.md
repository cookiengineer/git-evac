
# git-evac

## Remotes

### app/components

- [ ] Implement `RemotesTable` Component for the "Remotes" View

### app/views

- [ ] Implement a "Remotes" View that compares remotes and their remote URLs
      according to the settings schema. This View should offer to use the
      `FixRemote` step if the remotes don't match the schema.

### app/actions

- [ ] Implement a `FixRemote` action for remotes which are not matching the
      settings schema, where e.g. `origin`, `github`, and `gitlab` are universally
      set based on the schema's URL templates.

### app/controllers/Settings

- [ ] Implement `public/settings.html` Controller and View
- [ ] Change Remote Properties (URL, Type)
- [ ] Change Identity Properties (SSH Key, User Name, User Email)
- [ ] Save Settings



## Backend

### types

- [ ] Implement `types/Identity.go` method `IsValid()`

### server

- [ ] Implement `server/DispatchRoutes.go` route `POST /api/commit`
- [ ] Implement `server/DispatchRoutes.go` route `GET /api/diff`

### server/routes

- [ ] Implement `routes.Commit()`
- [ ] Implement `routes.Diff()`

### actions

- [ ] Implement `actions/Commit.go`
- [ ] Implement `actions/Diff.go`
- [ ] Implement `actions/FixRemotes.go`


## Frontend

### app/structs

- [ ] SchedulerTable Dialog needs to refresh `repo.Status()` of Table after actions are done.
      This is already inside the fetch response, but not written to the storage.

