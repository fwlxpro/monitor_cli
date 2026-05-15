package systems

import (
	"os/exec"
	"strconv"
	"strings"
)

func CpuTemp() (float64, error) {

	// Chci vytvořit funkci která za pomocí commandu zjistí teplotu a  následně jí vrátí
	cmd := exec.Command("vcgencmd", "measure_temp")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	s := string(output)
	parts := strings.Split(s, "=")
	final_part := strings.TrimSuffix(parts[1], "'C")
	temp, err := strconv.ParseFloat(final_part, 64)

	// Výstup bude ve formátu "temp=XX.X'C", takže musíme extrahovat číslo

	if err != nil {
		return 0, err
	}

	return temp, nil
}
