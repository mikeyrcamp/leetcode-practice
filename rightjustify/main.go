package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	err := rightJustifyText(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func rightJustifyText(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	var maxWidth int
	for scan.Scan() {
		s := scan.Text()
		if len(s) > maxWidth {
			maxWidth = len(s)
		}
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	scan = bufio.NewScanner(f)
	for scan.Scan() {
		s := scan.Text()

		var pad string
		for i := 0; i < maxWidth-len(s); i++ {
			pad += " "
		}

		fmt.Println(pad + s)
	}
	return nil
}
