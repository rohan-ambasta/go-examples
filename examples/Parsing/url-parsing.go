package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	fmt.Println("URL Parsing")

	// We’ll parse this example URL, which includes a scheme,
	// authentication info, host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the url and ensure there are no erros
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme
	fmt.Println(u.Scheme)

	// User contains username as well as authentication info
	user := u.User
	fmt.Println(user)
	// username
	fmt.Println(user.Username())
	// password
	fmt.Println(user.Password())

	// host contains host name and port
	host := u.Host
	hostname, port, _ := net.SplitHostPort(host)
	fmt.Println(hostname)
	fmt.Println(port)

	// extract the path and the fragment after the #
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// use RawQuery to get all the query params
	queryParams := u.RawQuery
	fmt.Println(queryParams)

	// Parse the query string using url.ParseQuery into a map[string][]string
	paramsMap, _ := url.ParseQuery(queryParams)
	fmt.Println(paramsMap)
	fmt.Println(paramsMap["k"][0])

}
