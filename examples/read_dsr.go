package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ddaws/go-maker/maker"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	infuraWebSocketAddr = "wss://%s.infura.io/ws/v3/%s"
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
	log.Print("Connected to Infura!")

	pot, err := maker.LoadPot(client)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Loaded the Pot!")

	// You shouldn't ignore errors like this!
	dsr, err := pot.Dsr(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Loaded the DSR!")
	log.Printf("DSR (per second): %s", dsr.Text(10))

	// Convert to float64. This makes our calculations simpler, but much less precise!
	rateScaled, _ := new(big.Float).SetInt(dsr).Float64()
	perSecondRate := rateScaled / math.Pow(10, maker.RayScale)
	secondsInYear := float64(365 * 24 * 60 * 60)

	// Round to the nearest 10th of a %
	approxAnnualizedRate := math.Pow(perSecondRate, secondsInYear)
	annualizedRate := math.Round(approxAnnualizedRate*1000) / 1000

	log.Printf("DSR (annualized): %.3f", annualizedRate)
}
