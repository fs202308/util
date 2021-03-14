package xos

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bsync-tech/util/xmetrics/cpu"
	"github.com/bsync-tech/util/xmetrics/disk"
	"github.com/bsync-tech/util/xmetrics/memory"
	"github.com/mackerelio/go-osstat/network"
)

func Cpu() (string, error) {
	before, err := cpu.Get()
	if err != nil {
		return "", err
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		return "", err
	}
	total := float64(after.Total - before.Total)
	user := fmt.Sprintf("%.2f%%", float64(after.User-before.User)/total*100)
	system := fmt.Sprintf("%.2f%%", (float64(after.System-before.System)/total)*100)
	idle := fmt.Sprintf("%.2f%%", (float64(after.Idle-before.Idle)/total)*100)
	info := map[string]interface{}{
		"User":   user,
		"System": system,
		"Idle":   idle,
	}
	d, err := json.Marshal(&info)
	if err != nil {
		return "", err
	}
	fmt.Println(string(d))
	return string(d), nil
}

func Memory() (string, error) {
	memory, err := memory.Get()
	if err != nil {
		return "", err
	}
	total := fmt.Sprintf("%dMB", memory.Total/1024/1024)
	used := fmt.Sprintf("%dMB", memory.Used/1024/1024)
	cached := fmt.Sprintf("%dMB", memory.Cached/1024/1024)
	free := fmt.Sprintf("%dMB", memory.Free/1024/1024)
	info := map[string]interface{}{
		"Total":  total,
		"Used":   used,
		"Cached": cached,
		"Free":   free,
	}
	d, err := json.Marshal(&info)
	if err != nil {
		return "", err
	}
	fmt.Println(string(d))
	return string(d), nil
}

func Disk() (string, error) {
	before, err := disk.Get()
	begin := time.Now()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	time.Sleep(time.Duration(1000) * time.Millisecond)
	after, err := disk.Get()
	elapse := time.Since(begin)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	ds := make([]*map[string]interface{}, len(after))
	for i := 0; i < len(after); i++ {
		info := map[string]interface{}{
			"DeviceName": after[i].DeviceName,
			"ReadUtil":   fmt.Sprintf("%.2f%%", 100*float64(after[i].TimeSpentReading-before[i].TimeSpentReading)/float64(elapse.Milliseconds())),
			"WriteUtil":  fmt.Sprintf("%.2f%%", 100*float64(after[i].TimeSpentWriting-before[i].TimeSpentWriting)/float64(elapse.Milliseconds())),
		}
		ds = append(ds, &info)
	}

	d, err := json.Marshal(&ds)
	if err != nil {
		return "", err
	}
	fmt.Println(string(d))
	return string(d), nil
}

func Network() (string, error) {
	before, err := network.Get()
	begin := time.Now()
	if err != nil {
		return "", err
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := network.Get()
	elapse := time.Since(begin)
	if err != nil {
		return "", err
	}
	ds := make([]*map[string]interface{}, len(after))
	for i := 0; i < len(after); i++ {
		info := map[string]interface{}{
			"DeviceName": after[i].Name,
			"Read":       fmt.Sprintf("%.2fBytes", 100*float64(after[i].RxBytes-before[i].RxBytes)/float64(elapse.Milliseconds())),
			"Write":      fmt.Sprintf("%.2fBytes", 100*float64(after[i].TxBytes-before[i].TxBytes)/float64(elapse.Milliseconds())),
		}
		ds = append(ds, &info)
	}

	d, err := json.Marshal(&ds)
	if err != nil {
		return "", err
	}
	fmt.Println(string(d))
	return string(d), nil
}
