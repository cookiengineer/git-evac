
# Architecture

## Workflow

Git Evac uses a WebView, which in return does Web Requests via the Fetch API to the local backend,
which in return executes the same actions.

An example workflow can look like this:

1. Click on a `Pull` button in the [RepositoriesTable](/app/components/RepositoriesTable.go)
2. Opens a `Dialog` with the [SchedulerTable](/app/components/SchedulerTable.go)
3. Click on `Confirm` in the Dialog calls `SchedulerTable.Start()`
4. The [Scheduler](/app/structs/Scheduler.go) method calls [actions.Pull](/app/actions/Pull.go)
5. The [actions.Pull](/app/actions/Pull.go) method requests `PATCH /api/pull/<owner>/<repository>`
6. The [routes.Pull](/source/server/routes/Pull.go) method calls `actions.Pull()`
7. The [actions.Pull](/source/actions/Pull.go) method executes the necessary git commands and handles potential error cases


## API Routes and Schemas

Idempotency and Management of State is pretty important to prevent data loss. All methods use the
correct HTTP verbs to reflect their idempotency. Each action on a repository will return an updated
[schemas.Repository](/source/schemas/Repository.go) instance.

### Basic APIs

|  ?  | Verb          | Path                             | Route                                                        | Response                                                |
|:---:|:--------------|:---------------------------------|:-------------------------------------------------------------|:--------------------------------------------------------|
| [x] | `GET`         | `/api/backups`                   | [routes.Backups](/source/server/routes/Backups.go)           | [schemas.Backups](/source/schemas/Backups.go)           |
| [x] | `GET`         | `/api/repositories`              | [routes.Repositories](/source/server/routes/Repositories.go) | [schemas.Repositories](/source/schemas/Repositories.go) |
| [ ] | `GET`         | `/api/diff/<owner>/<repository>` | [routes.Diff](/source/server/routes/Clone.go)                | [schemas.Diff](/source/schemas/Diff.go)                 |
| [ ] | `GET`, `POST` | `/api/settings`                  | [routes.Settings](/source/server/routes/Settings.go)         | [schemas.Settings](/source/schemas/Settings.go)         |

### Git Repository APIs

|  ?  | Verb            | Path                                 | Route                                                | Response                                            |
|:---:|:----------------|:-------------------------------------|:-----------------------------------------------------|:----------------------------------------------------|
| [x] | `PATCH`, `POST` | `/api/backup/<owner>/<repository>`   | [routes.Backup](/source/server/routes/Backup.go)     | [schemas.Repository](/source/schemas/Repository.go) |
| [ ] | `GET`           | `/api/clone/<owner>/<repository>`    | [routes.Clone](/source/server/routes/Clone.go)       | [schemas.Repository](/source/schemas/Repository.go) |
| [ ] | `PATCH`         | `/api/commit/<owner>/<repository>`   | [routes.Commit](/source/server/routes/Commit.go)     | [schemas.Repository](/source/schemas/Repository.go) |
| [x] | `GET`           | `/api/fix/<owner>/<repository>`      | [routes.Terminal](/source/server/routes/Terminal.go) | [schemas.Repository](/source/schemas/Repository.go) |
| [x] | `PATCH`         | `/api/pull/<owner>/<repository>`     | [routes.Pull](/source/server/routes/Pull.go)         | [schemas.Repository](/source/schemas/Repository.go) |
| [x] | `GET`           | `/api/push/<owner>/<repository>`     | [routes.Push](/source/server/routes/Push.go)         | [schemas.Repository](/source/schemas/Repository.go) |
| [x] | `PATCH`         | `/api/restore/<owner>/<repository>`  | [routes.Restore](/source/server/routes/Restore.go)   | [schemas.Repository](/source/schemas/Repository.go) |
| [x] | `GET`           | `/api/status/<owner>/<repository>`   | [routes.Status](/source/server/routes/Status.go)     | [schemas.Repository](/source/schemas/Repository.go) |
| [x] | `GET`           | `/api/terminal/<owner>/<repository>` | [routes.Terminal](/source/server/routes/Terminal.go) | [schemas.Repository](/source/schemas/Repository.go) |

