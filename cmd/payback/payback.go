package payback

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"strconv"
)

var CmdPayback = &cobra.Command{
	Use:   "payback",
	Short: "Command to payback dues",
	Run:   Payback,
}

func Payback(_ *cobra.Command, args []string) {
	user := cmd.UserDao.GetUser(args[0])
	amt, _ := strconv.ParseFloat(args[1], 32)

	cmd.UserDao.IncrementDues(user, float32(-1 * amt))
}
