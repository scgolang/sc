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


// NewSynthdef("sub", func(p *Params) UgenNode {
//     blip, sine := Blip{}.Rate(AR), SinOsc{}.Rate(AR)
//     Out{C(0), sine.Sub(blip)}.Rate(AR)
// })

//                 Out
//                  |
//             +---------+
//             |         |
//             0    BinaryOpUGen(sub)
//                       |
//                  +---------+
//                  |         |
//               SinOsc      Blip
//                  |         |
//               +-----+   +-----+
//               |     |   |     |
//              440    0  440   200
//
// constants: [440, 0, 200]
//
// ugens: [SinOsc, Blip, BinaryOpUGen, Out]
//
SynthDef(\sub, {
    Out.ar(0, SinOsc.ar() - Blip.ar());
}).writeDefFile(File.getcwd);

//
//                              Out
//                               |
//                          +---------+
//                          |         |
//                          0    BinaryOpUgen
//                                    |
//                               +---------+
//                               |         |
//                          PinkNoise    EnvGen
//                                         |
//           +---+---+---+---+---+---+---+---+---+----+---+---+---+---+---+---+
//           |   |   |   |   |   |   |   |   |   |    |   |   |   |   |   |   |
//           1   1   0   1   2   0   2  -99 -99  1  0.01  5  -4   0   1   5  -4
//
// constants: [1, 0, 2, -99, 0.01, 5, -4]
//
// ugens: [EnvGen, PinkNoise, BinaryOpUgen, Out]
//
SynthDef(\Envgen1, {
    // EnvGen.kr(Env.perc, doneAction: 2); becomes
    // EnvGen.kr(1, 1, 0, 1, 2,
    //           0, 2, -99, -99, 1, 0.01, 5, -4, 0, 1, 5, -4);
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

//
//                   Out
//                    |
//                 +-------+
//                 |       |
//                 0    BinaryOpUgen(mul)
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
//
// ugens: [SinOsc(0.1), SinOsc(440), BinaryOpUgen, Out]
//
SynthDef(\SineTone2, {
    Out.ar(0, SinOsc.ar(440, SinOsc.ar(0.1), 0.5));
}).writeDefFile(File.getcwd);

//
//                  Out
//                   |
//                +------+
//                |      |
//                0    BinaryOpUgen(add)
//                         |
//                    +---------+
//                    |         |
//                 SinOsc      0.5
//                   |
//               +-------+
//               |       |
//              440    SinOsc
//                       |
//                  +---------+
//                  |         |
//                 0.1        0
//
// constants: [0.1, 0, 440, 0.5]
//
// ugens: [SinOsc(0.1), SinOsc(440), BinaryOpUGen, Out]
//
SynthDef(\SineTone3, {
    Out.ar(0, SinOsc.ar(440, SinOsc.ar(0.1), add: 0.5));
}).writeDefFile(File.getcwd);

//
//                 Out
//                  |
//               +-----+
//               |     |
//               0   SinOsc
//                     |
//                 +-------+
//                 |       |
//             Control     0
//
SynthDef(\SineTone4, {
    arg freq=440;
    Out.ar(0, SinOsc.ar(freq));
}).writeDefFile(File.getcwd);

SynthDef(\UseParam, {
	arg freq=200;
	Out.ar(0, SinOsc.ar(freq + 20));
}).writeDefFile(File.getcwd);

//                        Out
//                         |
//              +----------+-----------+
//              |          |           |
//              0     SinOsc(440)  SinOsc(441)
//
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
SynthDef(\AllpassExample, {
    Out.ar(0, AllpassC.ar(WhiteNoise.ar(0.1), 0.01, XLine.kr(0.0001, 0.01, 20), 0.2));
}).writeDefFile(File.getcwd);

//                                  Out
//                                   |
//                             +-----------+
//                             |           |
//                             0        AllpassN
//                                         |
//                       +-----------+-----------+-----------+
//                       |           |           |           |
//                 BinaryOpUgen     0.2         0.2          3
//                       |
//                +-------------+
//                |             |
//              Decay      WhiteNoise
//                |
//          +-------------+
//          |             |
//     BinaryOpUGen      0.2
//          |
//    +------------+
//    |            |
//   Dust         0.5
//    |
//    1
//
// constants: [1, 0.5, 0.2, 3, 0]
//
// ugens: [Dust, BinaryOpUGen(Dust, 0.5), Decay, WhiteNoise, BinaryOpUGen(Decay, WhiteNoise), AllpassN, Out]
//
SynthDef(\AllpassnExample, {
    var noise = WhiteNoise.ar();
    var dust = Dust.ar(1, 0.5);
    var decay = Decay.ar(dust, 0.2, noise);
    var sig = AllpassN.ar(decay, 0.2, 0.2, 3);
    Out.ar(0, sig);
    // Out.ar(0, AllpassN.ar(Decay.ar(Dust.ar(1, 0.5), 0.2, WhiteNoise.ar()), 0.2, 0.2, 3));
}).writeDefFile(File.getcwd);

0.exit;
