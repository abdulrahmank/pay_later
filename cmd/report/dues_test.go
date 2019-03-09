package report

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/abdulrahmank/pay_later/internal/mock/dao"
	"github.com/abdulrahmank/pay_later/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetDues(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDao := dao.NewMockUserDao(ctrl)
	cmd.UserDao = userDao

	var expected float32 = 1000
	user := model.User{Dues: expected}
	userDao.EXPECT().GetUser("u1").Return(&user)

	actual := getDues("u1")

	if actual != expected {
		t.Errorf("Expected %f but was %f\n", expected, actual)
	}
}

