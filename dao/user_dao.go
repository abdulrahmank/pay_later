package dao

import (
	"fmt"
	"github.com/abdulrahmank/pay_later/model"
	"log"
)

type UserDao interface {
	SaveUser(name, mailId string, creditLimit float32) error
	GetUser(name string) *model.User
	IncrementDues(user *model.User, txnAmt float32)
}

type UserDaoImpl struct{}

func (u *UserDaoImpl) SaveUser(name, mailId string, creditLimit float32) error {
	if _, e := db.Exec(fmt.Sprintf("INSERT INTO users (name, mailid, creditlimit, dues) VALUES ('%s', '%s', %f, 0)",
		name, mailId, creditLimit)); e != nil {
		return e
	}
	return nil
}

func (u *UserDaoImpl) GetUser(name string) *model.User {
	rows, e := db.Query(fmt.Sprintf("SELECT * FROM users WHERE name='%s'", name))
	if e != nil {
		return nil
	}
	rows.Next()
	var user model.User
	if err := rows.Scan(&user.Name, &user.MailId, &user.CreditLimit, &user.Dues); err != nil {
		log.Fatal(err)
		return nil
	}
	return &user
}

func (u *UserDaoImpl) IncrementDues(user *model.User, txnAmt float32) {
	_, e := db.Exec(fmt.Sprintf("UPDATE users SET dues = %f WHERE name = '%s'", user.Dues + txnAmt, user.Name))
	if e != nil {
		log.Fatal(e)
	}
}