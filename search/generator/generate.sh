#!/bin/bash

go run generator.go binary_tmpl.go sequential_tmpl.go
goimports -w "$(pwd)/.."
