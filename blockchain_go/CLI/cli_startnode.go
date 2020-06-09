package CLI

import (
	"blockchain_go/BLC"
	"fmt"
	"log"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if BLC.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on.Address to receive rewards:", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	BLC.StartServer(nodeID, minerAddress)
}
