## Golang web3 example

This is a short example to demostrate the usage of interface to ethereum blockchain with golang. The example.go is generated by abigen.

#### Blockchain Commands 

Solidity contract to ABI

````bash
solc --abi contracts/Example.sol > contracts/Example.abi
````
ABI to GO package

```bash
abigen --abi contracts/Example.abi --pkg example --out example/example.go
```

#### GO commands
````
go get github.com/ethereum/go-ethereum
````
````
go get github.com/victoryeo/golang-web3/example
````
````
go run main.go
````
#### Problems
If you receive:
````
Subscribe: notifications not supported
````
This is because notifications require a connection oriented network protocol. Port 8545 by default is the HTTP RPC interface. HTTP does not support server push notifications, so subscriptions do not work on it. Port 8546 by default is the WebSocket interface, which does support notifications, but you have to use WebSocket.