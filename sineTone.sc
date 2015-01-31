SynthDef(\SineTone, {
    Out.ar(0, SinOsc.ar());
}).writeDefFile(File.getcwd);
0.exit;
