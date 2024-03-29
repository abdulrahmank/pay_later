package cmd

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/abdulrahmank/pay_later/internal/mock/dao"
	"github.com/abdulrahmank/pay_later/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPerformNewTransaction(t *testing.T) {
	userName := "u1"
	merchantName := "m1"
	txnAmt := "900"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tDao := dao.NewMockTxnDao(ctrl)
	uDao := dao.NewMockUserDao(ctrl)
	user := model.User{Name: "u1", MailId: "u1@email.in", CreditLimit: 1000, Dues: 100}
	cmd.UserDao = uDao
	cmd.TxnDao = tDao

	tDao.EXPECT().CreateTxnEntry("u1", "m1", float32(900)).MinTimes(1)
	uDao.EXPECT().GetUser("u1").Return(&user)
	uDao.EXPECT().IncrementDues(&user, float32(900)).Times(1)

	transactionSuccess := performTransaction([]string{userName, merchantName, txnAmt})
	if transactionSuccess != nil {
		t.Errorf("Expected nil but was %v", transactionSuccess)
	}
}

func TestShouldNotPerformNewTransactionIfLowUserCredit(t *testing.T) {
	userName := "u1"
	merchantName := "m1"
	txnAmt := "900"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tDao := dao.NewMockTxnDao(ctrl)
	uDao := dao.NewMockUserDao(ctrl)
	user := &model.User{Name: "u1", MailId: "u1@email.in", CreditLimit: 1000, Dues: 200}
	cmd.UserDao = uDao
	cmd.TxnDao = tDao

	tDao.EXPECT().CreateTxnEntry(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	uDao.EXPECT().GetUser("u1").Return(user)
	uDao.EXPECT().IncrementDues(gomock.Any(), gomock.Any()).Times(0)

	transactionSuccess := performTransaction([]string{userName, merchantName, txnAmt})
	if transactionSuccess == nil {
		t.Error("Expected not nil but was nil")
	}
}
