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

FIXTURES = testdata/AllpassExample.scsyndef          \
           testdata/AllpassnExample.scsyndef         \
           testdata/BPFExample.scsyndef              \
           testdata/BRFExample.scsyndef              \
           testdata/Balance2Test.scsyndef            \
           testdata/Beats.scsyndef                   \
           testdata/BlipExample.scsyndef             \
           testdata/BrownNoiseTest.scsyndef          \
           testdata/COsc.scsyndef                    \
           testdata/Cascade.scsyndef                 \
           testdata/CascadeExample.scsyndef          \
           testdata/ClipNoiseTest.scsyndef           \
           testdata/CombCTest.scsyndef               \
           testdata/CombLTest.scsyndef               \
           testdata/CombNTest.scsyndef               \
           testdata/Crackle.scsyndef                 \
           testdata/Decay2.scsyndef                  \
           testdata/DelayCTest.scsyndef              \
           testdata/DelayLTest.scsyndef              \
           testdata/DelayNTest.scsyndef              \
           testdata/DustTest.scsyndef                \
           testdata/Dust2Test.scsyndef               \
           testdata/EnvgenTest.scsyndef              \
           testdata/FSinOscExample.scsyndef          \
           testdata/FormletTest.scsyndef             \
           testdata/FreeVerbTest.scsyndef            \
           testdata/GateTest.scsyndef                \
           testdata/GrainBufTest.scsyndef            \
           testdata/GrainFMTest.scsyndef             \
           testdata/ImpulseExample.scsyndef          \
           testdata/InTest.scsyndef                  \
           testdata/IntegratorExample.scsyndef       \
           testdata/LFNoise1Test.scsyndef            \
           testdata/LFPulseTest.scsyndef             \
           testdata/LFSawExample.scsyndef            \
           testdata/LFTriExample.scsyndef            \
           testdata/MixTest.scsyndef                 \
           testdata/PMOscTest.scsyndef               \
           testdata/PlayBufExample.scsyndef          \
           testdata/PulseTest.scsyndef               \
           testdata/SameSame.scsyndef                \
           testdata/SawTone1.scsyndef                \
           testdata/SelectTest.scsyndef              \
           testdata/SimpleMulti.scsyndef             \
           testdata/SineTone.scsyndef                \
           testdata/SineTone2.scsyndef               \
           testdata/SineTone3.scsyndef               \
           testdata/SineTone4.scsyndef               \
           testdata/Sum3.scsyndef                    \
           testdata/TestEnvADSR.scsyndef             \
           testdata/THX1.scsyndef                    \
           testdata/UseParam.scsyndef                \
           testdata/defWith2Params.scsyndef          \
           testdata/foo.scsyndef                     \
           testdata/bar.scsyndef                     \
           testdata/baz.scsyndef                     \
           testdata/sub.scsyndef

# Default target
all:
	@go install

fixtures: $(FIXTURES)

$(FIXTURES): testdata/synthdefs.sc
	@cd testdata && $(SCLANG) synthdefs.sc

lint:
	@gometalinter

coverage:
	@rm -f cover.out cover.html
	@go test -coverprofile cover.out && go tool cover -html cover.out -o cover.html

clean:
	@rm -rf *~ *.gosyndef *.svg *.dot *.json *.xml fixtures/*.scsyndef

test: $(FIXTURES)
	@go test

.PHONY: all clean fixtures test
