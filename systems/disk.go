package systems

import (
	"os/exec"
	"strings"
	"strconv"
)

func DiskUsage()(int,int, error){
  
	cmd := exec.Command("df", "--output=used,avail", "/", "--block-size=1G")
	output, err := cmd.Output()
  if err != nil {
		return 0,0, err 
	}

	s := string(output)
	row := strings.Split(s, "\n")
	field := strings.Fields(row[1])
	diskUsage,err := strconv.Atoi(field[1])
	diskAvaible, err := strconv.Atoi(field[2])
  if err != nil {
		return 0,0, err
	} 


 return diskAvaible,diskUsage, nil
}
