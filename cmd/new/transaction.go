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
	CmdNew.AddCommand(cmdNewTxn)
}

func PerformTransaction(_ *cobra.Command, args []string) {
	userName := args[0]
	user := cmd.UserDao.GetUser(userName)
	merchantName := args[1]
	amt, _ := strconv.ParseFloat(args[2], 32)
	amtf := float32(amt)

	if user.CreditLimit-user.Dues >= amtf {
		if e := cmd.TxnDao.CreateTxnEntry(userName, merchantName, amtf); e != nil {
			log.Fatal(e)
		} else {
			cmd.UserDao.IncrementDues(user, amtf)
		}
	}

}
