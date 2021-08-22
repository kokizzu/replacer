package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
)

func main() {
	args := os.Args

	if len(args) < 5 {
		fmt.Println(`
usage:
  ` + args[0] + ` [searchTerm] [replaceTerm] [afterPrefix] [fileName]`)
		return
	}

	searchTerm := args[1]
	replaceTerm := args[2]
	afterPrefix := args[3]
	fileName := args[4]
	replaceCount := 0
	buff := bytes.Buffer{}

	{
		file, err := os.Open(fileName)
		if L.IsError(err, `os.Open file failed: `+fileName) {
			return
		}

		defer func() {
			err := file.Close()
			L.IsError(err, `file.Close failed: `+fileName)
		}()

		scanner := bufio.NewScanner(file) // max 64k
		startReplace := false
		for scanner.Scan() {
			line := scanner.Text()
			if startReplace {
				newLine := S.Replace(line, searchTerm, replaceTerm)
				if newLine != line {
					replaceCount++
					line = newLine
				}
			} else if S.StartsWith(line, afterPrefix) {
				startReplace = true
			}
			buff.WriteString(line)
			buff.WriteString("\n")
		}

		if err := scanner.Err(); L.IsError(err, `error while scanning`) {
			return
		}
	}

	if replaceCount == 0 {
		fmt.Println(`no searchTerm found`)
		return
	}

	err := ioutil.WriteFile(fileName, buff.Bytes(), 0755)
	if L.IsError(err, `ioutil.WriteFile failed: `+fileName) {
		return
	}

	fmt.Printf("Done %d replacement\n", replaceCount)
}
