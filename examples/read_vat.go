package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ddaws/go-maker/maker"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	InfuraWebSocketAddr = "wss://%s.infura.io/ws/v3/%s"
	VatMessageTemplate  = `

Vat @ %s
----------
  debt: %s
  vice: %s
  line: %s

`
)

func main() {
	var (
		network   = flag.String("network", "mainnet", "The Ethereum network on Infura to connect to")
		projectID = flag.String("project", "", "Your Infura project ID for connecting to and querying Infura")
	)
	flag.Parse()

	if projectID == nil || *projectID == "" {
		log.Fatalln("You must specify a valid project ID")
	}

	infuraAddr := fmt.Sprintf(InfuraWebSocketAddr, *network, *projectID)
	client, err := ethclient.Dial(infuraAddr)
	if err != nil {
		log.Fatalln(err)
	}

	vat, err := maker.LoadVat(client)
	if err != nil {
		log.Fatalln(err)
	}

	// You shouldn't ignore errors like this!
	var (
		debt, _ = vat.Debt(nil)
		vice, _ = vat.Vice(nil)
		line, _ = vat.Line(nil)
	)

	log.Printf(VatMessageTemplate, maker.VatAddress.Hex(), debt, vice, line)
}
