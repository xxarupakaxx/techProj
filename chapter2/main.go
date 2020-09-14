package main

import (
	"log"
	"os/exec"
	"runtime"
)

func main() {

	var cmd exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "myapp.bat")
	} else {
		cmd = exec.Command("/bin/sh", "-c", "myapp.sh")

	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
