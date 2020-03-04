package maker

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (jug *JugCaller) Vat(backend bind.ContractBackend, opts *bind.CallOpts) (*Vat, error) {
	addr, err := jug.VatAddress(opts)
	if err != nil {
		return nil, err
	}
	return NewVat(addr, backend)
}

func (jug *JugSession) Vat(backend bind.ContractBackend) (*Vat, error) {
	return jug.Contract.Vat(backend, &jug.CallOpts)
}

func (jug *JugCallerSession) Vat(backend bind.ContractBackend) (*Vat, error) {
	return jug.Contract.Vat(backend, &jug.CallOpts)
}
