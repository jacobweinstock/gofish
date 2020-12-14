#
# SPDX-License-Identifier: BSD-3-Clause
#

all: build test

test:
	go test -v ./...

build:
	go build

clean:
	go clean
