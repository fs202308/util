//go:build windows
// +build windows

package disk

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type disks map[string]map[string]string

func Iostat() disks {
	return disks
}
