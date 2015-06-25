## Translating synthdefs from sclang to golang

One of the first things you must learn in order to be able to use SuperCollider
is how to write synthdefs. After reading this guide you should be pretty
comfortable with creating synthdefs with golang.

This guide assumes that you are comfortable with reading and writing
synthdefs in sclang. If you aren't I suggest reading
[this](http://doc.sccode.org/Tutorials/Getting-Started/10-SynthDefs-and-Synths.html)
and creating some synthdefs with sclang before deciding to write them in golang.

### TL;DR

Before translating this synthdef to golang

```supercollider
SynthDef(\foo, {
    Out.ar(0, SinOsc.ar(mul: Blip.ar());
});
```

You should rewrite it as

```supercollider
SynthDef(\foo, {
    Out.ar(0, SinOsc.ar() * Blip.ar());
});
```

So that the golang translation

```go
NewSynthdef("foo", func(p Params) Ugen {
    bus, blip, sine := C(0), Blip{}.Rate(AR), SinOsc{}.Rate(AR)
    return Out{bus, sine.Mul(blip)}.Rate(AR)
})
```

will produce the exact same
[synthdef data](http://doc.sccode.org/Reference/Synth-Definition-File-Format.html)
as the sclang version.



### Synthdefs

A synthdef is how a SuperCollider (scsynth) client expresses the sounds they want to create.
A sound is defined by a directed graph of ugen's (unit generators), and a ugen
is anything that generates or processes an audio signal (e.g. SinOsc generates
a sine wave, whereas RLPF is a resonant lowpass filter that processes audio streams).
See [wikipedia](https://en.wikipedia.org/wiki/Unit_generator) for more background on
where ugens come from (because the history of computer music is really exciting).

Every synthdef, no matter what language you are using to create it, must have a
name and a ugen graph. The preferred way to construct a ugen graph is with code.
In sclang this code lives inside what is called a Ugen Graph Function. The sc package
uses a type called UgenFunc (see below).
Every scsynth client must flatten the synthdef's ugen graph to a
[blob of data](http://doc.sccode.org/Reference/Synth-Definition-File-Format.html)
which is then sent to scsynth as an OSC message. Then when you tell scsynth to
create new synth nodes, you can reference your synthdef by name to get the sound
you want.

### Ugen Inputs

Every ugen that has inputs (which is most ugens) will accept either constants or
other ugens as input. For instance, SinOsc has a `freq` input which controls the
frequency of the sine wave it generates. You could pass `440` to generate
a constant-frequency sine wave, or you could pass another SinOsc ugen to
oscillate the frequency of the first.

This ability to pass in different kinds of values as inputs to ugens means that ugen
inputs must be [polymorphic](https://en.wikipedia.org/wiki/Polymorphism_(computer_science)).
sclang is a dynamically typed language similar to Javascript or Python, so it will gladly
accept any value you care to pass as a ugen input and figure out what to do with it.
golang, on the other hand, is a statically typed language which means we have to be a bit
more explicit about the types we use for ugen inputs. Specifically, this means that
ugen inputs must be an interface type.

In sc all ugen inputs are implementations of the [Input](types/input.go)
interface. This has some very important ramifications on how you write synthdefs.

First of all, unlike sclang, you can not use numeric literals as inputs to ugens.
This is because a numeric literal does not implement the Input interface. In order
to pass a constant as a ugen input you must wrap it with the [C](ugens/c.go) type.

To pass a ugen as an input to another ugen, all you have to do is create your ugen
with the `Rate(int8)` method because it returns an Input.

There is also another way to control ugens in a synthdef: synthdef parameters.
If you define a synthdef with parameters, then you can pass in values for these
parameters when you create synth nodes from this synthdef.

### Synthdef Parameters

An sclang Ugen Graph Function can be a function with any number of parameters.

In golang we use this type to create ugen graphs:

```go
type UgenFunc func(p Params) Ugen
```

Here is a simple example to see what this looks like in practice:

```go
def := NewSynthdef("sineTone", func(p Params) Ugen {
    bus := C(0)
    freq := p.Add("freq", 440)
    // use the freq param as a ugen input
    sine := SinOsc{Freq:freq}.Rate(AR)
    return Out{bus, sine}.Rate(AR)
})
```

##### Technical note

The ideal way to create synthdefs might look like this:

```go
def := NewSynthdef("sineTone", func(freq C) Ugen {
    bus := C(0)
    // use the freq param as a ugen input
    sine := SinOsc{Freq:freq}.Rate(AR)
    return Out{bus, sine}.Rate(AR)
})
```

Here there is no need for a `Params` interface, but this
is not possible with golang because it is statically typed.
`NewSynthdef` needs to know what kind of argument to expect
_at compile time_ and not all synthdefs will want to use the
same function type with this approach.



## Problems

Below are some of the subtle differences encountered when 
translating synthdefs from sclang to golang.

### BinaryOpUgen

#### sclang

Consider the following two synthdefs:

```supercollider
SynthDef(\foo, {
    Out.ar(0, SinOsc.ar() * Blip.ar());
});
```

```supercollider
SynthDef(\bar, {
    Out.ar(0, SinOsc.ar(mul: Blip.ar()));
});
```

Synth nodes created from these two synthdefs should sound identical,
but there is a difference in how sclang serializes them to send them to scsynth.

The ugens in `foo` will be sorted in this order: `SinOsc, Blip, BinaryOpUGen, Out`,
whereas the ugens in `bar` will be sorted in this order: `Blip, SinOsc, BinaryOpUGen, Out`.

*Note:* BinaryOpUgen is a ugen that most SuperCollider users should not care about.
Any time you use an arithmetic operator such as `+` or `*` or use the `mul`
parameter a BinaryOpUGen will be created.

In both `foo` and `bar` the inputs to BinaryOpUGen are (in the order in which they appear
in the synthdef file) SinOsc and Blip.

But as you can see, when we multiply by passing Blip as an input to SinOsc (in `bar`),
Blip will appear first in the sorted ugen list. This is because ugens are sorted using
a [depth first search](https://en.wikipedia.org/wiki/Depth-first_search) algorithm, which
means that when flattening a ugen graph, inputs are sorted first. The confusing thing
is that the order of the inputs to BinaryOpUGen is the same for `foo` and `bar`, yet
the these inputs appear in a different order in the ugens list.

Note also that this synthdef

```supercollider
SynthDef(\baz, {
    Out.ar(0, Blip.ar(mul: SinOsc.ar()));
});
```

will sort the ugens in the same order as the `foo` synthdef: `SinOsc, Blip, BinaryOpUGen, Out`,
but the order of the inputs to BinaryOpUGen is switched! The inputs to BinaryOpUGen
for `baz` are `Blip, SinOsc`.

So the questions now are

1. These 3 synthdefs all create the same sound, so which one do we translate to Go?

2. How do we translate to Go in a way that will allow us to generate a
   binary representation that is identical to the one generated by sclang?

Note that question 2 is important if you care about automating the comparison
of sclang-generated synthdefs with golang-generated synthdefs.

#### golang

In sc there is only one way to multiply two ugens: with the `Mul` method.

```go
NewSynthdef("bar", func(p *Params) UgenNode {
    blip := Blip{}.Rate(AR)
    bus, sig := C(0), SinOsc{}.Rate(AR).Mul(blip)
    return Out{bus, sig}.Rate(AR)
})
```

When sc flattens the ugen graph (this happens when you call NewSynthdef),
the ugen order will be `SinOsc, Blip, BinaryOpUgen, Out`,
and the inputs to BinaryOpUgen will be `SinOsc, Blip`.
This is because the Mul method returns a BinaryOpUgen
with the receiver as its first input and the argument (to Mul) as its
second input.
Notice that the order of SinOsc and Blip is reversed in the ugen list when
compared to the sclang version.
What this means is that translating sclang synthdefs that use the `mul`
ugen parameter by using the golang `Mul` method of the Input interface
results in different synthdef files.

Because of this, if you are translating an sclang synthdef that uses the `mul` ugen
parameter you should first rewrite it to use the `*` binary operator.

The golang "bar" synthdef will have the inputs to BinaryOpUGen in the correct order,
but the order of the SinOsc and Blip ugens in the sorted ugens list will be wrong.

Similarly, the golang "baz" synthdef will have a properly sorted ugens list
but the order of the inputs to BinaryOpUGen will be switched.

Luckily the `+` and `*` operators are commutative, so the order of the inputs
to the resulting BinaryOpUGen does not matter.
