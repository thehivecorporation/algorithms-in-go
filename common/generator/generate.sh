#!/bin/bash

go run generator.go common_tmpl.go
goimports -w $(pwd)/..