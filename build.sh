#!/bin/bash

GOROOT="$(go env GOROOT)";
ROOT="$(pwd)";

build_wasm() {

	cd "${ROOT}/app";
	env GOOS="js" GOARCH="wasm" go build -o "${ROOT}/source/public/main.wasm" "${ROOT}/app/main.go";

	if [[ "$?" == "0" ]]; then

		cp "${ROOT}/source/public/main.wasm" "${ROOT}/build/main.wasm";
		cp "${GOROOT}/lib/wasm/wasm_exec.js" "${ROOT}/source/public/wasm_exec.js";

		echo -e "- Generate WASM code: [\e[32mok\e[0m]";
		return 0

	else

		echo -e "- Generate WASM code: [\e[31mfail\e[0m]";
		return 1

	fi;

}

build_webview() {

	local cmd="$1";
	local os="$2";
	local arch="$3";
	local folder="${ROOT}/build/${os}";

	local ext="";
	local ldflags="";

	if [[ "${cmd}" == *"-debug" ]]; then
		ldflags="-s -w";
	fi;

	if [[ "${os}" == "windows" ]]; then
		ext="exe";
	fi;

	mkdir -p "${folder}";

	cd "${ROOT}/source";

	if [[ "${ext}" != "" ]]; then
		env GOOS="${os}" GOARCH="${arch}" go build -ldflags="${ldflags}" -o "${folder}/${cmd}_${os}_${arch}.${ext}" "${ROOT}/source/cmds/${cmd}/main.go";
	else
		env GOOS="${os}" GOARCH="${arch}" go build -ldflags="${ldflags}" -o "${folder}/${cmd}_${os}_${arch}" "${ROOT}/source/cmds/${cmd}/main.go";
	fi;

	if [[ "$?" == "0" ]]; then

		echo -e "- Generate native binary: ${cmd} / ${os} / ${arch} [\e[32mok\e[0m]";
		return 0

	else

		echo -e "- Generate native binary: ${cmd} / ${os} / ${arch} [\e[31mfail\e[0m]";
		return 1

	fi;

}

build_wasm;

build_webview "git-evac" "linux" "amd64";
build_webview "git-evac-debug" "linux" "amd64";

