// Copyright © 2018 Alexandre Bourget <alex@eoscanada.com>

package main

import "github.com/eoscanada/eos-bios/eos-bios/cmd"

var version = "dev"

func init() {
	cmd.Version = version
}

func main() {
	cmd.Execute()
}
