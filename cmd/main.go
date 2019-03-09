package cmd

import "github.com/abdulrahmank/pay_later/dao"

var UserDao dao.UserDao
var MerchantDao dao.MerchantDao
var TxnDao dao.TxnDao

func init() {
	UserDao = &dao.UserDaoImpl{}
	MerchantDao = &dao.MerchantDaoImpl{}
	TxnDao = &dao.TxnDaoImpl{}
}