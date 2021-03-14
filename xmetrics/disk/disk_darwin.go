// +build darwin

package disk

func Get() ([]Stats, error) {
	return []Stats{}, nil
}

// Stats represents disk I/O statistics for linux.
type Stats struct {
	MajorNumber      uint64
	MinorNumber      uint64
	DeviceName       string
	ReadsCompleted   uint64
	ReadsMerged      uint64
	SectorsRead      uint64
	TimeSpentReading uint64
	WritesCompleted  uint64
	WritesMerged     uint64
	SectorsWritten   uint64
	TimeSpentWriting uint64
	IOInProgress     uint64
	TimeSpentInIO    uint64
	WeightedTimeInIO uint64
}
