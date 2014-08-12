package main

import (
	"fmt"
)

const maxSelection = 5

func selectOption(options []string) (selection int32, err error) {
	fmt.Println("Please select one of the following: \n")
	for i, option := range options {
		if i >= maxSelection {
			break
		}
		fmt.Printf("%d) %s\n", i+1, option)
	}
	fmt.Println("q) <None of the above>")
	fmt.Print("\n> ")

	_, err = fmt.Scanf("%d", &selection)
	if err != nil {
		return -1, err
	}
	return selection - 1, nil
}
