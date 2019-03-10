package update

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
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
	discount, e := strconv.ParseFloat(strings.Replace(args[1], "%", "", 1), 32)
	if e != nil {
		log.Fatal(e)
	}
	cmd.MerchantDao.UpdateMerchant(args[0], float32(discount))
}
