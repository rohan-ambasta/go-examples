package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("regular expressions")

	// test whether a pattern matches a string
	isMatch, err := regexp.MatchString("p[a-z]+ch", "peach")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isMatch)

	// compile a pattern and keep using it with anywhere you want
	regex, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(regex)

	// match the regex with an input string
	fmt.Println("peach is a match - ", regex.MatchString("peach"))

	// find a string from an input string using regex
	fmt.Println("What did we find - ", regex.FindString("peaches peach punch"))

	// return the first and last index of the sub string if we find a match
	fmt.Println("Index of peach is - ", regex.FindStringIndex("peaches peach pucnch"))

	// will match strings matching the complete pattern as well as the inner pattern - ([a-z]+)
	fmt.Println(regex.FindStringSubmatch("peach punch"))

	// will return index of strings matching all patterns
	fmt.Println(regex.FindStringSubmatchIndex("peach punch"))

	// return all matches
	fmt.Println(regex.FindAllString("peaches peach punch", -1))

	// return first two matches
	fmt.Println(regex.FindAllString("peaches peach punch", 2))

	// match a byte array
	fmt.Println(regex.Match([]byte("peach")))

	// MustCompile is like Compile but panics if the expression cannot be parsed.
	// It simplifies safe initialization of global variables holding compiled regular
	// expressions.
	r := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// replace subsets of matched strings with other strings
	fmt.Println(regex.ReplaceAllString("a bunch of peaches", "<fruit>"))

	// replace the matched string using a function
	in := []byte("a peach")
	out := regex.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("upper case - ", string(out))

}
