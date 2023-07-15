# Osmosis Math Wrap (Demo)

The Osmosis Math wrap aims to bring the math libraries found within the [Osmosis app-chain](https://github.com/osmosis-labs/cosmos-sdk) into front-end applications. It does this by compiling Golang into a WebAssembly-based "wrap" using Polywrap. This wrap can then be used within: TypeScript, JavaScript, Python, Rust, Swift, Kotlin, [...](https://github.com/polywrap/client-readiness)

What functionality does this wrap contain? The schema can be found in the [`polywrap.graphql`](./polywrap.graphql) file. The schema is implemented using [Go](./module).

The Go code is aware of all schema times through bindings in the [`./module/wrap`](./module/wrap) directory. This is done automatically by the `polywrap` CLI.

## Build
`npx polywrap build` or `yarn build`. This will build your Go-based wrap within a docker build image.

The resulting wrap can be found in [`./build`](./build):
* `wrap.wasm` - portable bytecode
* `wrap.info` - manifest & abi

## Test
Tests are written in TypeScript [here](./module/__tests__/).

Run all tests:
> `yarn test`

## Setup
### Pre-Requisites
* `node.js v18+`
* `yarn`
* `docker`

### Run
> `yarn install`
