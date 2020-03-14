package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/ddaws/go-maker/maker"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	// Connect to Infura
	infuraAddr := fmt.Sprintf(infuraWebSocketAddr, *network, *projectID)
	client, err := ethclient.Dial(infuraAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Connected to Infura!")
	// Subscribe to log events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{maker.FlipEthAAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Subscribed to Flipper events!")
	// Load the Flipper ABI for decoding flipper events
	flipAbi, err := abi.JSON(strings.NewReader(string(maker.FlipperABI)))
	if err != nil {
		log.Fatalln(err)
	}
	// Define a new template for formatting Kick events
	kickTmpl, err := template.New("kick").Parse(`
Kick({{ .Id }}):
	Lot {{ .Lot }}
	Bid {{ .Big }}
	Tab {{ .Tab }}
	Usr {{ .Usr }}
	Gal {{ .Gal }} 
`)
	if err != nil {
		log.Fatalln(err)
	}
	// Create a timer to print heartbeat messages
	heartbeat := time.Tick(10 * time.Second)
	// Start processing flip events
	for {
		select {
		case err := <-sub.Err():
			log.Fatalln(err)
		case vLog := <-logs:
			var kick maker.FlipperKick
			err = flipAbi.Unpack(&kick, "Kick", vLog.Data)
			if err != nil {
				log.Fatalln(err)
			}
			kickTmpl.Execute(os.Stdout, kick)
		case <-heartbeat:
			log.Println("Heartbeat...")
		}
	}
}
