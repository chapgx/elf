package elf

import (
	"errors"

	"github.com/chapgx/elf/db"
	"github.com/chapgx/morm"
)

type Admin struct {
	Id          int `morm:"id int primary key"`
	Username    string
	MasterKey   *string `morm:"master_key text"`
	Fingerprint *string `morm:"finger_print text"`
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
	if admin.Username != "" {
		return errors.New("username must be <nil> for the init function")
	}

	// TODO: need to read first to make sure root does not exists

	admin.Username = "root"
	e := morm.Insert(&admin)

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
