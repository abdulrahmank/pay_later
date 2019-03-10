package payback

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
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

	if err := cmd.UserDao.IncrementDues(user, float32(-1*amt)); err == nil {
		log.Printf("%s(dues: %f)", user.Name, float32(user.Dues-float32(amt)))
	}
}
