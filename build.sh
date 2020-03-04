#!/bin/bash

# TODO: Update this script to be position independant

# Generate Vat contact bindings
abigen --sol=./dss/src/vat.sol --pkg=maker --type=Vat --out=maker/vat.go

# Generate Pot contract bindings
abigen --sol=./dss/src/pot.sol --pkg=maker --type=Pot --alias=Pie=TotalPie --out=maker/pot.go

# Generate Jug contract bindings
abigen --sol=./dss/src/jug.sol --pkg=maker --type=Jug --alias=vat=VatAddress,vow=VowAddress --out=./maker/jug.go
