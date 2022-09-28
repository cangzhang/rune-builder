#!/usr/bin/env bash

args="${@}"

go=$(command -v go)
npm=$(command -v npm)

$go build -o ./data-maker cmd/main.go
./data-maker $args
cp output/index.json output/op.gg/
cp output/index.json output/murderbridge/

cd output/op.gg
$npm publish --access public

cd ../murderbridge
$npm publish --access public
