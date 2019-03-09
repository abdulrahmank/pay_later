package cmd

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var cmdNewTxn = &cobra.Command{
	Use:   "txn",
	Run:   PerformTransaction,
	Short: "To perform a new transaction",
}

func init() {
	cmdNew.AddCommand(cmdNewTxn)
}

func PerformTransaction(_ *cobra.Command, args []string) {
	userName := args[0]
	user := cmd.UserDao.GetUser(userName)
	merchantName := args[1]
	amt, _ := strconv.Atoi(args[2])

	if user.CreditLimit-user.Dues >= amt {
		if e := cmd.TxnDao.CreateTxnEntry(userName, merchantName, amt); e != nil {
			log.Fatal(e)
		} else {
			cmd.UserDao.IncrementDues(user, amt)
		}
	}

}
