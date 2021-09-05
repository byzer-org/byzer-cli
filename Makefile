export GO111MODULE=on

all: mlsql release

REVISION := $(shell git rev-parse --short HEAD 2>/dev/null)
REVISIONDATE := $(shell git log -1 --pretty=format:'%ad' --date short 2>/dev/null)
PKG := mlsql.tech/allwefantasy/mlsql-lang-cli/pkg/version
LDFLAGS = -s -w
ifneq ($(strip $(REVISION)),) # Use git clone
	LDFLAGS += -X $(PKG).revision=$(REVISION) \
		   -X $(PKG).revisionDate=$(REVISIONDATE)
endif

SHELL = /bin/sh

ifdef STATIC
	LDFLAGS += -linkmode external -extldflags '-static'
	CC = /usr/bin/musl-gcc
	export CC
endif

mlsql: Makefile cmd/*.go pkg/*/*.go
	go build -ldflags="$(LDFLAGS)"  -o mlsql ./cmd

release:
	cp ./mlsql /Users/allwefantasy/projects/mlsql/src/mlsql-lang/mlsql-app_2.4-2.1.0-SNAPSHOT/bin/

