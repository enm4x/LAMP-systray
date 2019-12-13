package main

import (
	"log"
	"os"
	"os/exec"
)

func cmdexec(rqstdAct string) {

	if rqstdAct == "start" {
		cmd := exec.Command("/bin/sh", "-c", "echo cmd:systemctl start apache2 mysql; systemctl start apache2 mysql; echo apache2 and mysql has been started")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	} else if rqstdAct == "stop" {
		cmd := exec.Command("/bin/sh", "-c", "echo cmd: systemctl stop apache2 mysql; systemctl stop apache2 mysql; echo apache2 and mysql has been stoped")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}

}
