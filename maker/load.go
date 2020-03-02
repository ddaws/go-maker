package maker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// https://changelog.makerdao.com/releases/mainnet/1.0.3/contracts.json

// TODO: Autogenerate this code or load it at runtime
var VatAddress = common.HexToAddress("0x35D1b3F3D7966A1DFe207aa4514C12a259A0492B")

func LoadVat(client *ethclient.Client) (*Vat, error) {
	return NewVat(VatAddress, client)
}

func LoadVatCaller(client *ethclient.Client) (*VatCaller, error) {
	return NewVatCaller(VatAddress, client)
}
