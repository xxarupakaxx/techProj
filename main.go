package main

import (
	"log"
	"os/user"
	"path/filepath"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// "\"でパス文字列を統合しない
	dir := filepath.Join(u.HomeDir, ".config", "myapp")
	if err != nil {
		log.Fatal(err)
	}
}
