package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Go PseudoRandom number generators")

	// Go’s math/rand package provides pseudorandom number generation.

	// What is Pseudorandom number generator?
	// A pseudorandom number generator (PRNG), also known as a deterministic
	// random bit generator (DRBG),[1] is an algorithm for generating a sequence
	// of numbers whose properties approximate the properties of sequences of random numbers.
	// The PRNG-generated sequence is not truly random, because it is completely
	// determined by a relatively small set of initial values, called the
	// PRNG's seed (which may include truly random values). Although sequences that are
	// closer to truly random can be generated using hardware random number generators,
	// pseudorandom number generators are important in practice for their speed in number
	// generation and their reproducibility
	// Please the wikipedia link for detail information
	// [https://en.wikipedia.org/wiki/Pseudorandom_number_generator]

	// For example, rand.Intn returns a random int n, 0 <= n < 100.
	// If you observe the output of this code, you will see that
	// same random number sequence is generated every time your run this code
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
	fmt.Println(rand.Float64())

	// This can be used to generate random floats in other ranges, for example 5.0 <= f' < 10.0.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// The default number generator is deterministic, so it’ll produce
	// the same sequence of numbers each time by default. To produce
	// varying sequences, give it a seed that changes. Note that this is not
	// safe to use for random numbers you intend to be secret, use crypto/rand for those.
	// In the below code create source s1 and use it for random number generator r1
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Now the random number sequence changes from the one with default seed
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// If you seed a source with the same number, it produces the same sequence of random numbers.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))

}
