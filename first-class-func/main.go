package main

import (
	"fmt"
	"strings"
	"unicode"
)

func normalize(in string) []byte {
	var res strings.Builder
	for _, r := range in {
		if !unicode.IsLetter(r) {
			continue
		}
		res.WriteRune(unicode.ToLower(r))
	}
	return []byte(res.String())
}

func transform(in string, transformer func(byte) byte) string {
	out := normalize(in)
	for i, letter := range out {
		out[i] = transformer(letter)
	}
	return string(out)
}

func composeTransform(in string, transformers ...func(byte) byte) string {
	out := normalize(in)
	for _, transformer := range transformers {
		for i, letter := range out {
			out[i] = transformer(letter)
		}
	}
	return string(out)
}

func withCaesarEncoder(b byte) byte { return withShiftEncoder(3)(b) }
func withCaesarDecoder(b byte) byte { return withShiftDecoder(3)(b) }
func withShiftEncoder(dist byte) func(byte) byte {
	return func(b byte) byte { return rotate(b + dist) }
}
func withShiftDecoder(dist byte) func(byte) byte {
	return func(b byte) byte { return rotate(b - dist) }
}

type vigenere struct {
	key string
	idx int
}

func (v *vigenere) nextDist() byte { dist := v.key[v.idx%len(v.key)] - 'a'; v.idx++; return dist }
func withVigenereEncoder(v interface{ nextDist() byte }) func(byte) byte {
	return func(b byte) byte { return rotate(b + v.nextDist()) }
}
func withVigenereDecoder(v interface{ nextDist() byte }) func(byte) byte {
	return func(b byte) byte { return rotate(b - v.nextDist()) }
}

func withAtbash(b byte) byte { return 'a' + 'z' - b }

func rotate(b byte) byte {
	switch {
	case b < 'a':
		b += ('z' - 'a' + 1)
	case b > 'z':
		b -= ('z' - 'a' + 1)
	}
	return b
}

func main() {
	s := "hello world"
	fmt.Println(string(normalize(s)), transform(s, withCaesarEncoder), composeTransform(s, withCaesarDecoder, withCaesarEncoder))
	fmt.Println(string(normalize(s)), transform(s, withShiftEncoder(5)), composeTransform(s, withShiftDecoder(5), withShiftEncoder(5)))
	v := &vigenere{key: "alpha"}
	fmt.Println(string(normalize(s)), transform(s, withVigenereEncoder(v)), composeTransform(s, withVigenereDecoder(v), withVigenereEncoder(v)))
	fmt.Println(string(normalize(s)), transform(s, withAtbash), composeTransform(s, withAtbash, withAtbash))
}
