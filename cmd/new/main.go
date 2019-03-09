package cmd

import (
	"github.com/abdulrahmank/pay_later/dao"
	"github.com/spf13/cobra"
)

var userDao dao.UserDao

var cmdNew = &cobra.Command{
	Use:   "new",
	Short: "To add new merchants/user and for new transactions",
}

func init() {
	userDao = &dao.UserDaoImpl{}
}

func Execute() {
	cmdNew.Execute()
}
