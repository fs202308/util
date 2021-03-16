package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

type cmdRunner struct{}
type disks map[string]map[string]string

func New() *cmdRunner {
	return &cmdRunner{}
}

func (c *cmdRunner) Run(cmd string, args []string) (io.Reader, error) {
	command := exec.Command(cmd, args...)
	resCh := make(chan []byte)
	errCh := make(chan error)
	go func() {
		out, err := command.CombinedOutput()
		if err != nil {
			errCh <- err
		}
		resCh <- out
	}()
	timer := time.After(2 * time.Second)
	select {
	case err := <-errCh:
		return nil, err
	case res := <-resCh:
		return bytes.NewReader(res), nil
	case <-timer:
		return nil, fmt.Errorf("time out (cmd:%v args:%v)", cmd, args)
	}
}

func (c *cmdRunner) Exec(cmd string, args []string) string {
	command := exec.Command(cmd, args...)
	outputBytes, _ := command.CombinedOutput()
	//if err != nil {
	//	log.Error(err)
	//}
	return string(outputBytes[:])
}

func parser_iostat(r io.Reader) disks {
	iostat_info := make(disks)
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		fields_list := []string{"rrqm/s", "wrqm/s", "r/s", "w/s", "rkB/s", "wkB/s", "avgrq-sz", "avgqu-sz", "await", "r_await", "w_await", "svctm", "%util"}
		line := scan.Text()
		fields := strings.Fields(line)
		if strings.HasPrefix(line, "sd") {
			iostat_info[fields[0]] = make(map[string]string)
			for i, f := range fields_list {
				iostat_info[fields[0]][f] = fields[i+1]
			}
		}
	}
	return iostat_info
}

func Iostat() disks {
	var cmdrun = cmdRunner{}
	rr, err := cmdrun.Run("iostat", []string{"-x"})
	if err != nil {
		fmt.Println(err)
	}
	disks := parser_iostat(rr)
	return disks
}

func main() {
	fmt.Println(Iostat())
}
