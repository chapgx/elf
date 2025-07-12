package elf

import (
	"github.com/chapgx/elf/db"
)

type Admin struct {
	Id        int
	Username  string
	MasterKey string
}

// Inserts initial administrator into the database
func (admin Admin) init() error {
	client := db.Connect(_dbpath)
	defer client.Close()

	_, e := client.Exec(`
	insert into admins(uname)
	values('admin');
	`)

	return e
}

func (admin Admin) SetKey(key string) error {
	client := db.Connect(_dbpath)
	defer client.Close()

	_, e := client.Exec(`
		update admins
		set masterkey = ?
		where uname = 'admin'
		and masterkey is null
		`, key)

	return e
}
