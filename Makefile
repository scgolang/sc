# Get the path to sclang
PLATFORM := $(shell uname -s)

ifeq ($(PLATFORM),Darwin)
SCLANG=/Applications/SuperCollider/SuperCollider.app/Contents//Resources/sclang
endif

ifeq ($(PLATFORM),Linux)
SCLANG=/usr/bin/sclang
endif

SYNTHDEF := Beats.scsyndef
SUBPKG := ugens pattern
EXAMPLES := $(wildcard examples/*.go)
EXAMPLES_BIN := $(patsubst examples/%.go,%,$(EXAMPLES))

.PHONY: synthdefs clean test tools examples clean_bin

all:
	cd types && go install
	go install
	for pkg in $(SUBPKG); do cd $$pkg && go install && cd ..; done

%.scsyndef: synthdefs.sc
	sclang $<

clean:
	rm -rf *~ $(EXAMPLES_BIN) *.scsyndef *.gosyndef *.svg *.dot

test: $(SYNTHDEF)
	go test
	for pkg in $(SUBPKG); do cd $$pkg && go test && cd ..; done

examples:
	for src in $(EXAMPLES); do go build $$src; done

graphs:
