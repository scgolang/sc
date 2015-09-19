# Get the path to sclang
PLATFORM := $(shell uname -s)

ifeq ($(PLATFORM),Darwin)
SCLANG=/Applications/SuperCollider/SuperCollider.app/Contents//Resources/sclang
endif

ifeq ($(PLATFORM),Linux)
SCLANG=/usr/bin/sclang
endif

ifeq ($(GUI),no)
SCLANG := /usr/bin/xvfb-run --server-args="-screen 0, 1280x800x24" $(SCLANG)
endif

SYNTHDEF := Beats.scsyndef

.PHONY: all clean test

all:
	go install

%.scsyndef: synthdefs.sc
	$(SCLANG) $<

clean:
	rm -rf *~ *.scsyndef *.gosyndef *.svg *.dot *.json *.xml

test: $(SYNTHDEF)
	go test
