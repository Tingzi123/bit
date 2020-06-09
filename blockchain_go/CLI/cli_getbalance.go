package CLI

import (
	"blockchain_go/BLC"
	"blockchain_go/UTIL"
	"fmt"
	"log"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !BLC.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	bc := BLC.NewBlockchain(nodeID)
	UTXOSet := BLC.UTXOSet{bc}
	defer bc.Db.Close()

	balance := 0
	pubKeyHash := UTIL.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s':%d\n", address, balance)
}
