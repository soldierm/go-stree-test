package main

import (
	"go-stress-test/components"
	"os"
)

func main() {
	os.Setenv("CURRENT_ENV", "terminal")
	components.InitRequest()
	components.Start()
}
