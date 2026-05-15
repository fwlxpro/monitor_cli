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
		return 0, err 
	}

	s := string(output)
	row := strings.Slice(s, "\n")
	field := strings.Fields(row[1])
	diskUsage,err := strconv.ParseInt(field[1], 64)
	diskAvaible err := strconv.ParseInt(field[2], 64)
  if err != nil {
		return 0, err
	} 


 return DiskAvaible,diskUsage, nil
}
