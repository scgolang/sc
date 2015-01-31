PLATFORM := $(shell uname -s)
ifeq ($(PLATFORM),Darwin)
SCLANG=/Applications/SuperCollider/SuperCollider.app/Contents//Resources/sclang
endif
ifeq ($(PLATFORM),Linux)
SCLANG=/usr/bin/sclang
endif

testTone:
	$(SCLANG) testTone.sc
