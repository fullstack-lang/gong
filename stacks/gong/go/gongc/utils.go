package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

/**
 * Insert sting to n-th line of file.
 * If you want to insert a line, append newline '\n' to the end of the string.
 */
func InsertStringToFile(path, str string, preceedingLineContent string) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for i, line := range lines {
		fileContent += line
		fileContent += "\n"
		if strings.Contains(lines[i], preceedingLineContent) {
			fileContent += str
			fileContent += "\n"
		}
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}
