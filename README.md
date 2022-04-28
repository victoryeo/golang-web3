## Golang web3 example

This is a short example to demostrate the usage of web3 with golang

#### Blockchain Commands 

Solidity contract to ABI

````bash
solc --abi contracts/Example.sol > contracts/Example.abi
````
ABI to GO package

```bash
abigen --abi contracts/Example.abi --pkg example --out example.go
```

#### GO commands
go get github.com/ethereum/go-ethereum
go get github.com/victoryeo/golang-web3/example
go run main.go