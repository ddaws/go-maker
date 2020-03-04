#!/bin/bash

# Generate Vat contact bindings
abigen --sol=./dss/src/vat.sol --pkg=maker --type=Vat --out=maker/vat.go

# Generate Pot contract bindings
abigen --sol=./dss/src/pot.sol --pkg=maker --type=Pot  --alias=Pie=TotalPie --out=maker/pot.go
