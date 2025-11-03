package dataset

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// PushJData return: writes the json array in the given jFile
// NOTE: the file suppose to be in the directory
// NOTE: right now the json data ends with a comma but in future this issue will be solved
// NOTE: it is for the array json-format
func PushJData(jFile string, i interface{}) {

	g, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	fs, err := os.OpenFile(jFile, os.O_APPEND|os.O_RDWR|os.O_SYNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	fs.Write([]byte(g))
	fs.Seek(-1, io.SeekEnd)
	fs.Write([]byte(","))
	fs.Seek(1, io.SeekStart)
	fs.Write([]byte("\n"))
}

// ToJSONformat  NOTE: it is not test with object Json-Format
// object Json-format:[ "compilerOption":{"strict":"true"},""..]
// it is designed for any array Json-format
// [{},{},{}]
// runs the data into json file
func ToJSONformat(path string) {
	lines := FileDataToString(path)
	j := ""
	j = strings.Join(lines, "")
	// string to rune to work with characters
	ru := []rune(j)

	// find the eof in json file in-between
	comp := regexp.MustCompile("]{")
	// finding the index to easily remove
	inx := comp.FindIndex([]byte(j))
	comp2 := regexp.MustCompile(`\[`)
	inx2 := comp2.FindIndex([]byte(j))

	// important else the range out of index
	if len(inx) != 0 {
		ru[inx[0]] = '\n'
	}
	// append the opening tag to the json file
	if len(inx2) == 0 {
		// in-short push front or prepend
		ru = append([]rune{'['}, ru...)
	}
	// check for the last char having eof
	if ru[len(ru)-1] != ']' {
		ru = append(ru, ']') // append the closing tag
	}

	// convert the char to string
	j = string(ru)
	// write the file
	err := os.WriteFile(path, []byte(j), 0644)
	if err != nil {
		panic(err)
	}
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
