#!/bin/bash

# TODO: Replace me with go generate invocations.
if [ -n "$(git diff)" ]; then
	exit 1
fi
