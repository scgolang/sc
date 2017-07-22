SynthDef(\Beats, {
    var lfo = SinOsc.kr(0.2, add: 440);
    Out.ar(0, SinOsc.ar(lfo), SinOsc.ar(lfo));
}).writeDefFile(File.getcwd);

SynthDef(\foo, {
    Out.ar(0, SinOsc.ar() * Blip.ar());
}).writeDefFile(File.getcwd);

SynthDef(\bar, {
    Out.ar(0, SinOsc.ar(mul: Blip.ar()));
}).writeDefFile(File.getcwd);

SynthDef(\baz, {
    Out.ar(0, Blip.ar(mul: SinOsc.ar()));
}).writeDefFile(File.getcwd);

SynthDef(\sub, {
    Out.ar(0, SinOsc.ar() - Blip.ar());
}).writeDefFile(File.getcwd);

SynthDef(\negExample, {
    Out.ar(0, LFNoise1.ar(1500).neg);
}).writeDefFile(File.getcwd);

SynthDef(\Envgen1, {
    Out.ar(0, PinkNoise.ar() * EnvGen.kr(Env.perc, doneAction: 2));
}).writeDefFile(File.getcwd);

SynthDef(\defWith2Params, {
    arg freq=440, gain=0.5;
    var env = EnvGen.kr(Env.perc, doneAction: 2, levelScale: gain);
    var sine = SinOsc.ar(freq);
    Out.ar(0, sine * env);
}).writeDefFile(File.getcwd);

SynthDef(\SameSame, {
    var s = SinOsc.ar(220);
    Out.ar(0, [s, s]);
}).writeDefFile(File.getcwd);

SynthDef(\SawTone1, {
    arg freq=440, cutoff=1200, q=0.5;
    Out.ar(0, RLPF.ar(Saw.ar(freq), cutoff, q));
}).writeDefFile(File.getcwd);

SynthDef(\SineTone, {
    Out.ar(0, SinOsc.ar(440));
}).writeDefFile(File.getcwd);

SynthDef(\SineTone2, {
    Out.ar(0, SinOsc.ar(440, SinOsc.ar(0.1), 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\SineTone3, {
    Out.ar(0, SinOsc.ar(440, SinOsc.ar(0.1), add: 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\SineTone4, {
    arg freq=440;
    Out.ar(0, SinOsc.ar(freq));
}).writeDefFile(File.getcwd);

SynthDef(\UseParam, {
	arg freq=200;
	Out.ar(0, SinOsc.ar(freq + 20));
}).writeDefFile(File.getcwd);

SynthDef(\SimpleMulti, {
    var sine = SinOsc.ar([440, 441]);
    Out.ar(0, sine);
}).writeDefFile(File.getcwd);

SynthDef(\AllpassExample, {
    Out.ar(0, AllpassC.ar(Decay.ar(Dust.ar(1,0.5), 0.2, WhiteNoise.ar), 0.2, 0.2, 3));
}).writeDefFile(File.getcwd);

SynthDef(\BAllPassExample, {
    var sig = Saw.ar();
    sig = BAllPass.ar(sig, MouseX.kr(10, 18000, \exponential), 0.8);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\IntegratorExample, {
    Out.ar(0, Integrator.ar(LFPulse.ar(1500 / 4, 0.2, 0.1), MouseX.kr(0.01, 0.999, 1)));
}).writeDefFile(File.getcwd);

SynthDef(\FSinOscExample, {
    Out.ar(0, FSinOsc.ar(FSinOsc.ar(XLine.kr(4, 401, 8), 0.0, 200, 800)) * 0.2);
}).writeDefFile(File.getcwd);

SynthDef(\BPFExample, {
    var line = XLine.kr(0.7, 300, 20);
    var saw = Saw.ar(200, 0.5);
    var sine = FSinOsc.kr(line, 0, 3600, 4000);
    Out.ar(0, BPF.ar(saw, sine, 0.3));
}).writeDefFile(File.getcwd);

SynthDef(\BRFExample, {
    var line = XLine.kr(0.7, 300, 20);
    var saw = Saw.ar(200, 0.5);
    var sine = FSinOsc.kr(line, 0, 3800, 4000);
    Out.ar(0, BRF.ar(saw, sine, 0.3));
}).writeDefFile(File.getcwd);

SynthDef(\Balance2Test, {
    var l = LFSaw.ar(44);
    var r = Pulse.ar(33);
    var pos = FSinOsc.kr(0.5);
    var level = 0.1;
    Out.ar(0, Balance2.ar(l, r, pos, level));
}).writeDefFile(File.getcwd);

SynthDef(\BlipExample, {
    var freq = XLine.kr(20000, 200, 6);
    var harms = 100;
    var mul = 0.2;
    Out.ar(0, Blip.ar(freq, harms, mul));
}).writeDefFile(File.getcwd);

SynthDef(\BrownNoiseTest, {
    var sig = SinOsc.ar(BrownNoise.ar(100, 200));
    Out.ar(0, sig * 0.1);
}).writeDefFile(File.getcwd);

SynthDef(\CascadeExample, {
    var mod1 = SinOsc.ar([440, 441]);
    var mod2 = SinOsc.ar(mod1);
    Out.ar(0, SinOsc.ar(mod2));
}).writeDefFile(File.getcwd);

SynthDef(\EnvgenTest, {
    Out.ar(0, PinkNoise.ar() * EnvGen.kr(Env.perc, doneAction: 2));
}).writeDefFile(File.getcwd);

SynthDef(\LFSawExample, {
    var freq = LFSaw.kr(4, 0, 200, 400);
    Out.ar(0, LFSaw.ar(freq, 0, 0.1));
}).writeDefFile(File.getcwd);

SynthDef(\LFPulseTest, {
    var freq = LFPulse.kr(3, 0, 0.3, 200, 200);
    Out.ar(0, LFPulse.ar(freq, 0, 0.2, 0.1));
}).writeDefFile(File.getcwd);

SynthDef(\ImpulseExample, {
    var freq = XLine.kr(800, 100, 5);
    var gain = 0.5;
    var phase = 0.0;
    var sig = Impulse.ar(freq, phase, gain);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\LFNoise1Example, {
    var freq = XLine.kr(1000, 10000, 10);
    Out.ar(0, LFNoise1.ar(freq, 0.25));
}).writeDefFile(File.getcwd);

SynthDef(\LFTriExample, {
    var freq = LFTri.kr(4, 0, 200, 400);
    Out.ar(0, LFTri.ar(freq, 0, 0.1));
}).writeDefFile(File.getcwd);

SynthDef(\PlayBufExample, {
    arg bufnum = 0;
    Out.ar(0, PlayBuf.ar(1, bufnum, 1.0, 1.0, 0, 0, 2));
}).writeDefFile(File.getcwd);

SynthDef(\TrigTest, {
        Out.ar(0, Trig.ar(Dust.ar(1), 0.2) * FSinOsc.ar(800, 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\Trig1Test, {
        Out.ar(0, Trig1.ar(Dust.ar(1), 0.2) * FSinOsc.ar(800, 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\VarSawTest, {
        Out.ar(0, VarSaw.ar(LFPulse.kr(3, 0, 0.3, 200, 200), 0, 0.2, 0.1));
}).writeDefFile(File.getcwd);

SynthDef(\OscTest, {
    arg bufnum = 0;
    Out.ar(0, Osc.ar(bufnum, XLine.kr(2000,200), 0, 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\OscNTest, {
    arg bufnum = 0;
    Out.ar(0, OscN.ar(bufnum, XLine.kr(2000,200), 0, 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\VOscTest, {
    arg bufnum = 0;
    Out.ar(0, VOsc.ar(bufnum, XLine.kr(2000,200), 0, 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\VOsc3Test, {
    arg bufnum = 0;
    var line1 = XLine.kr(2000, 200, 0.5);
    var line2 = XLine.kr(2000, 200, 1.5);
    var line3 = XLine.kr(2000, 200, 4.5);
    Out.ar(0, VOsc3.ar(bufnum, line1, line2, line3));
}).writeDefFile(File.getcwd);

SynthDef(\PSinGrainTest, {
    Out.ar(0, PSinGrain.ar(880, 0.1, 0.7));
}).writeDefFile(File.getcwd);

SynthDef(\ShaperTest, {
    arg bufnum = 0;
    Out.ar(0, Shaper.ar(bufnum, SinOsc.ar(440, 0.5, Line.kr(0,0.9,6))));
}).writeDefFile(File.getcwd);

SynthDef(\SinOscFBTest, {
	Out.ar(0, SinOscFB.ar(100*SinOscFB.ar(MouseY.kr(1,1000,'exponential'))+200,MouseX.kr(0.5pi,pi))*0.1);
}).writeDefFile(File.getcwd);

SynthDef(\RLPFTest, {
    Out.ar(0, RLPF.ar(Saw.ar(200, 0.1), FSinOsc.kr(XLine.kr(0.7, 300, 20), 0, 3600, 4000), 0.2));
}).writeDefFile(File.getcwd);

SynthDef(\BallTest, {
	var sf = LFNoise0.ar(MouseX.kr(1, 100, 1));
	var g  = MouseY.kr(0.1, 10, 1);
	var f  = Ball.ar(sf, g, 0.01, 0.01);
	f = f * 140 + 500;
	Out.ar(0, SinOsc.ar(f, 0, 0.2));
}).writeDefFile(File.getcwd);

SynthDef(\SlewTest, {
    Out.ar(0, Slew.ar(Saw.ar(800, mul: 0.2), 400, 400));
}).writeDefFile(File.getcwd);

SynthDef(\XFade2Test, {
    Out.ar(0, XFade2.ar( Saw.ar, SinOsc.ar , LFTri.kr(0.1) ));
}).writeDefFile(File.getcwd);

SynthDef(\LinXFade2Test, {
    Out.ar(0, LinXFade2.ar( Saw.ar, SinOsc.ar , LFTri.kr(0.1), 1 ));
}).writeDefFile(File.getcwd);

SynthDef(\CrackleTest, {
    var crack = Crackle.ar(Line.kr(1.0, 2.0, 3), 0.5, 0.5);
    Out.ar(0, crack);
}).writeDefFile(File.getcwd);

SynthDef(\GrainBufTest, {
    Out.ar(0, GrainBuf.ar(numChannels: 1, sndbuf: 0));
}).writeDefFile(File.getcwd);

SynthDef(\COscTest, {
    Out.ar(0, COsc.ar(0, 200, 0.7, 0.25));
}).writeDefFile(File.getcwd);

SynthDef(\ClipNoiseTest, {
    Out.ar(0, ClipNoise.ar(0.2));
}).writeDefFile(File.getcwd);

SynthDef(\CombCTest, {
    var line = XLine.kr(0.0001, 0.01, 20);
    var sig = CombC.ar(WhiteNoise.ar(0.01), 0.01, line, 0.2);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\CombLTest, {
    var line = XLine.kr(0.0001, 0.01, 20);
    var sig = CombL.ar(WhiteNoise.ar(0.01), 0.01, line, 0.2);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\CombNTest, {
    var line = XLine.kr(0.0001, 0.01, 20);
    var sig = CombN.ar(WhiteNoise.ar(0.01), 0.01, line, 0.2);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\Decay2Test, {
    var line = XLine.kr(1, 50, 20);
    var pulse = Impulse.ar(line, 0.25);
    Out.ar(0, Decay2.ar(pulse, 0.01, 0.2, FSinOsc.ar(600)));
}).writeDefFile(File.getcwd);

SynthDef(\DelayCTest, {
    var z = Decay.ar(Dust.ar(1, 0.5), 0.3, WhiteNoise.ar());
    Out.ar(0, DelayC.ar(z, 0.2, 0.2, 1, z));
}).writeDefFile(File.getcwd);

SynthDef(\DelayLTest, {
    var z = Decay.ar(Dust.ar(1, 0.5), 0.3, WhiteNoise.ar());
    Out.ar(0, DelayL.ar(z, 0.2, 0.2, 1, z));
}).writeDefFile(File.getcwd);

SynthDef(\DelayNTest, {
    var z = Decay.ar(Dust.ar(1, 0.5), 0.3, WhiteNoise.ar());
    Out.ar(0, DelayN.ar(z, 0.2, 0.2, 1, z));
}).writeDefFile(File.getcwd);

SynthDef(\DustTest, {
    Out.ar(0, Dust.ar(XLine.kr(20000, 2, 10), 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\Dust2Test, {
    Out.ar(0, Dust2.ar(XLine.kr(20000, 2, 10), 0.5));
}).writeDefFile(File.getcwd);

SynthDef(\PulseTest, {
    Out.ar(0, Pulse.ar(XLine.kr(40,4000,6),0.1, 0.2));
}).writeDefFile(File.getcwd);

SynthDef(\SyncSawTest, {
    Out.ar(0, SyncSaw.ar(800, Line.kr(800, 1600, 0.01)));
}).writeDefFile(File.getcwd);

SynthDef(\FormletTest, {
    var in = Blip.ar(SinOsc.kr(5, 0, 20, 300), 1000, 0.1);
    Out.ar(0, Formlet.ar(in, XLine.kr(1500, 700, 8), 0.005, 0.4));
}).writeDefFile(File.getcwd);

SynthDef(\FreeVerbTest, {
    arg mix=0.25, room=0.15, damp=0.5;
    var decay = Decay.ar(Impulse.ar(1), 0.25, LFCub.ar(1200, 0, 0.1));
    var sig = FreeVerb.ar(decay, mix, room, damp);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\LFCubTest, {
    var freq = LFCub.kr(LFCub.kr(0.2, 0, 8, 10), 0, 400, 800);
    var sig = LFCub.ar(freq, 0, 0.1);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\LFGaussTest, {
    Out.ar(0, LFGauss.ar(0.01, SampleDur.ir * MouseX.kr(10, 3000, 1)) * 0.2);
}).writeDefFile(File.getcwd);

SynthDef(\GateTest, {
    var noise = WhiteNoise.kr(1, 0);
    var pulse = LFPulse.kr(1.333, 0.5);
    Out.ar(0, Gate.ar(noise, pulse));
}).writeDefFile(File.getcwd);

SynthDef(\GrainFMTest, {
    arg gate=1, amp=1;
    var pan = MouseX.kr(-1, 1);
    var freqdev = WhiteNoise.kr(MouseY.kr(0, 400));
    var env = Env([0, 1, 0], [1, 1], \sin, 1);
    var ampenv = EnvGen.kr(env, gate, levelScale: amp, doneAction: 2);
    var trig = Impulse.kr(10);
    var modIndex = LFNoise1.kr.range(1, 10);
    var sig = GrainFM.ar(
        numChannels: 2,
        trigger: trig,
        dur: 0.1,
        carfreq: 440+freqdev,
        modfreq: 200,
        index: modIndex,
        pan: pan
    );
    Out.ar(0, sig * ampenv);
}).writeDefFile(File.getcwd);

SynthDef(\PMOscTest, {
    Out.ar(0, PMOsc.ar(Line.kr(600, 900, 5), 600, 3, 0, 0.1));
}).writeDefFile(File.getcwd);

//
//                     Out
//                      |
//                   +-----+
//                   |     |
//                   0   BinaryOpUgen
//                         |
//                    +---------+
//                    |         |
//                 Select      0.2
//                    |
//        +---------------------+--------+-------+
//        |                     |        |       |
//      MulAdd                SinOsc    Saw    Pulse
//        |                     |        |       |
//   +-------+-------+       +-----+    440   +-----+
//   |       |       |       |     |          |     |
// LFSaw    1.5     1.5     440    0         440   0.5
//
SynthDef(\SelectTest, {
    var a, cycle;
    a = [ SinOsc.ar, Saw.ar, Pulse.ar ];
    cycle = a.size  * 0.5;
    Out.ar(0, Select.ar(LFSaw.kr(1.0, 0.0, cycle, cycle), a) * 0.2);
}).writeDefFile(File.getcwd);

SynthDef(\Sum3Test, {
    Out.ar(0, Sum3.new(PinkNoise.ar(0.1), FSinOsc.ar(801, 0.1), LFSaw.ar(40, 0.1)));
}).writeDefFile(File.getcwd);

SynthDef(\MixTest, {
    Out.ar(0, Mix.new([
        PinkNoise.ar(0.1),
	FSinOsc.ar(801, 0.1),
	LFSaw.ar(40, 0.1),
	Pulse.ar(436.0),
	Dust.ar(4.0)
    ]));
}).writeDefFile(File.getcwd);

SynthDef(\TestEnvADSR, {
    Out.ar(0, SinOsc.ar() * EnvGen.kr(Env.adsr(), doneAction: 2));
}).writeDefFile(File.getcwd);

SynthDef(\THX, {
    var numVoices = 30;
    var fundamentals = ({rrand(200.0, 400.0)}!numVoices).sort;
    var finalPitches = (numVoices.collect({|nv| (nv/(numVoices/6)).round * 12; }) + 14.5).midicps;
    var sweepEnv = EnvGen.kr(Env([0, 0.1, 1], [5, 8], [2, 5]));
    var sig = Mix({ |numTone|
	var initRandomFreq = fundamentals[numTone] + LFNoise2.kr(0.5, 3 * (numTone + 1));
	var destinationFreq = finalPitches[numTone] + LFNoise2.kr(0.1, (numTone / 4));
	var freq = ((1 - sweepEnv) * initRandomFreq) + (sweepEnv * destinationFreq);
	Pan2.ar(
		BLowPass.ar(Saw.ar(freq), freq * 8, 0.5),
		rrand(-0.5, 0.5),
		numVoices.reciprocal
	)
    }!numVoices);
    Out.ar(0, sig);
}).writeDefFile(File.getcwd);

SynthDef(\DetectSilence, { arg out;
    var z;
    z = SinOsc.ar(Rand(400, 700), 0, LFDNoise3.kr(8).max(0)).softclip * 0.3;
    DetectSilence.ar(z, doneAction:2);
    Out.ar(out, z);
}).writeDefFile(File.getcwd);

SynthDef(\FittonBubbles, {
    var saw2 = LFSaw.ar([8,7.23],0,3,80);
    var saw1 = LFSaw.ar(0.4, 0, 24, saw2).midicps;
    
    Out.ar(0, CombC.ar(
        SinOsc.ar(saw1, 0, 0.04),
        0.2,
        0.2,
        4
    ));
}).writeDefFile(File.getcwd);

SynthDef(\InTest, { arg out=0, in=0;
    Out.ar(out, SinOsc.ar(In.kr(in, 2), 0, 0.1));
}).writeDefFile(File.getcwd);


SynthDef(\TGrainsExample, {
	Out.ar(0, TGrains.ar(2, Impulse.ar(4)));
}).writeDefFile(File.getcwd);

SynthDef(\SoundInTest0, {
    Out.ar(0, SoundIn.ar(0));
}).writeDefFile(File.getcwd);

SynthDef(\SoundInTest00, {
    Out.ar(0, SoundIn.ar([0, 0]));
}).writeDefFile(File.getcwd);

SynthDef(\SoundInTest01, {
    Out.ar(0, SoundIn.ar([0, 1]));
}).writeDefFile(File.getcwd);

SynthDef(\SoundInTest02, {
    Out.ar(0, SoundIn.ar([0, 2]));
}).writeDefFile(File.getcwd);

SynthDef(\SoundInTest12, {
    Out.ar(0, SoundIn.ar([1, 2]));
}).writeDefFile(File.getcwd);

SynthDef(\SoundInTest20, {
    Out.ar(0, SoundIn.ar([2, 0]));
}).writeDefFile(File.getcwd);

SynthDef(\Warp1Example, {
	Out.ar(0, Warp1.ar(2));
}).writeDefFile(File.getcwd);

SynthDef(\GVerbExample, {
	Out.ar(0, GVerb.ar(SinOsc.ar(220)));
}).writeDefFile(File.getcwd);

SynthDef(\LPFExample, {
	Out.ar(0, LPF.ar(SinOsc.ar(220)));
}).writeDefFile(File.getcwd);

SynthDef(\HPFExample, {
	Out.ar(0, HPF.ar(SinOsc.ar(220)));
}).writeDefFile(File.getcwd);

SynthDef(\FreeVerbExample, {
	Out.ar(0, FreeVerb.ar(SinOsc.ar(220)));
}).writeDefFile(File.getcwd);

SynthDef(\GrainFMExample, {
	Out.ar(0, GrainFM.ar());
}).writeDefFile(File.getcwd);

SynthDef(\TDelayTest, {
        z = Impulse.ar(2);
        Out.ar(0, [z * 0.1, ToggleFF.ar(TDelay.ar(z, 0.5)) * SinOsc.ar(mul: 0.1)]);
}).writeDefFile(File.getcwd);

SynthDef(\DCTest, {
	Out.ar(0, DC.ar());
}).writeDefFile(File.getcwd);

SynthDef(\LeakDCTest, {
	Out.ar(0, LeakDC.ar(LFPulse.ar(800, 0.5, 0.5, 0.5)));
}).writeDefFile(File.getcwd);

SynthDef(\SweepTest, {
	Out.ar(0, LFPulse.ar() * Sweep.ar());
}).writeDefFile(File.getcwd);

SynthDef(\HasherTest, {
	Out.ar(0, SinOsc.ar(Hasher.kr(MouseX.kr(0, 10), 300, 500)));
}).writeDefFile(File.getcwd);

SynthDef(\LinPan2Test, {
	Out.ar(0, LinPan2.ar(FSinOsc.ar(800, 0, 0.1), FSinOsc.kr(3)))
}).writeDefFile(File.getcwd);

SynthDef(\LagTest, {
	Out.ar(0, SinOsc.ar(
                Lag.kr(
                        LFPulse.kr(4, 0, 0.5, 50, 400),
                        Line.kr(0, 1, 15)
                ),
                0,
                0.3
        ));
}).writeDefFile(File.getcwd);

SynthDef(\LatchTest, {
	Out.ar(0, Blip.ar(Latch.ar(WhiteNoise.ar, Impulse.ar(9)) * 400 + 500, 4, 0.2));
}).writeDefFile(File.getcwd);

SynthDef(\LineTest, {
	Out.ar(0, SinOsc.ar(Line.kr(200, 17000, 10), 0, 0.1));
}).writeDefFile(File.getcwd);

SynthDef(\XLineTest, {
	Out.ar(0, SinOsc.ar(XLine.kr(200, 17000, 10), 0, 0.1));
}).writeDefFile(File.getcwd);

SynthDef(\OnePoleTest, {
	Out.ar(0, OnePole.ar(WhiteNoise.ar(0.5), Line.kr(-0.99, 0.99, 10)));
}).writeDefFile(File.getcwd);

SynthDef(\OneZeroTest, {
	Out.ar(0, OneZero.ar(WhiteNoise.ar(0.5), Line.kr(-0.5, 0.5, 10)));
}).writeDefFile(File.getcwd);

SynthDef(\Pan2Test, {
	Out.ar(0, Pan2.ar(PinkNoise.ar(0.4), FSinOsc.kr(2), 0.3));
}).writeDefFile(File.getcwd);

SynthDef(\PanAzTest, {
	Out.ar(0, PanAz.ar(2, DC.ar(1), Line.ar(0, 1/2, 0.1)));
}).writeDefFile(File.getcwd);

SynthDef(\PanB2Test, {
	var w, x, y, p, a, b, c, d;
        p = PinkNoise.ar; // source
        // B-format encode
        #w, x, y = PanB2.ar(p, MouseX.kr(-1,1), 0.1);
        // B-format decode to quad
        #a, b, c, d = DecodeB2.ar(4, w, x, y);
        Out.ar(0, [a, b, d, c]); // reorder to my speaker arrangement: Lf Rf Lr Rr
}).writeDefFile(File.getcwd);

SynthDef(\Pan4Test, {
	Out.ar(0, Pan4.ar(PinkNoise.ar, FSinOsc.kr(2), FSinOsc.kr(1.2), 0.3))
}).writeDefFile(File.getcwd);

SynthDef(\FormantTest, {
	Out.ar(0, Formant.ar(XLine.kr(400,1000, 8), 2000, 800, 0.125));
}).writeDefFile(File.getcwd);

SynthDef(\LFParTest, {
	Out.ar(0, LFPar.ar(XLine.kr(100,8000,30),0,0.1));
}).writeDefFile(File.getcwd);

SynthDef(\MedianTest, {
	Out.ar(0, LeakDC.ar(Median.ar(31, WhiteNoise.ar(0.1) + SinOsc.ar(800,0,0.1)), 0.9));
}).writeDefFile(File.getcwd);

SynthDef(\ResonzTest, {
	Out.ar(0, Resonz.ar(WhiteNoise.ar(0.5), XLine.kr(1000,8000,10), 0.05));
}).writeDefFile(File.getcwd);

SynthDef(\RingzTest, {
	Out.ar(0, Ringz.ar(Impulse.ar(6, 0, 0.3), 2000, XLine.kr(4, 0.04, 8)));
}).writeDefFile(File.getcwd);

SynthDef(\KlangTest, {
        var sig = Pan2.ar(
	        Klang.ar(`[[
                        561.384644,
                        1043.168701,
                        237.107315,
                        303.264008,
                        927.150208,
                        833.526123,
                        509.927826,
                        946.380005,
                        752.409973,
                        525.558716,
                        1111.182129,
                        715.820068
		], nil, nil], 1, 0),
		0.307131
	);
        Out.ar(0, sig * EnvGen.kr(Env.sine(4), 1, 0.02, doneAction: 2));
}).writeDefFile(File.getcwd);

SynthDef(\KlankTest1, {
        Out.ar(0, Klank.ar(`[[800, 1071, 1353, 1723], nil, [1, 1, 1, 1]], PinkNoise.ar(0.007)));
}).writeDefFile(File.getcwd);

SynthDef(\GrainInTest, {
        arg gate = 1, amp = 1, envbuf;
        var pan, env;
        // use mouse x to control panning
        pan = MouseX.kr(-1, 1);
        env = EnvGen.kr(
            Env([0, 1, 0], [1, 1], \sin, 1),
            gate,
            levelScale: amp,
            doneAction: 2);
        Out.ar(0,
            GrainIn.ar(2, Impulse.kr(32), 1, PinkNoise.ar * 0.05, pan, envbuf) * env)
}).writeDefFile(File.getcwd);

SynthDef(\GrainSinTest, {
        arg gate = 1, amp = 1, envbuf;
        var freqdev, pan, env;
        // use mouse x to control panning
	freqdev = WhiteNoise.kr(MouseY.kr(0, 400));
        pan = MouseX.kr(-1, 1);
        env = EnvGen.kr(
                Env([0, 1, 0], [1, 1], \sin, 1),
                gate,
                levelScale: amp,
                doneAction: 2
        );
	Out.ar(0, GrainSin.ar(2, Impulse.kr(10), 0.1, 440 + freqdev, pan, envbuf) * env);
}).writeDefFile(File.getcwd);

SynthDef(\FFTTest, {
        var in, chain;
        in = WhiteNoise.ar(0.2);
        chain = FFT(LocalBuf(2048), in);
        chain = PV_BrickWall(chain, SinOsc.kr(0.1));
        Out.ar(0, IFFT(chain));
}).writeDefFile(File.getcwd);

SynthDef(\PulseDividerTest, { arg out = 0;
    var p, a, b;
    p = Impulse.ar(8);
    a = SinOsc.ar(1200, 0, Decay2.ar(p, 0.005, 0.1));
    b = SinOsc.ar(600,  0, Decay2.ar(PulseDivider.ar(p, 4), 0.005, 0.5));
    Out.ar(out, (a + b) * 0.4)
}).writeDefFile(File.getcwd);

SynthDef(\PulseCountTest, {
        Out.ar(0, SinOsc.ar(PulseCount.ar(Impulse.ar(10), Impulse.ar(0.4)) * 200, 0, 0.05));
}).writeDefFile(File.getcwd);

0.exit;
