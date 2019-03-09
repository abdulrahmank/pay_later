package report

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/abdulrahmank/pay_later/internal/mock/dao"
	"github.com/abdulrahmank/pay_later/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetUsersAtCLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	userDao := dao.NewMockUserDao(ctrl)
	cmd.UserDao = userDao

	user1 := model.User{Name: "u1", CreditLimit: float32(400), Dues: float32(400)}
	user2 := model.User{Name: "u2", CreditLimit: float32(400), Dues: float32(200)}
	user3 := model.User{Name: "u3", CreditLimit: float32(400), Dues: float32(500)}

	userDao.EXPECT().GetAllUsers().Return([]*model.User{&user1, &user2, &user3})

	actual := getUsersAtCLimit()

	if len(actual) != 2 {
		t.Errorf("Expected %d, but was %d", 2, len(actual))
	}

	for i, val := range []string{"u1", "u3"} {
		if actual[i] != val {
			t.Errorf("Expected %s, but was %s", val, actual[i])
		}
	}
}
