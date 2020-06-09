package CLI

import (
	"blockchain_go/BLC"
	"fmt"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := BLC.NewBlockchain(nodeID)
	UTXOSet := BLC.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
