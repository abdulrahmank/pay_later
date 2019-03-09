package cmd

import (
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
	cmdNew.AddCommand(cmdNewUser)
}

func AddNewUser(_ *cobra.Command, arg []string) {
	userName := arg[0]
	emailId := arg[1]
	if creditLimit, err := strconv.Atoi(arg[2]); err != nil {
		log.Printf("Invalid input")
	} else {
		userDao.SaveUser(userName, emailId, creditLimit)
	}
}