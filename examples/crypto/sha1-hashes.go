package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	// Go implements several hash functions in various crypto/* packages.

	fmt.Println("SHA1 Hashes")

	s := "sha1 this string"

	// the pattern for generating a hash is sha1.new
	// sha1.write(bytes) and then sha1.Sum([]byte{})
	h := sha1.New()

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice.
	// The argument to Sum can be used to append to an existing
	// byte slice: it usually isn’t needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the %x format verb to convert a
	// hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// You can compute other hashes using a similar pattern to the one shown above.
	// For example, to compute MD5 hashes import crypto/md5 and use md5.New().
	// Note that if you need cryptographically secure hashes, you should carefully
	// research hash strength!
	// https://en.wikipedia.org/wiki/Cryptographic_hash_function

}
