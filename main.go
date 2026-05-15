package main

import (
	"fmt"
	"log"
	systems "pimonitor/systems"
)

func main() {

	log.SetFlags(0)
 
  avai,used, err := systems.DiskUsage()
  if err != nil {
		log.Fatal(err)
	} 

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
	fmt.Printf('

    ___       ___       ___            ___       ___       ___       ___   
   /\  \     /\__\     /\  \          /\  \     /\  \     /\  \     /\__\  
  /::\  \   /:/  /    _\:\  \         \:\  \   /::\  \   /::\  \   /:/  /  
 /:/\:\__\ /:/__/    /\/::\__\        /::\__\ /:/\:\__\ /:/\:\__\ /:/__/   
 \:\ \/__/ \:\  \    \::/\/__/       /:/\/__/ \:\/:/  / \:\/:/  / \:\  \   
  \:\__\    \:\__\    \:\__\         \/__/     \::/  /   \::/  /   \:\__\  
   \/__/     \/__/     \/__/                    \/__/     \/__/     \/__/  

	')
	fmt.Printf("CPU Temperature: %.1f°C\n", temp)
	fmt.Printf("Ram Usage: %.1f\n", usage)
	fmt.Printf("Uptime: %d days\n",uptime )
	fmt.Printf("Used Disk: %d \n Avaible Storage on disk: %d ",used,avai)
}
