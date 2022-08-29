<h1 align="center">
   Qtum Qtool
</h1>

## Overview

`qtool` provides tools for handling cryptographic related objects (like *addresses, keys and scriptPubKey*) using **qtum** specific parameters.

`qtool` is both a CLI tool and a JSON-RPC server, both exposing the same functionality

The `qtool` web UI is available at [qtool.qtum.info](https://qtool.qtum.info)

---

## Tools included

### 1. `qtool-cli` 

- Command line appllication to run qtool utilities

- Available commands

  ```bash
    convertaddress     Converts a legacy address from hex encoding to base58 encoding (and vice-versa)
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

- Detailed qtool-cli *sub command* info

    [convertaddress](./qtool-cli/docs/qtool_convertaddress.md)

    [convertprivkey](./qtool-cli/docs/qtool_convertprivkey)

    [getaddfromprivkey](./qtool-cli/docs/qtool_getaddrfromprivkey.md)

    [p2pktoaddr](./qtool-cli/docs/qtool_p2pktoaddr.md)

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


### 2. `qtool-api` 
- web server that exposes `qtool` utilities via a JSON-RPC api
  
- starting `qtool-api` on a docker container (listening on port 8080)
  ```bash
  make run-api
  ```
- Available endpoints

  `/privatekey`
  ```bash
  ❯ curl -X POST -d '{"jsonrpc":"2.0","id":1,"method":"convertprivkey","params":{"data":"00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35","format":"hex","network":"testnet", "blockchain":"qtum"}}' https://qtool.qtum.info/api/privatekey
  {
    "jsonrpc": "2.0",
    "result": {
      "privKey": "cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk"
    },
    "error": null,
    "id": 1
  }
  ```
  `/address`
  ```bash
  ❯ curl -X POST -d '{"jsonrpc":"2.0","id":1,"method":"convertaddress","params":{"data":"qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW","format":"b58","network":"testnet", "blockchain":"qtum"}}' https://qtool.qtum.info/api/address
  {
    "jsonrpc": "2.0",
    "result": {
      "address": "7926223070547d2d15b2ef5e7383e541c338ffe9"
    },
    "error": null,
    "id": 1
  }
  ```
  `/script`
  ```bash
  ❯ curl -X POST -d '{"jsonrpc":"2.0","id":1,"method":"getaddressfromscriptpubKey","params":{"data":"210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112ac","network":"testnet", "blockchain":"qtum"}}' https://qtool.qtum.info/api/script
  {
    "jsonrpc": "2.0",
    "result": {
      "scriptPubKey": {
        "hex": "210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112ac",
        "asm": "0299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112 OP_CHECKSIG"
      },
      "pubKey": "0299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112",
      "addressHex": "7926223070547d2d15b2ef5e7383e541c338ffe9",
      "addressBase58": "qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW"
    },
    "error": null,
    "id": 1
  }
  ```

- API specs

(TBD)

### 3. `qtool pkg`
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

- Starting / stoppping qtool-api and react web UI in dev environment

  ```bash
  make start-compose-dev
  make stop-compose-dev
  ```

- Starting / stopping qtool-api and react web UI in prod environment

  ```bash
  make start-compose-prod
  make stop-compose-prod
  ```


