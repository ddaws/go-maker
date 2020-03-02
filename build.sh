#!/bin/bash

solc --abi --bin dss/src/vat.sol --overwrite -o build
abigen --bin=./build/Vat.bin --abi=./build/Vat.abi --pkg=maker --type=Vat --out=maker/vat.go
