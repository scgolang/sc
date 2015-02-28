PLATFORM := $(shell uname -s)

ifeq ($(PLATFORM),Darwin)
SCLANG=/Applications/SuperCollider/SuperCollider.app/Contents//Resources/sclang
endif

ifeq ($(PLATFORM),Linux)
SCLANG=/usr/bin/sclang
endif

SYNTHDEFS := SineTone            \
             SineTone2           \
             SineTone3           \
             SineTone4           \
             SawTone1            \
             Beats

SYNTHDEFS := $(addsuffix .scsyndef,$(SYNTHDEFS))

SUBPKG := ugens
PROGS := sdef status inspect
TOOLS := sdef.go status.go inspect.go
TOOLS := $(addprefix tools/,$(TOOLS))

.PHONY: synthdefs clean test tools

all:
	cd types && go install
	go install
	for pkg in $(SUBPKG); do cd $$pkg && go install && cd ..; done

%.scsyndef: %.sc
	sclang $<

synthdefs: $(SYNTHDEFS)

clean:
	rm -rf *~ *.scsyndef $(PROGS)

test:
	go test
	for pkg in $(SUBPKG); do cd $$pkg && go test && cd ..; done

tools:
	for t in $(TOOLS); do go build $$t; done
