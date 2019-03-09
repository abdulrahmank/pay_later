package dao

import (
	"log"
	"os"
	"testing"
)

func TestShouldSaveNewMerchant(t *testing.T) {
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PORT", "5432")
	os.Setenv("DATABASE_USER", "test")
	os.Setenv("DATABASE_DB", "test_pay_later")
	os.Setenv("DATABASE_PSSWD", "password")
	Reinit()
	t.Run("Should be able to save new merchant", func(t *testing.T) {
		dao := MerchantDaoImpl{}
		e := dao.SaveMerchant("m1", 10)

		var count int
		rows, _ := db.Query("Select count(*) from merchants where name = 'm1'")
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
	})

	db.Exec("TRUNCATE merchants")
}
