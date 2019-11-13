package main

import (
	"os"

	"github.com/liu-xiao-guo/readjson/cmd"

	_ "github.com/liu-xiao-guo/readjson/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
