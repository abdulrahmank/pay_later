package cmd

import (
	"github.com/abdulrahmank/pay_later/internal/mock/dao"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAddNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := dao.NewMockUserDao(ctrl)
	mockDao.EXPECT().SaveUser("u1", "u1@email.in", 1000).
		Times(1)
	userDao = mockDao

	AddNewUser(nil, []string{"u1", "u1@email.in", "1000"})
}

func TestShouldNotAddNewUserWhenInputInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := dao.NewMockUserDao(ctrl)
	mockDao.EXPECT().SaveUser(gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)
	userDao = mockDao

	AddNewUser(nil, []string{"u1", "u1@email.in", "a"})
}