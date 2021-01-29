#!/bin/bash

mkdir -p bin

echo "Building to bin/ folder"

# CGO required to build example relying on crypto11 and imagick dependency.
find . -name "*.go" ! -name "*_cgo.go" -print0 | CGO_ENABLED=0 xargs -0 -n1 go build -o bin/
find . -name "*_cgo.go" -print0 | CGO_ENABLED=1 CGO_CFLAGS_ALLOW='-Xpreprocessor' xargs -0 -n1 go build

mv *pkcs11*_cgo *custom_encoder*_cgo bin/
