package report

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
)

var cmdDiscount = &cobra.Command{
	Use:   "discount",
	Short: "To get the latest amount of discount from a merchant",
	Run:   GetDiscounts,
}

func init() {
	CmdReport.AddCommand(cmdDiscount)
}

func GetDiscounts(_ *cobra.Command, args []string) {
	log.Println(getDiscounts(args[0]))
}

func getDiscounts(merchantName string) float32 {
	merchantDetails := cmd.MerchantDao.GetMerchantDetails(merchantName)

	totalTxnAmt := 0
	for _, val := range merchantDetails.Txns {
		totalTxnAmt += val
	}
	var totalDiscount float32
	totalDiscount = float32((float32(totalTxnAmt * merchantDetails.Discount)) / float32(100))

	return totalDiscount
}
