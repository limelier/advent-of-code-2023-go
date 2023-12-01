package common

import (
	"bufio"
	"os"
)

// ReadInput returns a channel which will contain each line of the input file
func ReadInput(relPath string) chan string {
	ch := make(chan string)
	file, err := os.Open(relPath)

	if err != nil {
		panic("Tried to read missing input file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		_ = file.Close()
		close(ch)
	}()

	return ch
}
