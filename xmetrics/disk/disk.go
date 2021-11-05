package disk

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

// Get disk statistics
func Get() string {
	buffer := make([]string, 0)
	ys, _ := disk.Partitions(true)
	for _, y := range ys {
		d, _ := disk.IOCounters(y.Device)
		if len(d) > 0 {
			temp := fmt.Sprintf("disk %s: %v", y.Device, d)
			buffer = append(buffer, temp)
		}
	}
	return strings.Join(buffer, ",")
}
