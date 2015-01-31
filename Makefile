PLATFORM := $(shell uname -s)

ifeq ($(PLATFORM),Darwin)
SCLANG=/Applications/SuperCollider/SuperCollider.app/Contents//Resources/sclang
endif

ifeq ($(PLATFORM),Linux)
SCLANG=/usr/bin/sclang
endif

SYNTHDEF_GENERATORS := sineTone.sc

.PHONY: synthdefs clean

synthdefs:
	for sd in $(SYNTHDEF_GENERATORS); do $(SCLANG) $$sd; done

clean:
	rm -rf *~ *.scsyndef
