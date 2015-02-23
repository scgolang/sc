PLATFORM := $(shell uname -s)

ifeq ($(PLATFORM),Darwin)
SCLANG=/Applications/SuperCollider/SuperCollider.app/Contents//Resources/sclang
endif

ifeq ($(PLATFORM),Linux)
SCLANG=/usr/bin/sclang
endif

SYNTHDEFS := sineTone.sc          \
             sineTone2.sc         \
             sineTone3.sc

SUBPKG := ugens
TOOLS := sdef.go
TOOLS := $(addprefix tools/,$(TOOLS))

.PHONY: synthdefs clean test tools

all:
	cd types && go install
	go install
	for pkg in $(SUBPKG); do cd $$pkg && go install && cd ..; done

synthdefs:
	for sd in $(SYNTHDEFS); do $(SCLANG) $$sd; done

clean:
	rm -rf *~ *.scsyndef $(TOOLS)

test:
	go test
	for pkg in $(SUBPKG); do cd $$pkg && go test && cd ..; done

tools:
	for t in $(TOOLS); do go build $$t; done
