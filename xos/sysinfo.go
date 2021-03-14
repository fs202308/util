package xos

import (
	"fmt"

	linux "github.com/Anlandme/sysinfo"
)

func GetCpuInfo() {
	cpuinfo, _ := linux.ReadCPUInfo("/proc/cpuinfo")
	fmt.Println(cpuinfo.Processors[0].MHz)
}
