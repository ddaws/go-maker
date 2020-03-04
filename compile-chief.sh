#!/bin/bash

( cd lib/ds-chief && dapp build )
abigen --combined-json lib/ds-chief/out/dapp.sol.json \
  --pkg maker \
  --type DSChief \
  --exc DSAuth,DSAuthEvents,DSToken,DSTokenBase,DSMath,DSThing,DSStop,DSNote \
  --out tmp/ds_chief.go