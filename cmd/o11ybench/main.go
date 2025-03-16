package main

import (
	"fmt"
	"os"

	"github.com/zyy17/o11ybench/pkg/cmd/root"
)

func main() {
	if err := root.NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
