#!/bin/bash

set -ex

# Required for format check.
go get golang.org/x/tools/cmd/goimports
# Required for imports check.
go get gonum.org/v1/tools/cmd/check-imports
# Required for copyright header check.
go get gonum.org/v1/tools/cmd/check-copyright
# Required for coverage.
go get golang.org/x/tools/cmd/cover
go get github.com/mattn/goveralls
