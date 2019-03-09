package report

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
)

var cmdDues = &cobra.Command{
	Use:   "dues",
	Short: "To get the due amount of a user",
	Run:   GetDues,
}

func init() {
	CmdReport.AddCommand(cmdDues)
}

func GetDues(_ *cobra.Command, args []string) {
	log.Println(getDues(args[0]))
}

func getDues(username string) float32 {
	return cmd.UserDao.GetUser(username).Dues
}
