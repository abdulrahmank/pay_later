package cmd

import (
	"github.com/spf13/cobra"
)

var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "To add new merchants/user and for new transactions",
}
