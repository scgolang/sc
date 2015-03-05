# Get the path to sclang
PLATFORM := $(shell uname -s)

ifeq ($(PLATFORM),Darwin)
SCLANG=/Applications/SuperCollider/SuperCollider.app/Contents//Resources/sclang
endif

ifeq ($(PLATFORM),Linux)
SCLANG=/usr/bin/sclang
endif

# Synthdefs needed for testing
SYNTHDEFS := SineTone            \
             SineTone2           \
             SineTone3           \
             SineTone4           \
             SawTone1            \
             Beats

SYNTHDEFS := $(addsuffix .scsyndef,$(SYNTHDEFS))
# SYNTHDEFS := $(addprefix synthdefs/,$(SYNTHDEFS))

SUBPKG := ugens
EXAMPLES := $(wildcard examples/*.go)
EXAMPLES_BIN := $(patsubst examples/%.go,%,$(EXAMPLES))

.PHONY: synthdefs clean test tools examples clean_bin

all:
	cd types && go install
	go install
	for pkg in $(SUBPKG); do cd $$pkg && go install && cd ..; done

%.scsyndef: synthdefs/%.sc
	sclang $<

synthdefs: $(SYNTHDEFS)

clean:
	rm -rf *~ *.scsyndef $(EXAMPLES_BIN)

test: $(SYNTHDEFS)
	go test
	for pkg in $(SUBPKG); do cd $$pkg && go test && cd ..; done

examples:
	for src in $(EXAMPLES); do go build $$src; done
