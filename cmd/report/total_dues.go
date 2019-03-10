package report

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
)

var cmdTotalDues = &cobra.Command{
	Use:   "total-dues",
	Short: "To get total dues of all the users",
	Run:   GetTotalDues,
}

func init() {
	CmdReport.AddCommand(cmdTotalDues)
}

func GetTotalDues(_ *cobra.Command, _ []string) {
	totalDues := float32(0)
	for _, val := range cmd.UserDao.GetAllUsers() {
		if val.Dues > float32(0) {
			log.Printf("%s: %f", val.Name, val.Dues)
			totalDues = totalDues + val.Dues
		}
	}
	log.Printf("total: %f", totalDues)
}
