#!/bin/bash

GOROOT=$(go env GOROOT);
ROOT=$(pwd)/..;

env GOOS=js GOARCH=wasm go build -o "${ROOT}/source/public/main.wasm" main.go;

if [[ "$?" == "0" ]]; then

	cp "${GOROOT}/lib/wasm/wasm_exec.js" "${ROOT}/source/public/wasm_exec.js";

	cd "${ROOT}/source";

	go run "./cmds/git-evac-debug/main.go";

fi;

