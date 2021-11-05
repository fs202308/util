package cpu

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

// Get cpu statistics
func Get() string {
	w, _ := cpu.Percent(time.Second, true)
	buffer := make([]string, 0)
	for i := range w {
		temp := fmt.Sprintf("cpu%d:%.2f%%", i, w[i])
		buffer = append(buffer, temp)
	}
	return strings.Join(buffer, ", ")
}
