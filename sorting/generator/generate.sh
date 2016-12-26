#!/bin/bash

go run generator.go bubble_tmpl.go heap_tmpl.go insertion_tmpl.go merge_tmpl.go quick_tmpl.go selection_tmpl.go shell_tmpl.go
goimports -w $(pwd)/..