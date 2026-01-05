package main

import (
	"fmt"

	"gonum.org/v1/gonum"
)

func main() {
	v, c := gonum.Version()
	fmt.Println(v, c)
}
