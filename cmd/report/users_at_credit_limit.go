package report

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
)

var cmdUserAtCLimit = &cobra.Command{
	Use:   "users-at-credit-limit",
	Short: "Get all the users equal to or greater than credit limit",
	Run:   GetUsersAtCLimit,
}

func init() {
	CmdReport.AddCommand(cmdUserAtCLimit)
}

func GetUsersAtCLimit(_ *cobra.Command, _ []string) {
	for _, user := range getUsersAtCLimit() {
		log.Printf("%s\n", user)
	}
}

func getUsersAtCLimit() []string {
	users := make([]string, 0)
	for _, user := range cmd.UserDao.GetAllUsers() {
		if user.Dues >= user.CreditLimit {
			users = append(users, user.Name)
		}
	}
	return users
}
