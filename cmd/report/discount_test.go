package report

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/abdulrahmank/pay_later/internal/mock/dao"
	"github.com/abdulrahmank/pay_later/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetDiscounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	mDao := dao.NewMockMerchantDao(ctrl)
	cmd.MerchantDao = mDao
	merchant := &model.Merchant{Name: "m1", Discount: 20, Txns: []int{400, 500, 543}}
	mDao.EXPECT().GetMerchantDetails("m1").Return(merchant)

	actual := getDiscounts("m1")

	var expected float32
	expected = 288.6

	if expected != actual {
		t.Errorf("Expected %f but was %f", expected, actual)
	}

}
