package schemas

import "time"

type Repository struct {

	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	FullName      string    `json:"full_name"`
	Description   string    `json:"description"`

	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DefaultBranch string    `json:"default_branch"`
	Size          int64     `json:"size"`

	Fork          bool      `json:"fork"`
	Private       bool      `json:"private"`

	WebURL        string    `json:"html_url"`
	SSHURL        string    `json:"ssh_url"`

	Permissions struct {
		Admin bool `json:"admin"`
		Pull  bool `json:"pull"`
		Push  bool `json:"push"`
	} `json:"permissions"`

}

