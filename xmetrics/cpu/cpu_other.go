//go:build !linux && !darwin
// +build !linux,!darwin

package cpu

// Get cpu statistics
func Get() (*Stats, error) {
	return nil, nil
}

// Stats represents cpu statistics
type Stats struct {
	User, System, Idle, Nice, Total uint64
}
