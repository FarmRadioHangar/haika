package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	script := "/home/fri/fessbox/backend-node/sh/past-hour-metrics.sh"
	table := []struct {
		tag     string
		command string
		args    []string
	}{
		{"incoming_sms", script, []string{"sms_in"}},
		{"incoming_call", script, []string{"from ring to master"}},
	}
	rst := make(map[string]interface{})
	for _, v := range table {
		out, err := exec.Command(v.command, v.args...).Output()
		if err != nil {
			os.Exit(1)
		}
		i, err := strconv.Atoi(string(out))
		if err != nil {
			os.Exit(1)
		}
		rst[v.tag] = i
	}
	json.NewEncoder(os.Stdout).Encode(rst)
}
