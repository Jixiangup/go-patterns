package singleton

import "log"

func Usage() {
	user := NewUser()
	log.Println(user)
}
