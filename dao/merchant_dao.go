package dao

import (
	"fmt"
	"github.com/abdulrahmank/pay_later/model"
	"github.com/lib/pq"
	"log"
	"strconv"
)

type MerchantDao interface {
	SaveMerchant(name string, discount float32) error
	UpdateMerchant(name string, discount float32) error
	GetMerchantDetails(name string) *model.Merchant
}

type MerchantDaoImpl struct{}

func (m *MerchantDaoImpl) SaveMerchant(name string, discount float32) error {
	if _, e := db.Exec(fmt.Sprintf("INSERT INTO merchants VALUES ('%s', %d)", name, discount)); e != nil {
		return e
	}
	return nil
}

func (m *MerchantDaoImpl) UpdateMerchant(name string, discount float32) error {
	if _, e := db.Exec(fmt.Sprintf("UPDATE merchants SET discount = %d WHERE name = '%s'", discount, name)); e != nil {
		return e
	}
	return nil
}

func (m *MerchantDaoImpl) GetMerchantDetails(name string) *model.Merchant {
	rows, e := db.Query(fmt.Sprintf("SELECT m.name, m.discount, array_agg(t.amt) "+
		"FROM merchants m LEFT JOIN txn t on m.name = t.merchant_name WHERE m.name = '%s' GROUP BY m.name, m.discount", name))
	if e != nil {
		return nil
	}
	rows.Next()
	merchant := &model.Merchant{}
	txns := make([]string, 5)
	if err := rows.Scan(&merchant.Name, &merchant.Discount, pq.Array(&txns)); err != nil {
		log.Fatal(err.Error())
		return nil
	} else {
		for _, txn := range txns {
			amt, _ := strconv.Atoi(txn)
			merchant.Txns = append(merchant.Txns, amt)
		}
		return merchant
	}
}
