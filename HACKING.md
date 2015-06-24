How to add a Ugen
-----------------
All the supported ugens are created by files in the ugens/ directory.
For a good example, look at [ugens/sinosc.go](ugens/sinosc.go).

Please add a test if you implement a ugen.
See [sinosc_test.go](sinosc_test.go) (in the sc package) for an example of a
test involving the SinOsc ugen.

Writing a test for a new ugen should involve:

1. Adding an sclang synthdef to synthdefs.sc that will
   generate a synthdef file that uses the ugen you are working on.
   A good place to go to find such a synthdef is the SuperCollider
   documentation for the ugen.

2. Translating the synthdef to Go. See [this file](translate-synthdefs.md)
   for any issues you may encounter when translating synthdefs to Go.

3. Comparing the Go synthdef to the sc synthdef. This is simplified by
   a utility function called [compareAndWrite](caw.go). This function will
   read the synthdef file created by sclang in step 1, create a synthdef
   file for your Go synthdef called `NAME.gosyndef`, compare the
   two files byte-for-byte, and fail your test if the two files differ.
