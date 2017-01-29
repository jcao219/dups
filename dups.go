package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("need to specify filename")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic("could not find file")
	}
	reader := bufio.NewReader(f)
	l, pre, err := reader.ReadLine()

	var lineAll []byte
	lineNo := 0
	allLines := make(map[string][]int)
	for ; err == nil; l, pre, err = reader.ReadLine() {
		lineAll = append(lineAll, l...)
		if pre {
			continue
		}
		lineNo++

		lineStr := string(lineAll)
		lNums, _ := allLines[lineStr]
		allLines[lineStr] = append(lNums, lineNo)

		lineAll = lineAll[:0] // Do not release memory back to GC
	}

	for lineText, lNums := range allLines {
		if len(lNums) > 1 {
			for i, lineNo := range lNums {
				if i == len(lNums)-1 { // Last line.
					fmt.Printf("%d\t", lineNo)
				} else {
					fmt.Printf("%d,", lineNo)
				}
			}
			fmt.Println(lineText)
		}
	}
}
