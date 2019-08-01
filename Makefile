.DEFAULT_GOAL := help

SHELL := /bin/bash
PROJECT_DIR=$(shell pwd)

# Run all tests.
test:
	@scripts/test.sh
.PHONY: test

cmd:	
	@scripts/cmd.sh
.PHONY: cmd

web:	
	@scripts/web.sh
.PHONY: web

all:
	@echo "Nothing to do for all"
.PHONY: all

build:
	@echo "Nothing to do for build"
.PHONY: build

clean:
	@echo "Nothing to do for clean"
.PHONY: clean

help:	
	@echo
	@echo "   cmd – run the snip command line client"
	@echo "  test – run 'go test' for the entire project"
	@echo "   web – run the snip web client"
	@echo
.PHONY: help
