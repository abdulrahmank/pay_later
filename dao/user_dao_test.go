package dao

import (
	"github.com/abdulrahmank/pay_later/model"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PORT", "5432")
	os.Setenv("DATABASE_USER", "test")
	os.Setenv("DATABASE_DB", "test_pay_later")
	os.Setenv("DATABASE_PSSWD", "password")
	Reinit()

	m.Run()

	db.Exec("TRUNCATE users")
}

func TestShouldSaveNewUser(t *testing.T) {
	dao := UserDaoImpl{}
	e := dao.SaveUser("u1", "u1@email.in", 1000)

	var count int
	rows, _ := db.Query("Select count(*) from users where name = 'u1'")

	rows.Next()
	if e = rows.Scan(&count); e != nil {
		log.Fatal(e)
	}

	if e != nil {
		t.Errorf("Expected %v, but was %v", nil, e)
	}
	if count != 1 {
		t.Errorf("Expected %d, but was %d", 1, count)
	}
}

func TestShouldGetUser(t *testing.T) {
	dao := UserDaoImpl{}
	dao.SaveUser("u1", "u1@email.in", 1000)

	user := dao.GetUser("u1")

	if user.Name != "u1" {
		t.Errorf("Expected this %s but was %s", "u1", user.Name)
	}
	if user.MailId != "u1@email.in" {
		t.Errorf("Expected this %s but was %s", "u1@email.in", user.MailId)
	}
	if user.CreditLimit != 1000 {
		t.Errorf("Expected this %d but was %f", 1000, user.CreditLimit)
	}
}

func TestShouldGetAllUsers(t *testing.T) {
	dao := UserDaoImpl{}
	names := []string{"u1", "u2"}
	mailIds := []string{"u1@email.in", "u2@email.in"}
	creditLimits := []float32{1000, 1000}
	dao.SaveUser("u1", "u1@email.in", 1000)
	dao.SaveUser("u2", "u2@email.in", 1000)

	users := dao.GetAllUsers()

	for i, user := range users {
		if user.Name != names[i] {
			t.Errorf("Expected this %s but was %s", names[i], user.Name)
		}
		if user.MailId != mailIds[i] {
			t.Errorf("Expected this %s but was %s", mailIds[i], user.MailId)
		}
		if user.CreditLimit != creditLimits[i] {
			t.Errorf("Expected this %d but was %f", creditLimits[i], user.CreditLimit)
		}
	}
}

func TestShouldIncrementDueAmount(t *testing.T) {
	dao := UserDaoImpl{}
	dao.SaveUser("u1", "u1@email.in", 1000)

	dao.IncrementDues(&model.User{Name:"u1", Dues:200}, 200)

	user := dao.GetUser("u1")
	if user.Dues != 400 {
		t.Errorf("Expected %d but was %f", 200, user.Dues)
	}
}
