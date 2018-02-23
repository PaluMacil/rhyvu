package ry

import (
	"fmt"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/gob"
)

// Db is a global database object which I am likely to move or replace
// with another approach
var Db Database

type Database struct {
	*storm.DB
}

func (db Database) Open() error {
	d, err := storm.Open("ry.db", storm.Codec(gob.Codec))
	if err != nil {
		return err
	}
	db.DB = d

	err = db.Init(User{})
	if err != nil {
		return fmt.Errorf("could not register User{}: %s", err)
	}
	err = db.Init(Session{})
	if err != nil {
		return fmt.Errorf("could not register Session{}: %s", err)
	}
	err = db.Init(Group{})
	if err != nil {
		return fmt.Errorf("could not register User{}: %s", err)
	}
	err = db.Init(Permission{})
	if err != nil {
		return fmt.Errorf("could not register User{}: %s", err)
	}
	err = db.Init(GroupPermission{})
	if err != nil {
		return fmt.Errorf("could not register User{}: %s", err)
	}
	err = db.Init(UserGroup{})
	if err != nil {
		return fmt.Errorf("could not register User{}: %s", err)
	}

	return nil
}
