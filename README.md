# Go Maker

> Golang bindings to the Multi-Colatoral Dai stablecoin system

**Note:** This library is still under development and currently only support Vat access

## Usage 

Assuming you're connected to an Ethereum RPC gateway (geth, parity, Infura, etc...) via go-ethereum

```go
// Connect to Ethereum gateway...
client := //...

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

log.Printf(`

Vat @ %s
----------
  debt: %s
  vice: %s
  line: %s

`, maker.VatAddress.Hex(), debt, vice, line)
```

For the complete example check out the `examples/` directory!

## To Do

- Autogenerate maker/load.go for easy reference and loading of deployed contracts
- Add support for the complete MCD system of contracts

## Developing 

**Note:** Compiled bindings are comitted to the repo. You can consume this package directly like any other Go library

### Tooling

- solc - Solidity compiler
- abigen - Generate language specific bindings to Solidity contracts

### Setup

To get started you will need to clone update submodules

```bash
git submodule update --init --recursive
```
