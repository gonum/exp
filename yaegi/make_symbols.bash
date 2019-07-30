#!/usr/bin/env bash

for p in $(go list gonum.org/v1/gonum/...); do
	if [[ $p = *"/internal/"* ]]; then
		continue
	fi

	if [[ $(go list -f '{{.Name}}' $p) = main ]]; then
		continue
	fi

	goexports -license LICENSE $p
done

goimports -w .
