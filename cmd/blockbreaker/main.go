package main

import (
	"fmt"

	"github.com/cmarkh/blockbreaker/pkg/blockbreaker"
)

func main() {
	err := blockbreaker.Start()
	fmt.Println(err)
}
