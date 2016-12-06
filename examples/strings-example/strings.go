package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//:= is a shorthand for declaring and initializing a variable
	// e.g var str = "HI, I AM UPPER CASE" in this case
	str := "HI, WELCOME TO GOLANG"

	// need to define the type for uninitialized variables
	var test string
	fmt.Println("test = ", test)

	// go will infer type of initialized variables
	// string is this case
	var lower = strings.ToLower(str)
	fmt.Println(lower)

	//contains is case sensitive
	//in case of insensitive contains write a func that converts the input
	// strings to lower case and then call strings.Contains
	if strings.Contains(lower, "golang") {
		fmt.Println("Yes, exists")
	}

	//strings can be used as character array
	// Print characters starting from 4 upto 10 [10th is excluded]
	fmt.Println("Print Characters 4 to 10 -" + str[4:10])

	//print first 3 Characters [0,1,2]
	fmt.Println("First 3 characters - " + str[:3])

	//print all characters starting from 4th character
	fmt.Println("Print all after 3rd character -" + str[4:])

	//split a string
	//returns an array of words
	words := strings.Split(str, ",")
	fmt.Printf("words - %v\n", words)

	//split a string using Fields, here the separator is whitespace
	//returns an array of words
	words = strings.Fields(str)
	fmt.Printf("words - %v\n", words)

	//print all characters in a string
	for _, char := range str {
		fmt.Printf("Character %c \n", char)
	}

	//"strings are built from bytes, not characters" - Rob Pike
	//https://blog.golang.org/strings
	//print rune values with the index, %#U prints the value as well as unicode
	for index, runeValue := range str {
		fmt.Printf("The RuneValue %#U at byte position %d\n", runeValue, index)
	}

	//case of multi byte unicode values
	//traversing one rune at a time using for range loop
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	//instead of traversing using a for range loop
	// traverse a string one rune at a time by getting the width of the unicode
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

	//converting a string to array of runes
	runes := []rune("日本語")
	for _, rune := range runes {
		fmt.Printf("%#U is the rune value\n", rune)
	}

}
