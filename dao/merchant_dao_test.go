package dao

import (
	"fmt"
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

func TestShouldUpdateMerchantDiscounts(t *testing.T) {
	dao := MerchantDaoImpl{}
	_ = dao.SaveMerchant("m1", 10)

	dao.UpdateMerchant("m1", 20)

	rows, _ := db.Query(fmt.Sprintf("select discount from merchants where name='%s'", "m1"))
	rows.Next()
	var discount int
	rows.Scan(&discount)

	if discount != 20 {
		t.Errorf("Expected %d, but was %d", 20, discount)
	}
}

func TestGetMerchantDetails(t *testing.T) {
	dao := MerchantDaoImpl{}
	_ = dao.SaveMerchant("m1", 10)
	_, e := db.Exec("Insert into txn (user_name, merchant_name, amt) values ('u1', 'm1', 100)")
	if e != nil {
		log.Fatal(e)
	}
	db.Exec("Insert into txn (user_name, merchant_name, amt) values ('u1', 'm1', 300)")
	db.Exec("Insert into txn (user_name, merchant_name, amt) values ('u1', 'm1', 500)")

	merchantDetails := dao.GetMerchantDetails("m1")

	if merchantDetails.Name != "m1" {
		t.Errorf("Expected %s, but was %s", "m1", merchantDetails.Name)
	}
	if merchantDetails.Discount != 10 {
		t.Errorf("Expected %d, but was %d", 10, merchantDetails.Discount)
	}

	expected := []int{100, 300, 500}
	for i, val := range expected {
		if merchantDetails.Txns[i] != val {
			t.Errorf("Expected %v, but was %v", "m1", merchantDetails.Txns)
			break
		}
	}
	db.Exec("truncate merchants")
	db.Exec("truncate txn")
}
