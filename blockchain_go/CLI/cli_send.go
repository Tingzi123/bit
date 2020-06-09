package CLI

import (
	"blockchain_go/BLC"
	"log"
	"fmt"
)

//发送交易
func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !BLC.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !BLC.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}
	bc := BLC.NewBlockchain(nodeID)
	UTXOSet := BLC.UTXOSet{bc}
	defer bc.Db.Close()

	wallets, err := BLC.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := BLC.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)
	if mineNow {
		cbTx := BLC.NewCoinbaseTX(from, "")
		txs := []*BLC.Transaction{cbTx, tx}
		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		BLC.SendTx(BLC.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
