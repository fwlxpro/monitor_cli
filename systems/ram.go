package systems

import (
	"os/exec"
	"strconv"
	"strings"
)

func RamUsage() (float64, error) {
	cmd := exec.Command("free", "-m")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	s := string(output)
	lines := strings.Split(s, "\n")
	fields := strings.Fields(lines[1])
	ram_usage, err := strconv.ParseFloat(fields[2], 64)

	if err != nil {
		return 0, err
	}

	return ram_usage, nil
}
