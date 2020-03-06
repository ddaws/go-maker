#!/bin/bash

# TODO: Update this script to be position independant

# Generate Vat contact bindings
abigen --sol=./lib/dss/src/vat.sol --pkg=maker --type=Vat --out=maker/vat.go

# Generate Pot and Jug contracts in two steps to prevent also generating bindings for the DappSys library contracts
solc --abi --bin --overwrite -o build/ lib/dss/src/pot.sol
abigen --pkg maker --type Pot \
  --abi build/Pot.abi \
  --bin build/Pot.bin \
  --alias Pie=TotalPie \
  --out maker/pot.go

solc --abi --bin --overwrite -o build/ lib/dss/src/jug.sol
abigen --pkg maker --type Jug \
  --abi build/Jug.abi \
  --bin build/Jug.bin \
  --alias vat=VatAddress,vow=VowAddress \
  --out maker/jug.go
