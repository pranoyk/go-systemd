package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	_, err := exec.Command("service", "urlshortner", "start").Output()
	if err != nil {
		fmt.Println("start service failed :: ", err.Error())
	}

	_, err = exec.Command("service", "urlshortner", "enable").Output()
	if err != nil {
		fmt.Println("start service failed :: ", err.Error())
	}
}