package payback

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"strconv"
)

var cmdPayback = &cobra.Command{
	Use:   "payback",
	Short: "Command to payback dues",
	Run:   Payback,
}

func Payback(_ *cobra.Command, args []string) {
	user := cmd.UserDao.GetUser(args[0])
	amt, _ := strconv.Atoi(args[1])

	cmd.UserDao.IncrementDues(user, -1*amt)
}
