package singleton

import "sync"

type User struct {
	Nickname string
	Email    string
}

var (
	user *User
	once sync.Once
)

func NewUser() *User {
	once.Do(func() {
		user = new(User)
		user.Nickname = "guest"
		user.Email = "guest@gmail.com"
	})
	return user
}
