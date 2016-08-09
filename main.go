package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	logFile := "/home/fri/fessbox/backend-node/log/debug.log"
	table := []struct {
		tag string
		arg string
	}{
		{"incoming_sms", "sms_in"},
		{"incoming_calls", "from ring to master"},
	}
	layout := "2006.01.02-15"
	now := time.Now().Format(layout)
	rst := make(map[string]interface{})
	file := logFile + "-" + now
	_, err := os.Stat(file)
	if err != nil {
		return
	}
	for _, v := range table {
		out, err := exec.Command("/usr/bin/grep",
			"-cF", fmt.Sprintf("\"%s\"", v.arg), file).Output()
		if err != nil {
			if out == nil {
				rst[v.tag] = 0
				continue
			}
		}
		i, err := strconv.Atoi(
			strings.TrimSpace(string(out)))
		if err != nil {
			fmt.Println(err)
			continue
		}
		rst[v.tag] = i
	}
	json.NewEncoder(os.Stdout).Encode(rst)
}
