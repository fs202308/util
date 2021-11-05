package main

import (
	"fmt"

	"github.com/bsync-tech/util/xmetrics/cpu"
	"github.com/bsync-tech/util/xmetrics/disk"
	"github.com/bsync-tech/util/xmetrics/memory"
	network "github.com/bsync-tech/util/xmetrics/net"
)

func main() {
	fmt.Println(cpu.Get())
	fmt.Println(network.Get())
	fmt.Println(memory.Get())
	fmt.Println(disk.Get())
}
