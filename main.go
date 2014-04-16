package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	config := NewConfig(Arguments.ConfigFile)
	fmt.Println("Config", config.Test)
}
