package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Go offers built-in support for JSON encoding and decoding,
// including to and from built-in and custom data types.

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	fmt.Println("JSON examples")

	//encoding basic data types

	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(5)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(3.14)
	fmt.Println(string(floatB))

	stringB, _ := json.Marshal("str")
	fmt.Println(string(stringB))

	fruits := []string{"apple", "mango", "grapes"}
	fruitsB, _ := json.Marshal(fruits)
	fmt.Println(string(fruitsB))

	fruitsWithLength := map[string]int{"apple": 5, "mango": 5, "grapes": 6}
	fruitsWithLengthB, _ := json.Marshal(fruitsWithLength)
	fmt.Println(string(fruitsWithLengthB))

	// The JSON package can automatically encode your custom data types.
	// It will only include exported fields in the encoded output and will
	// by default use those names as the JSON keys.
	res1 := &Response1{
		Page:   10,
		Fruits: []string{"apple", "mango", "grapes"},
	}
	res1B, _ := json.Marshal(res1)
	fmt.Println(string(res1B))

	// You can use tags on struct field declarations to customize the encoded
	// JSON key names. Check the definition of Response2 above
	// to see an example of such tags.
	res2 := &Response2{
		Page:   11,
		Fruits: []string{"apple", "mango", "grapes"},
	}
	res2B, _ := json.Marshal(res2)
	fmt.Println(string(res2B))

	// sample json data
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// map to hold json data
	var dat map[string]interface{}

	// decode json data, panic in case of error
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// In order to use the values in the decoded map, we’ll need
	// to cast them to their appropriate type. For example here we
	// cast the value in num to the expected float64 type.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of casts.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// We can also decode JSON into custom data types.
	// This has the advantages of adding additional type-safety
	// to our programs and eliminating the need for type assertions
	// when accessing the decoded data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and strings
	// as intermediates between the data and JSON representation on standard out.
	// We can also stream JSON encodings directly to os.Writers like os.Stdout or
	// even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
