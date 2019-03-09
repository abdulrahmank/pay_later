package dao

import "fmt"

type MerchantDao interface {
	SaveMerchant(name string, discount int) error
}

type MerchantDaoImpl struct{}

func (m *MerchantDaoImpl) SaveMerchant(name string, discount int) error {
	if _, e := db.Exec(fmt.Sprintf("INSERT INTO merchants VALUES ('%s', %d)", name, discount)); e != nil {
		return e
	}
	return nil
}
