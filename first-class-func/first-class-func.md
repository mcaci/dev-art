# Exploring first class functions in Go

Among the many type Go has to offer us __first class functions__ are ones that in my opinion need more love by gophers. With this thread I want to share an example of its power and usefulness.

This thread is also inspired by Dave Cheney's talk [Do not fear first class functions](http) at 2015's edition of the DotGo conference. In his talk...

## Let's start

To show how we can truly use first class functions let's start with implementing the Caesar cipher algorithm, a simple shift-based cipher where every letter of a message/string is replaced with the third next letter rounding back to 'a' if the third next letter passes 'z' (e.g. 'a' becomes 'd', 'm' becomes 'p', 'y' becomes 'b'). To decode the encoded message you just replace each letter with the third previous letter.

Here is the code for encoding with the Caesar algorithm:

```go
func normalize(in string) (out []byte) { /*...;*/ return []byte(in) }
func encWithCaesar(in string) string {
    out := normalize(in)
    for i, letter := range out {
        out[i] = letter + 3
        switch {
        case out[i] < 'a':
            out[i] += ('z' - 'a' + 1)
        case out[i] > 'z':
            out[i] -= ('z' - 'a' + 1)
        }
    }
    return string(out)
}
```

The code itself is not too difficult and to decode it is the same except replacing `letter + 3` with `letter - 3`.

For the sake of simplicity we delegate the rotation of the output letter to an external `rotate` function as follows:

```go
func rotate(b byte) byte {
    switch {
    case b < 'a':
        b += ('z' - 'a' + 1)
    case b > 'z':
        b -= ('z' - 'a' + 1)
    }
    return b
}
```

Including the decoding function the initial example becomes this one:

```go
func normalize(in string) (out []byte) { /*...;*/ return []byte(in) }
func encWithCaesar(in string) string {
    out := normalize(in)
    for i, letter := range out {
        out[i] = rotate(letter + 3)
    }
    return string(out)
}
func decWithCaesar(in string) string {
    out := normalize(in)
    for i, letter := range out {
        out[i] = rotate(letter + 3)
    }
    return string(out)
}
```

This code can be extended to include the generic shift algorithm with any specified distance.

```go
func encWithShift(in string, b byte) string {
    out := normalize(in)
    for i, letter := range out {
        out[i] = rotate(letter + b)
    }
    return string(out)
}
func decWithShift(in string, b byte) string {
    out := normalize(in)
    for i, letter := range out {
        out[i] = rotate(letter - b)
    }
    return string(out)
}
```

## From lots to one

Coming here there's a lot of repetition, but one key process steps out here: the process of encoding and decoding a string boils down to get a __`byte`__ (`letter`) from the input and processing it to get another __`byte`__ to assign to `out[i]`, a `func (byte) byte` in Go words.

Please look at this snippet below:

```go
func Transform(in string, transformer func(byte) byte) string {
    out := normalize(in)
    for i, letter := range out {
        out[i] = transformer(letter)
    }
    return string(out)
}
```

Here we don't have anymore a link between the encoding/decoding function (`transform`) and the technique used to encode/decode (_caesar_, _shift_, ...). The latter is now a __behaviour__ encapsulated by the `transformer` parameter. This means that from now on, we don't need anymore to create functions for the different operations of encoding/decoding with Ceasar and with the generic Shift ciphers, but just provide the techinque as a __behaviour__ that the `transform` function will use.

What does this concretely mean for the Caesar and generic Shift ciphers? Let's have a look at the Ceasar cipher first.

If we use take the behaviours from the Caesar encoder and decoder we can write the following two functions:

```go
func WithCaesarEncoder(b byte) byte { return rotate(b + 3) }
func WithCaesarDecoder(b byte) byte { return rotate(b - 3) }
```

And this means that anyone that wishes to use the Ceasar cipher would now call:

```go
Transform("messagetoencode", WithCaesarEncoder)
Transform("encodedmessage", WithCaesarDecoder)
```

Apart from the fact that they almost read as plain English, you can see that we successfully passed the technique to encode/decode with the Caesar cipher as the __behaviour__ that the `transform` function will use.

Let's shift to the `shift` algorithm, pun very much intended... can we create a `func WithShiftEncoder/Decoder(b byte) byte { ... }` as with the Caesar one? No, we can't. We need a `distance` parameter to tell the shift cipher with distance to use, we would need something like `func WithShiftEncoder/Decoder(b, dist byte) byte { ... }`. So what can we do about this? Simply let's use a function uses the distance `dist` and returns a `func (byte) byte`, like so:

```go
func WithShiftEncoder(dist byte) func(byte) byte { return func(b byte) byte { return rotate(b + dist) } }
func WithShiftDecoder(dist byte) func(byte) byte { return func(b byte) byte { return rotate(b - dist) } }
```

It looks a little cryptic but what happens here is that we use the functions `WithShiftEncoder` and `WithShiftDecoder` and the `dist` parameter to prepare the function that we will be passing as behavior for the Shift cipher. Then similarly to what we have done for the Caesar cipher the client calls for the Shift one will look like this

```go
dist := 1 // or any byte var/param
Transform("messagetoencode", WithShiftEncoder(dist))
Transform("encodedmessage", WithShiftDecoder(dist))
```

Notice that we are not looking anymore at how a generic cipher algorithm should work but merely at the __behaviour__ that it should use for the encoding/decoding a message.

## Extend to other ciphers

Let's go beyond and add more ciphers, for example the atbash cipher where you replace every letter with it's mirrored one, e.g. 'a' becomes 'z' and 'z' becomes 'a'.

For this cipher coding the behavior is pretty easily done via

```go
func withAtbash(b byte) byte { return 'a' + 'z' - b }
```

As this is a reflective function we can apply once to encode and once again to decode

```go
Transform("messagetoencode", withAtbash) // encode
Transform(Transform("messagetoencode", withAtbash), withAtbash) // encode and decode
```

Another example would be the vigenere cipher, a shift-based cipher where the distance is deduced from a key string composed of letters from 'a' to 'z'.

```go
type vigenere struct {
    key string
    idx int
}

func (v *vigenere) nextDist() byte {
    dist := v.key[v.idx%len(v.key)] - 'a';
    v.idx++;
    return dist
}
func withVigenereEncoder(v interface{ nextDist() byte }) func(byte) byte {
    return func(b byte) byte { return rotate(b + v.nextDist()) }
}
func withVigenereDecoder(v interface{ nextDist() byte }) func(byte) byte {
    return func(b byte) byte { return rotate(b - v.nextDist()) }
}
```

Eventually if you go in depth on vigenere, shift and caesar you will notice that caesar can be reduced to shift which can be reduced to vigenere.

Full reference can be located in a gist

How about communicating via network channels?
