package update

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var cmdUpdateMerchant = &cobra.Command{
	Use:   "merchant",
	Short: "To update merchant discounts",
	Run:   Update,
}

func init() {
	CmdUpdate.AddCommand(cmdUpdateMerchant)
}

func Update(_ *cobra.Command, args []string) {
	discount, e := strconv.Atoi(args[1])
	if e != nil {
		log.Fatal(e)
	}
	cmd.MerchantDao.UpdateMerchant(args[0], discount)
}
