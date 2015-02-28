SynthDef(\SineTone, {
    Out.ar(0, SinOsc.ar(440));
}).writeDefFile(File.getcwd);
0.exit;
