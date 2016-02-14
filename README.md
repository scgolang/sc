# sc

Package sc provides a way to interface directly with the
supercollider server from Go programs.

[![Go Report Card](https://goreportcard.com/badge/github.com/scgolang/sc)](https://goreportcard.com/report/github.com/scgolang/sc)

## Why

Supercollider has a well-deserved reputation for being difficult
to learn. I believe that sclang is a big reason for this.
I also believe that Go is quite easy to learn, hence
replacing sclang with Go makes a lot of sense to me.

## Who

See https://github.com/scgolang/sc/blob/master/AUTHORS

## What this is not

sc is not intended for livecoding (for more about livecoding see
http://en.wikipedia.org/wiki/Live_coding). I may decide someday
to support some livecoding-esque features, but for now the
goal is to be able to build music applications
that use supercollider as the audio engine.

## Install

First, and most importantly, you must have supercollider
installed (see http://supercollider.github.io/download.html).

Next, you must install go (see https://golang.org/dl).

Then

```
go get github.com/scgolang/sc
```

## Test

There are a handful of synthdef files that are necessary for the tests.

These synthdef files are created by sclang, so you have to

```
make test
```

instead of

```
go test
```

Also, there are tests that attempt to connect to scsynth on port 57120 (see [client_test.go](client_test.go)).

If you do not have scsynth running the test will hang, so run

```
scsynth -u 57120
```

## Usage

Find godocs at http://godoc.org/github.com/scgolang/sc.
Also, take a look at some of the example programs here: https://github.com/scgolang/examples.

The list of supported ugens is here: https://github.com/scgolang/sc/blob/master/UGENS.md

## Develop

See https://github.com/scgolang/sc/blob/master/CONTRIBUTING and
https://github.com/scgolang/sc/blob/master/HACKING.md.

## Roadmap

See https://github.com/scgolang/sc/milestones

## Thanks

This project is hugely indebted to

* The supercollider developer community
* Overtone (https://github.com/overtone/overtone)
