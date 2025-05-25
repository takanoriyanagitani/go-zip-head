#!/bin/sh

inputz="./sample.d/input.zip"

geninput(){
	echo generating input zip file...

	mkdir -p sample.d

	echo hw1 > ./sample.d/hw1.txt
	echo hw2 > ./sample.d/hw2.txt
	echo hw3 > ./sample.d/hw3.txt

	ls ./sample.d/*.txt |
		zip \
			-0 \
			-@ \
			-T \
			-v \
			-o \
			"${inputz}"
}

test -f "${inputz}" || geninput

echo printing input zip file...
unzip -lv "${inputz}"

echo
echo printing the head of the zip file...
wazero \
	run \
	-mount "${PWD}/sample.d:/guest.d:ro" \
	-env ENV_HEAD=2 \
	-env ENV_INPUT_ZIP_FILENAME=/guest.d/input.zip \
	./ziphead.wasm |
	fq \
		--raw-output \
		--value-output \
		-d zip \
		'.local_files[] | .file_name'
