package main

import (
	"fmt"
	"log"
	systems "pimonitor/systems"
)

func main() {

	log.SetFlags(0)
  
  uptime , err := systems.Uptime()
	if err != nil {
		log.Fatal(err)
	}

	usage, err := systems.RamUsage()
	if err != nil {
		log.Fatal(err)
	}

	temp, err := systems.CpuTemp()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("CPU Temperature: %.1f°C\n", temp)
	fmt.Printf("Ram Usage: %.1f\n", usage)
	fmt.Printf("Uptime: %d days\n",uptime )
}
