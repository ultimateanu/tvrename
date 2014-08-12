package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: tvrename show_directory")
	}
	fixShowDir(os.Args[1])
}

func fixShowDir(dir string) {
	shows, _ := findShow(filepath.Base(dir))

	var showNames []string
	for _, s := range shows {
		showNames = append(showNames, s.String())
	}

	i, err := selectOption(showNames)
	if err != nil {
		log.Fatal("no matching show")
	}
	fmt.Println(shows[i])
}
