SynthDef(\Beats, {
    var lfo = SinOsc.kr(0.2, add: 440);
    Out.ar(0, SinOsc.ar(lfo), SinOsc.ar(lfo));
}).writeDefFile(File.getcwd);

SynthDef(\Envgen1, {
    // EnvGen.kr(Env.perc, doneAction: 2); becomes
    // EnvGen.kr(1, 1, 0, 1, 2,
    //           0, 2, -99, -99, 1, 0.01, 5, -4, 0, 1, 5, -4);
    Out.ar(0, PinkNoise.ar(EnvGen.kr(Env.perc, doneAction: 2)));
}).writeDefFile(File.getcwd);

SynthDef(\SawTone1, {
    arg freq=440, cutoff=1200, q=0.5;
    Out.ar(0, RLPF.ar(Saw.ar(freq), cutoff, q));
}).writeDefFile(File.getcwd);

SynthDef(\SineTone, {
    Out.ar(0, SinOsc.ar(440));
}).writeDefFile(File.getcwd);

//
//                   Out
//                    |
//                 +-------+
//                 |       |
//                 0    BinaryOpUgen
//                         |
//                    +--------+
//                    |        |
//                  SinOsc    0.5
//                    |
//                 +-------+
//                 |       |
//                440    SinOsc
//                         |
//                      +-----+
//                      |     |
//                     0.1    0
//
// constants: [0.1, 0, 440, 0.5]
// ugens: [SinOsc(0.1), SinOsc(440), BinaryOpUgen, Out]
//
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

//                                 Out
//                                  |
//                              +-------+
//                              |       |
//                              0       AllpassC
//                                         |
//                          +--------+--------+--------+
//                          |        |        |        |
//               BinaryOpUgen      0.01     XLine     0.2
//                  |                         |
//              +--------+          +------+-------+-------+
//              |        |          |      |       |       |
//      WhiteNoise      0.1      0.0001   0.01     20      0
//
// constants: [0.1, 0.0001, 0.01, 20, 0, 0.2]
//
// ugens: [WhiteNoise, BinaryOpUgen, XLine, AllpassC, Out]
//
// should ugens get pushed onto the stack?
//                                       
SynthDef(\AllpassExample, {
    // var noise = WhiteNoise.ar(0.1);
    // var line = XLine.kr(0.0001, 0.01, 20);
    // var all = AllpassC.ar(noise, 0.01, line, 0.2);
    // Out.ar(0, all);
    Out.ar(0, AllpassC.ar(WhiteNoise.ar(0.1), 0.01, XLine.kr(0.0001, 0.01, 20), 0.2));
}).writeDefFile(File.getcwd);

0.exit;
