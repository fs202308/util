package memory

import (
	"github.com/shirou/gopsutil/v3/mem"
)

// Get memory statistics
func Get() string {
	v, _ := mem.VirtualMemory()

	return v.String()
}
