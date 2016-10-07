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

SynthDef(\Cascade, {
    var mod1 = SinOsc.ar([440, 441]);
    var mod2 = SinOsc.ar(mod1);
    Out.ar(0, SinOsc.ar(mod2));
}).writeDefFile(File.getcwd);

SynthDef(\AllpassExample, {
    Out.ar(0, AllpassC.ar(WhiteNoise.ar(0.1), 0.01, XLine.kr(0.0001, 0.01, 20), 0.2));
}).writeDefFile(File.getcwd);

SynthDef(\AllpassnExample, {
    var noise = WhiteNoise.ar();
    var dust = Dust.ar(1, 0.5);
    var decay = Decay.ar(dust, 0.2, noise);
    var sig = AllpassN.ar(decay, 0.2, 0.2, 3);
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

SynthDef(\THX, {
var numVoices = 30;
var fundamentals = ({rrand(200.0, 400.0)}!numVoices).sort;
var finalPitches = (numVoices.collect({|nv| (nv/(numVoices/6)).round * 12; }) + 14.5).midicps;
var sweepEnv = EnvGen.kr(Env([0, 0.1, 1], [5, 8], [2, 5]));
var sig = Mix
({|numTone|
	var initRandomFreq = fundamentals[numTone] + LFNoise2.kr(0.5, 3 * (numTone + 1));
	var destinationFreq = finalPitches[numTone] + LFNoise2.kr(0.1, (numTone / 4));
	var freq = ((1 - sweepEnv) * initRandomFreq) + (sweepEnv * destinationFreq);
	Pan2.ar
	(
		BLowPass.ar(Saw.ar(freq), freq * 8, 0.5),
		rrand(-0.5, 0.5),
		numVoices.reciprocal
	)
}!numVoices);
Out.ar(0, sig);
}).writeDefFile(File.getcwd);

0.exit;
