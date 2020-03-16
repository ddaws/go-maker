#!/bin/bash

# TODO: Update this script to be position independant

docker run \
  -v `pwd`/lib/dss/src:/src \
  -v `pwd`/out:/out \
  ethereum/solc:0.5.15 --abi --bin --overwrite -o /out/ /src/cat.sol

abigen --pkg maker --type Cat \
  --abi out/Cat.abi \
  --bin out/Cot.bin \
  --out maker/cat.go

solc --abi --bin --overwrite -o out/ lib/dss/src/dai.sol
abigen --pkg maker --type Dai \
  --abi out/Dai.abi \
  --bin out/Dai.bin \
  --out maker/dai.go

solc --abi --bin --overwrite -o out/ lib/dss/src/end.sol
abigen --pkg maker --type End \
  --abi out/End.abi \
  --bin out/End.bin \
  --out maker/end.go

solc --abi --bin --overwrite -o out/ lib/dss/src/end.sol
abigen --pkg maker --type End \
  --abi out/End.abi \
  --bin out/End.bin \
  --out maker/end.go

solc --abi --bin --overwrite -o out/ lib/dss/src/flap.sol
abigen --pkg maker --type Flapper \
  --abi out/Flapper.abi \
  --bin out/Flapper.bin \
  --alias vat=VatAddress,gem=GemAddress \
  --out maker/flap.go

solc --abi --bin --overwrite -o out/ lib/dss/src/flip.sol
abigen --pkg maker --type Flipper \
  --abi out/Flipper.abi \
  --bin out/Flipper.bin \
  --out maker/flip.go

solc --abi --bin --overwrite -o out/ lib/dss/src/flop.sol
abigen --pkg maker --type Flopper \
  --abi out/Flopper.abi \
  --bin out/Flopper.bin \
  --alias vow=VowAddress \
  --out maker/flop.go

solc --abi --bin --overwrite -o out/ lib/dss/src/join.sol
abigen --pkg maker --type Join \
  --abi out/Join.abi \
  --bin out/Join.bin \
  --out maker/join.go

solc --abi --bin --overwrite -o out/ lib/dss/src/jug.sol
abigen --pkg maker --type Jug \
  --abi out/Jug.abi \
  --bin out/Jug.bin \
  --alias vat=VatAddress,vow=VowAddress \
  --out maker/jug.go

solc --abi --bin --overwrite -o out/ lib/dss/src/pot.sol
abigen --pkg maker --type Pot \
  --abi out/Pot.abi \
  --bin out/Pot.bin \
  --alias Pie=TotalPie \
  --out maker/pot.go

solc --abi --bin --overwrite -o out/ lib/dss/src/spot.sol
abigen --pkg maker --type Spot \
  --abi out/Spot.abi \
  --bin out/Spot.bin \
  --out maker/spot.go

# Generate Vat contact bindings
abigen --sol=./lib/dss/src/vat.sol --pkg=maker --type=Vat --out=maker/vat.go

solc --abi --bin --overwrite -o out/ lib/dss/src/vow.sol
abigen --pkg maker --type Vow \
  --abi out/Vow.abi \
  --bin out/Vow.bin \
  --alias vat=VatAddress,flapper=FlapperAddress,flopper=FlopperAddress \
  --out maker/vow.go
