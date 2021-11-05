package network

import (
	"fmt"
	"strings"

	netwk "github.com/shirou/gopsutil/v3/net"
)

// Get memory statistics
func Get() string {
	buffer := make([]string, 0)
	x, _ := netwk.IOCounters(true)
	for _, v := range x {
		temp := fmt.Sprintf("netcard %s: %s", v.Name, v.String())
		buffer = append(buffer, temp)
	}
	return strings.Join(buffer, ",")
}
