package elf

import (
	"errors"

	"github.com/chapgx/elf/db"
)

type Admin struct {
	Id          int
	Username    string
	MasterKey   *string
	Fingerprint *string
}

// IsComplete determines if the root admin is completed
func (a *Admin) IsRootComplete() error {
	if a.Username != "root" {
		return errors.New("your not the root user")
	}

	if a.MasterKey == nil || a.Fingerprint == nil {
		return ErrRootIsNotComplete
	}

	return nil
}

// ReadRoot reads root admin from local database
func (a Admin) ReadRoot() (Admin, error) {
	cl := db.Connect(_dbpath)
	defer cl.Close()

	rows, e := cl.Query(`
		select *
		from admins
		where uname = 'root'
	`)
	if e != nil {
		return Admin{}, e
	}

	if !rows.Next() {
		return Admin{}, errors.New("no rows found")
	}

	var admin Admin
	e = rows.Scan(&admin.Id, &admin.Username, &admin.MasterKey, &admin.Fingerprint)

	return admin, e
}

// Inserts initial administrator into the database
func (admin Admin) init() error {
	client := db.Connect(_dbpath)
	defer client.Close()

	_, e := client.Exec(`
	insert into admins(uname)
	values('root');
	`)

	return e
}

func (admin Admin) SetKey(key string) error {
	client := db.Connect(_dbpath)
	defer client.Close()

	_, e := client.Exec(`
		update admins
		set masterkey = ?
		where uname = 'root'
		and masterkey is null
		`, key)

	return e
}
