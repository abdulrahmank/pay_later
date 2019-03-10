package cmd

import (
	"github.com/abdulrahmank/pay_later/cmd"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var cmdNewMerchant = &cobra.Command{
	Use:   "merchant",
	Run:   AddNewMerchant,
	Short: "To add a new merchant",
}

func init() {
	CmdNew.AddCommand(cmdNewMerchant)
}

func AddNewMerchant(_ *cobra.Command, arg []string) {
	merchantName := arg[0]
	discount := arg[1]
	discount = strings.Replace(discount, "%", "", 1)
	if discountNum, err := strconv.ParseFloat(discount, 32); err != nil {
		log.Println("Invalid discount")
	} else {
		cmd.MerchantDao.SaveMerchant(merchantName, float32(discountNum))
	}
}
