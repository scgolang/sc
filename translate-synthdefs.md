TODO: this should be a proper guide on how to translate synthdefs!

For now this is just an account of some the subtleties I have discovered
that people will have to take into account when translating
synthdefs from sclang to golang.

### BinaryOpUgen

#### sclang

There are subtle differences between the following two
synthdefs:

```supercollider
SynthDef(\foo, {
    Out.ar(0, SinOsc.ar() * Blip.ar());
});
```

and

```supercollider
SynthDef(\bar, {
    Out.ar(0, SinOsc.ar(mul: Blip.ar()));
});
```

The difference is not in how they sound (they should sound identical),
but in how sclang treats them when it sorts the ugen graph.

The ugens in `foo` will be sorted in this order: `SinOsc, Blip, BinaryOpUGen, Out`,
whereas the ugens in `bar` will be sorted in this order: `Blip, SinOsc, BinaryOpUGen, Out`.

BinaryOpUgen is a ugen that most supercollider users should not care about.
Any time you use an arithmetic operator such as `+` or `*` with at least one ugen,
a BinaryOpUGen will be created.

In both `foo` and `bar` the inputs to BinaryOpUGen are (in the order in which they appear
in the synthdef file) SinOsc and Blip.

But as you can see, when we multiply by passing Blip as an argument to SinOsc (`bar`),
Blip will appear first in the topologically sorted ugen list. This is because
sclang treats the Blip ugen _as an input_ to SinOsc (and hence it should be deeper
than SinOsc), when in fact they both become inputs to BinaryOpUGen and are at the
same depth in the ugen tree.

It is tempting to say that the `foo` synthdef is equivalent to

```supercollider
SynthDef(\baz, {
    Out.ar(0, Blip.ar(mul: SinOsc.ar()));
});
```

because sclang treats the second operand of the `*` operator as the
receiver and the first operand as an input. The ugens for the `baz` synthdef
are sorted into the same order as the `foo` synthdef: `SinOsc, Blip, BinaryOpUGen, Out`.

There is a very subtle difference between the two though. You have to inspect the
synthdef file and look at the order of the inputs to BinaryOpUGen. The BinaryOpUGen in the
`foo` synthdef effectively looks like `BinaryOpUGen.ar(SinOsc.ar(), Blip.ar())`
but the one from the `baz` synthdef looks like `BinaryOpUGen.ar(Blip.ar(), SinOsc.ar())`.
The inputs are switched! Why would this be the case?

#### golang

In sc there is only one way to multiply two ugens: with the `Mul` method.

This will sort SinOsc before Blip:

```go
NewSynthdef("bar", func(p *Params) UgenNode {
    blip := Blip{}.Rate(AR)
    bus, sig := C(0), SinOsc{}.Rate(AR).Mul(blip)
    return Out{bus, sig}.Rate(AR)
})
```

and this will sort Blip before SinOsc:

```go
NewSynthdef("baz", func(p *Params) UgenNode {
    sine := SinOsc{}.Rate(AR)
    bus, sig := C(0), Blip{}.Rate(AR).Mul(sine)
    return Out{bus, sig}.Rate(AR)
})
```

So the question is, how do we translate the `foo` synthdef?

Unfortunately the answer is that you can't.

The golang "bar" synthdef will have the inputs to BinaryOpUGen in the correct order,
but the order of the SinOsc and Blip ugens in the sorted ugens list will be wrong.

Similarly, the golang "baz" synthdef will have a properly sorted ugens list
but the order of the inputs to BinaryOpUGen will be switched.

Luckily the `+` and `*` operators are commutative, so the order of the inputs
to the resulting BinaryOpUGen does not matter.
