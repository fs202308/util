//go:build windows
// +build windows

package disk

type data map[string]map[string]string

func Iostat() data {
	iostat_info := make(data)

	return iostat_info
}

func Get() (data, error) {
	return Iostat()
}
