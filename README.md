# Osmosis Math Wrap (POC)

The Osmosis Math wrap aims to bring the math libraries found within the Osmosis app-chain into front-end applications. It does this by compiling Golang into WebAssembly-based "wrap" using Polywrap.

The schema for the wrap can be found in the [`polywrap.graphql`](./polywrap.graphql) file. The implementation of this schema can be found in the [`./module`](./module) directory. All bindings for the schema will be generated in the `./module/wrap` directory by the `polywrap build` and `polywrap codegen` commands. 

## Pre-Requisites
* `nvm` (Node Version Manager)
* `yarn`
* `docker`

## Install
1. Use correct version of node.js  
> `nvm install && nvm use`

2. Install `polywrap` CLI  
> `yarn`

## Build
`yarn build` or `npx polywrap build` both work. This will build your Go-based wrap within a docker build image.

The resulting wrap can be found in `./build`, where you'll find:
* `wrap.wasm` - portable bytecode
* `wrap.info` - manifest & abi

## Test
`yarn test` will run JavaScript-based tests, found within the [`./module/__tests__/`](./module/__tests__/) directory.
