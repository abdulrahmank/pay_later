package cmd

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/abdulrahmank/pay_later/internal/mock/dao"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAddNewMerchant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := dao.NewMockMerchantDao(ctrl)
	cmd.MerchantDao = mockDao
	mockDao.EXPECT().SaveMerchant("m1", float32(20)).Times(1)

	AddNewMerchant(nil, []string{"m1", "20%"})
}

func TestShouldNotAddNewMerchantIfInvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := dao.NewMockMerchantDao(ctrl)
	cmd.MerchantDao = mockDao
	mockDao.EXPECT().SaveMerchant(gomock.Any(), gomock.Any()).Times(0)

	AddNewMerchant(nil, []string{"m1", "%"})
}
