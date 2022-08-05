package main

import (
	"go-stress-test/components"
	"os"
)

func main() {
	err := os.Setenv("CURRENT_ENV", "terminal")
	if err != nil {
		return
	}
	components.InitRequest()
	components.Start()
}
