go mod tidy

# tinygo build -o ./build/wrap.wasm -target ./wasm-memory.json ./module/wrap/main/main.go

tinygo build -o ./build-staging/module.wasm -target=wasi ./module/wrap/main/main.go

# put wrap imports first
wrap-sort-imports

wasm-snip ./build-staging/module.wasm -o ./build-staging/module_snip.wasm -p wasi_snapshot_preview1.*

wasm-opt -Os ./build-staging/module_snip.wasm -o ./build-staging/module_snip_opt.wasm

# wasm-snip -o ./build/wrap.wasm ./build/main.wasm -p syscall runtime.ticks fd_write tinygo

# rm ./build/main.wasm
