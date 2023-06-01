#!/usr/bin/env bash

# bun
bun build --compile rpj.bun.js

# deno
deno compile --allow-read rpj.deno.js

# go
go build rpj.go
