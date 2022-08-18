<h1 align="center">
   Qtum Qtool
</h1>

## Overview

`qtool` provides tools for handling crypto objects (like `addresses`, `keys` and `scriptPubKey`) using *qtum* specific parameters.

---

## Tools included

`qtool-cli` : command line appllication to run qtool utilities (more info: [qtool-cli README](./qtool-api/README.md))

`qtool-api`: web server that exposes all `qtool` utilities via a JSON-RPC api (more info: [qtool-api README](./qtool-cli/README.md)) 

`pkg`: golang pkg that implements `qtool` utilities

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