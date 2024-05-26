package server

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Printascii(input , banner string) string {
	res := ""
	style := "./func/" + banner + ".txt"
	r := []rune(input)
	for i := 0; i < len(r); i++ {
		if r[i] < 32 ||
			r[i] > 127 {
			fmt.Println("you need to choose a character from the ascii table.")
			os.Exit(1)
		}
	}
	words := strings.Split(input, "\\n")

	if checker(words) {
		for k := 0; k < len(words)-1; k++ {
			res += "\n"
		}
		return res
	}
	file, err := os.Open(style)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	for _, word := range words {
		var char []string
		var chars [][]string
		r = []rune(word)
		if word == "" {
			fmt.Println("")
			continue
		}
		for i := 0; i < len(r); i++ {
			char = printingchar(r[i], file)
			chars = append(chars, char)
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(chars); j++ {
				res += chars[j][i]
			}
			res += "\n"
		}
	}
	defer file.Close()
	return res
}

func checker(s []string) bool {
	for _, t := range s {
		if t != "" {
			return false
		}
	}
	return true
}

func printingchar(r rune, file *os.File) []string {
	n := int(r - 32) 
	startLine := 3 + (n)*9 
	// fmt.Println(startLine)
	endLine := startLine + 8
	// Seek to the beginning of the file
	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("error: ",err)
		os.Exit(2)
	}
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Scan lines from the file
	lineNumber := 1
	var lines []string
	for scanner.Scan() {
		lineNumber++
		if lineNumber < startLine {
			continue // Skip until the start line
		}
		if lineNumber > endLine {
			break // Stop scanning after the end line
		}
		line := scanner.Text() // Get the current line
		// fmt.Println(line)      // Process the line (print in this case)
		lines = append(lines, line)
	}
	// Check for errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	return lines
}
