package main

import "strconv"

var (
	testInt    int
	testString string
	testInt64  int64
)

func main() {

	// string到int
	testInt, _ = strconv.Atoi(testString)

	// string到int64
	testInt64, _ = strconv.ParseInt(testString, 10, 64)

	// int到string
	testString = strconv.Itoa(testInt)

	// int64到string
	testString = strconv.FormatInt(testInt64, 10)
}
