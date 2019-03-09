package cmd

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var cmdNewUser = &cobra.Command{
	Use:   "user",
	Run:   AddNewUser,
	Short: "To add a new user",
}

func init() {
	CmdNew.AddCommand(cmdNewUser)
}

func AddNewUser(_ *cobra.Command, arg []string) {
	userName := arg[0]
	emailId := arg[1]
	if creditLimit, err := strconv.ParseFloat(arg[2], 32); err != nil {
		log.Printf("Invalid input")
	} else {
		cmd.UserDao.SaveUser(userName, emailId, float32(creditLimit))
	}
}
