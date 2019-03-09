package common

import (
	"github.com/abdulrahmank/pay_later/cmd/new"
	"github.com/abdulrahmank/pay_later/cmd/payback"
	"github.com/abdulrahmank/pay_later/cmd/report"
	"github.com/abdulrahmank/pay_later/cmd/update"
	"github.com/spf13/cobra"
)

var CommonCmd = &cobra.Command{}

func Execute() {
	CommonCmd.Execute()
}

func init() {
	CommonCmd.AddCommand(cmd.CmdNew)
	CommonCmd.AddCommand(payback.CmdPayback)
	CommonCmd.AddCommand(update.CmdUpdate)
	CommonCmd.AddCommand(report.CmdReport)
}
