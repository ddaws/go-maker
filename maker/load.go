package maker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// https://changelog.makerdao.com/releases/mainnet/1.0.3/contracts.json

// TODO: Autogenerate this code or load it at runtime
var (
	VatAddress = common.HexToAddress("0x35D1b3F3D7966A1DFe207aa4514C12a259A0492B")
	PotAddress = common.HexToAddress("0x197E90f9FAD81970bA7976f33CbD77088E5D7cf7")
	JugAddress = common.HexToAddress("0x19c0976f590D67707E62397C87829d896Dc0f1F1")
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
