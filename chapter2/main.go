package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("my-app")
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	b, err = japanese.shiftJIS.NewDecoder().Bytes(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.print(string(b))
}
