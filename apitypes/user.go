package apitypes

type User struct {
	Username string
	Name     string
}

func (u *User) String() string {
	return u.Username
}
