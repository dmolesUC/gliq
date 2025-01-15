package messages

type Issue struct {
	Iid         int64
	Title       string
	Description string
	State       string
	Labels      []string
	Milestone   *Issue
	Author      *User
	Assignees   []*User
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type User struct {
	Username string
	Name     string
}

func (u *User) String() string {
	return u.Username
}
