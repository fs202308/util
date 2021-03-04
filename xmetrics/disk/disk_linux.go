// +build linux

package disk

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 0:   major number
// 1:   minor number
// 2:   device name
// 3:   reads completed successfully
// 4:   reads merged
// 5:   sectors read
// 6:   time spent reading (ms)
// 7:   writes completed
// 8:   writes merged
// 9:   sectors written
// 10:  time spent writing (ms)
// 11:  I/Os currently in progress
// 12:  time spent doing I/Os (ms)
// 13:  weighted time spent doing I/Os (ms)

func Get() ([]Stats, error) {
	file, err := os.Open("/proc/diskstats")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return collect(file)
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

func collect(out io.Reader) ([]Stats, error) {
	scanner := bufio.NewScanner(out)
	var diskStats []Stats
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 14 {
			continue
		}
		majorNumber, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return nil, err
		}
		minorNumber, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return nil, err
		}
		readsCompleted, err := strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			return nil, err
		}
		readsMerged, err := strconv.ParseUint(fields[4], 10, 64)
		if err != nil {
			return nil, err
		}
		sectorsRead, err := strconv.ParseUint(fields[5], 10, 64)
		if err != nil {
			return nil, err
		}
		timeSpentReading, err := strconv.ParseUint(fields[6], 10, 64)
		if err != nil {
			return nil, err
		}
		writesCompleted, err := strconv.ParseUint(fields[7], 10, 64)
		if err != nil {
			return nil, err
		}
		writesMerged, err := strconv.ParseUint(fields[8], 10, 64)
		if err != nil {
			return nil, err
		}
		sectorsWritten, err := strconv.ParseUint(fields[9], 10, 64)
		if err != nil {
			return nil, err
		}
		timeSpentWriting, err := strconv.ParseUint(fields[10], 10, 64)
		if err != nil {
			return nil, err
		}
		ioInProgress, err := strconv.ParseUint(fields[11], 10, 64)
		if err != nil {
			return nil, err
		}
		timeSpentInIO, err := strconv.ParseUint(fields[12], 10, 64)
		if err != nil {
			return nil, err
		}
		weightedTimeInIO, err := strconv.ParseUint(fields[13], 10, 64)
		if err != nil {
			return nil, err
		}
		diskStats = append(diskStats, Stats{
			MajorNumber:      majorNumber,
			MinorNumber:      minorNumber,
			DeviceName:       fields[2],
			ReadsCompleted:   readsCompleted,
			ReadsMerged:      readsMerged,
			SectorsRead:      sectorsRead,
			TimeSpentReading: timeSpentReading,
			WritesCompleted:  writesCompleted,
			WritesMerged:     writesMerged,
			SectorsWritten:   sectorsWritten,
			TimeSpentWriting: timeSpentWriting,
			IOInProgress:     ioInProgress,
			TimeSpentInIO:    timeSpentInIO,
			WeightedTimeInIO: weightedTimeInIO,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for /proc/diskstats: %s", err)
	}
	return diskStats, nil
}