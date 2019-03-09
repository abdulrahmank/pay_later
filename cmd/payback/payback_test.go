package payback

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/abdulrahmank/pay_later/internal/mock/dao"
	"github.com/abdulrahmank/pay_later/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPayback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userDao := dao.NewMockUserDao(ctrl)
	cmd.UserDao = userDao

	user := model.User{Name: "u1", Dues: 900}
	userDao.EXPECT().GetUser("u1").Return(&user)
	userDao.EXPECT().IncrementDues(&user, -700).Times(1)

	Payback(nil, []string{"u1", "700"})
}

