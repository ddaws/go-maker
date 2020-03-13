#!/bin/bash

# TODO: Update this script to be position independant

# Generate Vat contact bindings
abigen --sol=./lib/dss/src/vat.sol --pkg=maker --type=Vat --out=maker/vat.go

# Generate Pot and Jug contracts in two steps to prevent also generating bindings for the DappSys library contracts
solc --abi --bin --overwrite -o out/ lib/dss/src/pot.sol
abigen --pkg maker --type Pot \
  --abi out/Pot.abi \
  --bin out/Pot.bin \
  --alias Pie=TotalPie \
  --out maker/pot.go

solc --abi --bin --overwrite -o out/ lib/dss/src/jug.sol
abigen --pkg maker --type Jug \
  --abi out/Jug.abi \
  --bin out/Jug.bin \
  --alias vat=VatAddress,vow=VowAddress \
  --out maker/jug.go

solc --abi --bin --overwrite -o out/ lib/dss/src/flop.sol
abigen --pkg maker --type Flopper \
  --abi out/Flopper.abi \
  --bin out/Flopper.bin \
  --alias vow=VowAddress \
  --out maker/flop.go

solc --abi --bin --overwrite -o out/ lib/dss/src/flip.sol
abigen --pkg maker --type Flipper \
  --abi out/Flipper.abi \
  --bin out/Flipper.bin \
  --out maker/flip.go
