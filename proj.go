package main

import (
	"os"
	"fmt"

	"github.com/es3649/proj/internal/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Failed with error `%s'\n", err)
		os.Exit(1)
	}
}