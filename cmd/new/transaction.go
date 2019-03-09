package cmd

import (
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
	user := userDao.GetUser(userName)
	merchantName := args[1]
	amt, _ := strconv.Atoi(args[2])

	if user.CreditLimit-user.Dues >= amt {
		if e := txnDao.CreateTxnEntry(userName, merchantName, amt); e != nil {
			log.Fatal(e)
		} else {
			userDao.IncrementDues(user, amt)
		}
	}

}
