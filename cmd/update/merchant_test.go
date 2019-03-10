package update

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/abdulrahmank/pay_later/internal/mock/dao"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUpdateMerchant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := dao.NewMockMerchantDao(ctrl)
	cmd.MerchantDao = dao
	dao.EXPECT().UpdateMerchant("m1", float32(30)).Times(1)

	Update(nil, []string{"m1", "30%"})
}
