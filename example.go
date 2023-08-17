package main

import (
	"fmt"

	"github.com/fs202308/util/xmetrics/cpu"
	"github.com/fs202308/util/xmetrics/disk"
	"github.com/fs202308/util/xmetrics/memory"
	network "github.com/fs202308/util/xmetrics/net"
)

func main() {
	fmt.Println(cpu.Get())
	fmt.Println(network.Get())
	fmt.Println(memory.Get())
	fmt.Println(disk.Get())
}
