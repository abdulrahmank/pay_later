package dao

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestShouldAddTransactionEntry(t *testing.T) {
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PORT", "5432")
	os.Setenv("DATABASE_USER", "test")
	os.Setenv("DATABASE_DB", "test_pay_later")
	os.Setenv("DATABASE_PSSWD", "password")
	Reinit()
	t.Run("Should be able to save new transaction", func(t *testing.T) {
		userName := "u1"
		merchantName := "m1"
		txnAmt := 300
		dao := TxnDaoImpl{}

		dao.CreateTxnEntry(userName, merchantName, txnAmt)

		rows, e := db.Query(fmt.Sprintf("Select count(*) from txn where user_name = '%s' and merchant_name = '%s'",
			userName, merchantName))
		if e != nil {
			log.Fatal(e)
		}
		var count int
		rows.Next()
		rows.Scan(&count)
		if count != 1 {
			t.Errorf("Expected %d but was %d", 1, count)
		}
	})
	db.Exec("truncate txn")
}
