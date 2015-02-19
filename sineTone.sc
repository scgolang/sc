SynthDef(\SineTone, {
    Out.ar(0, SinOsc.ar(440, SinOsc.ar(0.1), 0.5));
    // 
    // which becomes
    //
    // Out.ar(0, BinaryOpUGen(SinOsc.ar(440, SinOsc.ar(0.1, 0)), 0.5))
    //
    // where the BinaryOpUGen synth has a "special index" of 2
    //
    // so...
    // BinaryOpUGen special indices
    // 2 => multiplication
    //
}).writeDefFile(File.getcwd);
0.exit;
