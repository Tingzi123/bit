package CLI

import (
	"blockchain_go/BLC"
	"fmt"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := BLC.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address:%s\n", address)
}
