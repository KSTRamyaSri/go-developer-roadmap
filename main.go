package main

import (
	"fmt"

	"go-developer-roadmap/01-fundamentals"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(
		fundamentals.InvestmentCalc(1000, 0.05, 1),
	)

	fmt.Println("Global Variable:", fundamentals.GlobalVar)

	fundamentals.Variables()
}