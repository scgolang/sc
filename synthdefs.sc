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
	Out.ar(0, SinOsc.ar([440, 441]));
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
    var line = XLine.kr(0.7,300,20);
    var saw = Saw.ar(200,0.5);
    var sine = FSinOsc.kr(line,0,3600,4000);
    Out.ar(0, BPF.ar(saw, sine, 0.3));
}).writeDefFile(File.getcwd);

0.exit;
