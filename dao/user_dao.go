package dao

import "fmt"

type UserDao interface {
	SaveUser(name, mailId string, creditLimit int) error
}

type UserDaoImpl struct{}

func (u *UserDaoImpl) SaveUser(name, mailId string, creditLimit int) error {
	if _, e := db.Exec(fmt.Sprintf("INSERT INTO users VALUES ('%s', '%s', %d)", name, mailId, creditLimit)); e != nil {
		return e
	}
	return nil
}
