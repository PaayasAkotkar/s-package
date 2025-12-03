// Package dataset algorithm related os

package dataset

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// EIndex process of enum
type EIndex int

const (
	firstIndex EIndex = iota
	lastIndex
	midIndex
)

// IOSeeFirstMiddleLast return value at the first, middle, last in the file
// NOTE: not all
func IOSeeFirstMiddleLast(filePath string, index EIndex) string {
	element := ""
	fs, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	fls, err := fs.Stat()
	if err != nil {
		panic(err)
	}

	// get the index of the first char
	// i1Char := make([]byte, fls.Size()-(fls.Size()-1))
	i1Char := make([]byte, fls.Size()-(fls.Size()-1))

	// get the max len of the file
	iChar := make([]byte, fls.Size())

	// get the first char
	fChar, err := fs.Read(i1Char)
	if err != nil {
		panic(err)
	}

	fChar2, err := fs.Read(iChar)
	if err != nil {
		panic(err)
	}

	switch index {
	case firstIndex:
		{
			element = string(i1Char[:fChar])
			break
		}
	case midIndex:
		{
			element = string(iChar[fChar2])
			break
		}
	case lastIndex:
		{
			element = string(iChar[fChar2-1])
			break
		}
	default:
		{
			element = "not found"
		}

	}
	return element
}

// IOGetFirstMiddleLast return index at the first, middle, last in the file
// NOTE: not all
func IOGetFirstMiddleLast(filePath string, index EIndex) int {
	elementIndex := 1
	fs, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	fls, err := fs.Stat()
	if err != nil {
		panic(err)
	}

	//// get the max len of the file
	iChar := make([]byte, fls.Size())

	fChar2, err := fs.Read(iChar)
	if err != nil {
		panic(err)
	}

	switch index {
	case firstIndex:
		{

			elementIndex = 0
			break
		}
	case midIndex:
		{
			elementIndex = int(iChar[fChar2-2])

			break
		}
	case lastIndex:
		{
			elementIndex = int(iChar[fChar2-1])
			break
		}
	default:
		{
			elementIndex = int(iChar[fChar2-2])
		}

	}
	return elementIndex
}

// FileDataToString return converts file data to string
func FileDataToString(filepath string) []string {
	//// o.App, os.Create
	fs, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fs.Close()
	lines := []string{}
	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		//// with this we will get the data at each line
		if trimmedLine != "" {
			lines = append(lines, trimmedLine)
		}
	}
	return lines
}

// Example demonstrates the use of open and format to array string
func Example() {
	pattern := regexp.MustCompile(`\b(?:GK|DF|MD|FW|\d{1,2}\s+(?:Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\s+\d{4}.*$)\b|[0-9]`)
	file := "test.txt"
	Open(file, *pattern)
	FormatToArray(file)
	// output:
	// "Iker CASILLAS",
	// "Cristobal CURRO TORRES",
	// "JUAN FRANcisco Garcia",
	// "Ivan HELGUERA",
	// "Carles PUYOL",
	// "Fernando HIERRO",
	// "RAUL Gonzalez Blanco",
	// "Ruben BARAJA",
	// "Fernando MORIENTES",
	// "Diego TRISTAN",
	// "Francisco DE PEDRO",
	// "Alberto LUQUE",
	// "RICARDO Lopez Felipe",
	// "David ALBELDA Aliques",
	// "Enrique ROMERO",
	// "Gaizka MENDIETA            Mar    Lazio (ITA)",
	// "Juan Carlos VALERON",
	// "SERGIO Gonzalez Soriano",
	// "XAVI Hernandez",
	// "Miguel Angel NADAL",
	// "LUIS ENRIQUE",
	// "JOAQUIN Sanchez",
	// "Pedro CONTRERAS",
}

// Open opens the file and removes all the requested data
func Open(f string, pattern regexp.Regexp) {
	fs, err := os.OpenFile(f, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fs.Close()
	lines := []string{}
	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		line := scanner.Text()
		line = string(pattern.ReplaceAll([]byte(line), []byte("")))
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	sentences := strings.Join(lines, "\n")

	if err := os.WriteFile(f, []byte(sentences), 0644); err != nil {
		panic(err)
	}
}

// FormatToArray converts the data to array
// the format must be:
// input:
// hey i am king
// i have 10 kids
// i have 5 wives
// output:
// "hey i am key"
// "i have 10 kids"
// "i have 5 wives"
func FormatToArray(filepath string) {
	fs, err := os.OpenFile(filepath, os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer fs.Close()
	lines := []string{}
	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		mod := fmt.Sprintf("\"%s\",", line) // does "line"

		lines = append(lines, mod)
	}
	sentences := strings.Join(lines, "\n")
	if err := os.WriteFile(filepath, []byte(sentences), 0644); err != nil {
		panic(err)
	}
}
