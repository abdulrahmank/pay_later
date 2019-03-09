package cmd

import "github.com/spf13/cobra"

var cmdUpdate = &cobra.Command{
	Use:   "update",
	Short: "To add new merchants/user and for new transactions",
}

func Execute() {
	cmdUpdate.Execute()
}