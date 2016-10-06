# Get the path to sclang
PLATFORM = $(shell uname -s)

ifeq ($(PLATFORM),Darwin)
SCLANG = /Applications/SuperCollider/SuperCollider.app/Contents//Resources/sclang
endif

ifeq ($(PLATFORM),Linux)
SCLANG = /usr/bin/sclang
endif

# HACK for running sclang w/o a GUI on linux
ifeq ($(GUI),no)
SCLANG = /usr/bin/xvfb-run --server-args="-screen 0, 1280x800x24" $(SCLANG)
endif

FIXTURES = fixtures/AllpassExample.scsyndef          \
           fixtures/AllpassnExample.scsyndef         \
           fixtures/BPFExample.scsyndef              \
           fixtures/BRFExample.scsyndef              \
           fixtures/Balance2Test.scsyndef            \
           fixtures/Beats.scsyndef                   \
           fixtures/BlipExample.scsyndef             \
           fixtures/BrownNoiseTest.scsyndef          \
           fixtures/COsc.scsyndef                    \
           fixtures/Cascade.scsyndef                 \
           fixtures/CascadeExample.scsyndef          \
           fixtures/ClipNoiseTest.scsyndef           \
           fixtures/CombCTest.scsyndef               \
           fixtures/CombLTest.scsyndef               \
           fixtures/CombNTest.scsyndef               \
           fixtures/Crackle.scsyndef                 \
           fixtures/Decay2.scsyndef                  \
           fixtures/DelayCTest.scsyndef              \
           fixtures/DelayLTest.scsyndef              \
           fixtures/DelayNTest.scsyndef              \
           fixtures/DustTest.scsyndef                \
           fixtures/Dust2Test.scsyndef               \
           fixtures/EnvgenTest.scsyndef              \
           fixtures/FSinOscExample.scsyndef          \
           fixtures/FormletTest.scsyndef             \
           fixtures/FreeVerbTest.scsyndef            \
           fixtures/GateTest.scsyndef                \
           fixtures/GrainBufTest.scsyndef            \
           fixtures/GrainFMTest.scsyndef             \
           fixtures/ImpulseExample.scsyndef          \
           fixtures/IntegratorExample.scsyndef       \
           fixtures/LFNoise1Test.scsyndef            \
           fixtures/LFPulseTest.scsyndef             \
           fixtures/LFSawExample.scsyndef            \
           fixtures/LFTriExample.scsyndef            \
           fixtures/MixTest.scsyndef                 \
           fixtures/PMOscTest.scsyndef               \
           fixtures/PlayBufExample.scsyndef          \
           fixtures/PulseTest.scsyndef               \
           fixtures/SameSame.scsyndef                \
           fixtures/SawTone1.scsyndef                \
           fixtures/SelectTest.scsyndef              \
           fixtures/SimpleMulti.scsyndef             \
           fixtures/SineTone.scsyndef                \
           fixtures/SineTone2.scsyndef               \
           fixtures/SineTone3.scsyndef               \
           fixtures/SineTone4.scsyndef               \
           fixtures/Sum3.scsyndef                    \
           fixtures/THX1.scsyndef                    \
           fixtures/UseParam.scsyndef                \
           fixtures/defWith2Params.scsyndef          \
           fixtures/foo.scsyndef                     \
           fixtures/bar.scsyndef                     \
           fixtures/baz.scsyndef                     \
           fixtures/sub.scsyndef

# Default target
all:
	@go install

$(FIXTURES): fixtures/synthdefs.sc
	@cd fixtures $(SCLANG) synthdefs.sc

lint:
	@gometalinter

coverage:
	@rm -f cover.out cover.html
	@go test -coverprofile cover.out && go tool cover -html cover.out -o cover.html

clean:
	@rm -rf *~ *.gosyndef *.svg *.dot *.json *.xml

test: $(FIXTURES)
	@go test

.PHONY: all clean test
