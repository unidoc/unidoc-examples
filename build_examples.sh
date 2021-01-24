#!/bin/bash

mkdir -p bin

echo "Building to bin/ folder"

#CGO required to build example relying on crypto11 and imagick dependency.
find . -name "*.go" ! -name "*pkcs11*.go" ! -name "*custom_encoder*.go" -print0 | CGO_ENABLED=0 xargs -0 -n1 go build -o bin
find . -name "*pkcs11*.go" -name "*custom_encoder*.go" -print0 | CGO_ENABLED=1 xargs -0 -n1 go build -o bin
