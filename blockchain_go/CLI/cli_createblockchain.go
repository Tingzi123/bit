package CLI

import (
	"blockchain_go/BLC"
	"log"
	"fmt"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !BLC.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := BLC.CreateBlockchain(address, nodeID)
	defer bc.Db.Close()

	UTXOSet := BLC.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
