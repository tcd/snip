# Commonly used targets (see each target for more information):
#   all: Build code.
#   test: Run tests.
#   clean: Clean up.

SHELL := /bin/bash

.PHONY: testing
test-run:
	scripts/test.sh

.PHONY: test
test: test-run

