# Commonly used targets (see each target for more information):
#   all: Build code.
#   test: Run tests.
#   clean: Clean up.

SHELL := /bin/bash

test-run:
	scripts/test.sh

test: test-run

.PHONY: testing test
