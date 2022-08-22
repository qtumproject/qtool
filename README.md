<h1 align="center">
   Qtum Qtool
</h1>

## Overview

`qtool` provides tools for handling crypto objects (like `addresses`, `keys` and `scriptPubKey`) using *qtum* specific parameters.

The `qtool` web UI is available at [qtool.qtum.info](https://qtool.qtum.info)

---

## Tools included

### `qtool-cli` 

- Command line appllication to run qtool utilities

- Available commands

```bash
  convertaddress     Converts a legacy address from one encoding to another
  convertprivkey     Converts the encoding of a ECDSA private key
  getaddrfromprivkey Gets a base 58 address from a given private key
  p2pktoaddr         Gets the b58 encoded address from a p2pk script

Flags:
  -b, --blockchain string   blockchain: "qtum" or "btc" (default "qtum")
  -h, --help                help for qtool
  -n, --network string      network type: "testnet" or "mainnet" (default "mainnet")
  -v, --verbose             verbose output
      --version             version for qtool
```

- More info

command: [convertaddress](./qtool-cli/docs/qtool_convertaddress.md)

command: [convertprivkey](./qtool-cli/docs/qtool_convertprivkey)

command: [getaddfromprivkey](./qtool-cli/docs/qtool_getaddrfromprivkey.md)

command: [p2pktoaddr](./qtool-cli/docs/qtool_p2pktoaddr.md)

- Usage:
  
Building `qtool-cli`
```bash
make build-cli
```
Converting `private key` format from `b58` to `hex`:
```bash
$ qtool convertprivkey cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -f b58
> Result: 00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35
```
Converting `address` format to from `b58` to `hex`
```bash
$ qtool convertaddress qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW -f b58
> Result: 7926223070547d2d15b2ef5e7383e541c338ffe9
```
Getting `address` from `private key`:
```bash
$ qtool getaddrfromprivkey cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -n testnet -b qtum -f b58
> Result: qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW
```


### `qtool-api` 
- web server that exposes all `qtool` utilities via a JSON-RPC api
  
- starting `qtool-api` on a docker container (listening on port 8080)
  ```bash
  make run-api
  ```
- Available endpoints

(TBD)

- API specs

(TBD)

### `qtool pkg`
- golang pkg that implements `qtool` utilities

---

## Project structure

```javascript
qtool
  |
  |-- qtool-api   // code base for the qtool JSON-RPC api server
  |-- qtool-cli   // code base for qtool cli
  |-- pkg     // qtool library
  |-- react-web-app   // react based web UI 
```

---

## Qtool web ui

- Running react based UI for Qtool

(TBD)

