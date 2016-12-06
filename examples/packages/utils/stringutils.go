package utils

//ExportedVar this is an exported variable
//starts with capital letter and needs a comment that starts
//with the name of the exported variable as it is done here
// getting a warning if I use the shorthand declaration :=, looks like
// declaration is required outside function body
var ExportedVar = "gnalog"

func reverse(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//Reverse Exported function reverse also needs a comment
func Reverse(input string) string {
	return reverse(input)
}
