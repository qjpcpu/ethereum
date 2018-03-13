#!/bin/bash
mkdir -p erc20
abigen --sol ./erc20.sol --pkg erc20 --out erc20/token.go
