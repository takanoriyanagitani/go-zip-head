#!/bin/sh

tinygo \
	build \
	-o ./ziphead.wasm \
	-target=wasip1 \
	-opt=z \
	-no-debug \
	./zhead.go
