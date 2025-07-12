package elf

type User struct {
	Username string
	Passwd   Password
}

func (u *User) ReadUserFromAdmin(a Admin) {
	u.Username = a.Username
}

func (u User) IsNil() bool {
	if u.Username == "" {
		return true
	}
	return false
}
