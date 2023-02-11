package main

import (
	"github.com/issue9/watermark"
)

func main() {
	w, err := watermark.NewFromFile("./mark1.png", 0, watermark.TopLeft)
	if err != nil {
		panic(err)
	}

	err = w.MarkFile("./1.jpg")
	if err != nil {
		panic(err)
	}
}
