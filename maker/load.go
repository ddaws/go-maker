package maker

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadVat(client *ethclient.Client) (*Vat, error) {
	return NewVat(VatAddress, client)
}

func LoadVatCaller(client *ethclient.Client) (*VatCaller, error) {
	return NewVatCaller(VatAddress, client)
}

func LoadPot(client *ethclient.Client) (*Pot, error) {
	return NewPot(PotAddress, client)
}

func LoadPotCaller(client *ethclient.Client) (*PotCaller, error) {
	return NewPotCaller(PotAddress, client)
}

func LoadJug(client *ethclient.Client) (*Jug, error) {
	return NewJug(JugAddress, client)
}

func LoadJugCaller(client *ethclient.Client) (*JugCaller, error) {
	return NewJugCaller(JugAddress, client)
}

func LoadFlipperEthA(client *ethclient.Client) (*Flipper, error) {
	return NewFlipper(FlipEthAAddress, client)
}

func LoadFlipperEthACaller(client *ethclient.Client) (*FlipperCaller, error) {
	return NewFlipperCaller(FlipEthAAddress, client)
}
