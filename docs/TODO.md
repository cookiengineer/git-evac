
# git-evac

## Remotes

### app/components

- [ ] Implement `RemotesTable` Component for the "Remotes" View
- [ ] Implement a `RemotesDialog` Component for the remotes view

### app/views and app/controllers

- [ ] Implement a "Remotes" View and Remotes Controller that compares remotes
      and their remote URLs according to the settings schema. This View should
      offer to use the `FixRemote` step if the remotes don't match the schema.
- [ ] Integrate the view to the app

### app/actions

- [x] Implement a `FixRemotes` action for remotes which are not matching the
      settings schema, where e.g. `origin`, `github`, and `gitlab` are universally
      set based on the schema's URL templates.
- [x] Implement a `FixIdentity` action for identities which are not matching the
      settings schema, applying `user.name`, `user.email` and `core.sshCommand`.
- [ ] Implement a "Remotes" View that offers the `FixRemotes` step for the
      repositories whose remotes don't match the schema.

### app/controllers/Settings

- [x] Implement `public/settings.html` Controller and View
- [x] Change Remote Properties (URL, Type)
- [x] Change Identity Properties (SSH Key, User Name, User Email)
- [x] Save Settings
- [x] Persist settings to the `--config` file path
- [x] Pin per-repository remote/identity overrides from the Repositories view



## Backend

### types

- [x] Implement `types/Identity.go` method `IsValid()`

### server

- [ ] Implement `server/DispatchRoutes.go` route `POST /api/commit`
- [ ] Implement `server/DispatchRoutes.go` route `GET /api/diff`
- [x] Implement `server/DispatchRoutes.go` route `PATCH /api/fixremotes/<owner>/<repository>`
- [x] Implement `server/DispatchRoutes.go` route `PATCH /api/fixidentity/<owner>/<repository>`

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

