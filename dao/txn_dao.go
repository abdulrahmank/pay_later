package dao

import "fmt"

type TxnDao interface {
	CreateTxnEntry(userName, merchantName string, txnAmt int) error
}

type TxnDaoImpl struct{}

func (t *TxnDaoImpl) CreateTxnEntry(userName, merchantName string, txnAmt int) error {
	_, e := db.Exec(fmt.
		Sprintf("INSERT INTO txn (user_name, merchant_name, amt) values ('%s', '%s', '%d')", userName, merchantName, txnAmt))
	if e != nil {
		return e
	}
	return nil
}
