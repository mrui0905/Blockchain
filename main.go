package main

import (
	"os"
	"github.com/mrui0905/Blockchain/cli"
	//"github.com/mrui0905/Blockchain/wallet"
)

func main() {
	defer os.Exit(0) // Error management
	cmd := cli.CommandLine{}
	cmd.Run()
}