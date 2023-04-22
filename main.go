package main

import (
	"fmt"
	"os"

	"github.com/olihough86/rottenhaddock/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
