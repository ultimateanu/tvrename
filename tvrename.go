package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//fmt.Println(findShow("legend of korra$%"))
	if len(os.Args) < 2 {
		log.Fatal("usage: tvrename show_directory")
	}
	fixShowDir(os.Args[1])
}

func fixShowDir(dir string) {
	//fmt.Println(filepath.Base(dir))
	fmt.Println(findShow(filepath.Base(dir)))
}
