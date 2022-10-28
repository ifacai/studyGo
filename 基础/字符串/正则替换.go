package main

import (
	"fmt"
	"regexp"
)

func regexpReplace(input string) string {
	sampleRegexp := regexp.MustCompile(`[^\x{4e00}-\x{9fa5}-[0-9]+`)
	result := sampleRegexp.ReplaceAllString(input, "")
	return result
}

func main() {
	str := "是我呀,123,,5512312../a~~!!@##!@"
	fmt.Println(regexpReplace(str))
}
