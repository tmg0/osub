package main

import (
	"fmt"
	"os"
	"osub/pkg/console"
)

func main() {
	if err := console.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
