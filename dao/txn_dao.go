package dao

import "fmt"

type TxnDao interface {
	CreateTxnEntry(userName, merchantName string, txnAmt float32) error
}

type TxnDaoImpl struct{}

func (t *TxnDaoImpl) CreateTxnEntry(userName, merchantName string, txnAmt float32) error {
	_, e := db.Exec(fmt.
		Sprintf("INSERT INTO txn (user_name, merchant_name, amt) values ('%s', '%s', '%f')", userName, merchantName, txnAmt))
	if e != nil {
		return e
	}
	return nil
}
