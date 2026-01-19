package schemas

import "time"

type Repository struct {

	ID                 int64     `json:"id"`
	Name               string    `json:"name"`
	Path               string    `json:"path"`
	PathWithNamespace  string    `json:"path_with_namespace"`
	Description        string    `json:"description"`

	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"last_activity_at"`
	DefaultBranch      string    `json:"default_branch"`

	ForkedFromProject  *struct {
		ID int64 `json:"id"`
	} `json:"forked_from_project"`

	Visibility         string    `json:"visibility"`
	Private            bool      `json:"private"`

	WebURL             string    `json:"web_url"`
	SSHURL             string    `json:"ssh_url_to_repo"`

}

