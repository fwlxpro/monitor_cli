package systems


import (
	"os/exec"
)

func GetUser()(string, error){
 
	cmd := exec.Command("whoami")
	output,err := cmd.Output()
  
	sname := string(output)
 
if err != nil {
	 return "",err
 }

 return sname,nil 
}

