# Get the path to sclang
PLATFORM = $(shell uname -s)

ifeq ($(PLATFORM),Darwin)
SCLANG = /Applications/SuperCollider.app/Contents/MacOS/sclang
endif

ifeq ($(PLATFORM),Linux)
SCLANG = /usr/bin/sclang
endif

# HACK for running sclang w/o a GUI on linux
ifeq ($(GUI),no)
SCLANG = /usr/bin/xvfb-run --server-args="-screen 0, 1280x800x24" $(SCLANG)
endif

# Default target
all:
	@go install

fixtures: .fixtures
.fixtures: testdata/synthdefs.sc
	@cd testdata && $(SCLANG) synthdefs.sc
	@touch $@

lint:
	@gometalinter

coverage:
	@rm -f cover.out cover.html
	@go test -coverprofile cover.out && go tool cover -html cover.out -o cover.html

clean:
	@rm -rf *~ *.gosyndef *.svg *.dot *.json *.xml testdata/*.scsyndef .fixtures

test: .fixtures
	@go test

.PHONY: all clean fixtures test
