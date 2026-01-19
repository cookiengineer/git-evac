
# TODO

## Backend

- [ ] Implement `types/Identity.go` method `IsValid()`
- [ ] Implement `server/DispatchRoutes.go` route `POST /api/commit`
- [ ] Implement `server/DispatchRoutes.go` route `GET /api/diff`
- [ ] Implement `services/github/FetchRepositories.go`
- [ ] Implement `services/gitlab/FetchRepositories.go`
- [ ] Implement `services/gitea/FetchRepositories.go`

## Frontend

- [ ] SchedulerTable Dialog needs to refresh `repo.Status()` of Table after actions are done.
      This is already inside the fetch response, but not written to the storage.
- [ ] `public/settings.html`: Implement Settings View

