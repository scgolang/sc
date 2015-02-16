PLATFORM := $(shell uname -s)

ifeq ($(PLATFORM),Darwin)
SCLANG=/Applications/SuperCollider/SuperCollider.app/Contents//Resources/sclang
endif

ifeq ($(PLATFORM),Linux)
SCLANG=/usr/bin/sclang
endif

SYNTHDEF_GENERATORS := sineTone.sc
SUBPKG := ugens

.PHONY: synthdefs clean test

all:
	go install
	for pkg in $(SUBPKG); do cd $$pkg && go install && cd ..; done

synthdefs:
	for sd in $(SYNTHDEF_GENERATORS); do $(SCLANG) $$sd; done

clean:
	rm -rf *~ *.scsyndef

test:
	go test
	for pkg in $(SUBPKG); do cd $$pkg && go test && cd ..; done
