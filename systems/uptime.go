package systems 


import ( 
"strconv"
"os/exec"
"strings"

)

// 1. Závorka co funckce příjmá 2. závorka co funkce vrací (return)
func Uptime()(int, error){
// začínáme CMD my chceme s těmi daty dál pracovat tak že si je musíme pokaždé nějak pojmenovat takhle funguje GO
    cmd := exec.Command("uptime")
  /* data mi to ale ještě stále nevypsalo nedalo jen našlo (dáme si to do output se kterým budeme dál pracovat) err slouží pokud by se nám třeba už data ani nedala */    
    output, err := cmd.Output()
    
    if err != nil {
        return 0, err
    }

 // teď to potebuji rozdělit 
 s := string(output)
 field := strings.Fields(s)
 days, err := strconv.Atoi(field[2])

 if err != nil {
	return 0, err
}


return days , nil
}

 
