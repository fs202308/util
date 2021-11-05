//go:build windows
// +build windows

package disk

type disks map[string]map[string]string

func Iostat() disks {
	return disks
}
