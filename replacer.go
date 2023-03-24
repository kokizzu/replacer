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
  ` + args[0] + ` [searchTerm] [replaceTerm] [afterTerm] [fileName]

optional argument before [searchTerm]
  -afterprefix afterTerm will become prefix search instead of default substrsearch
  -untilsubstr [untilSubstring] 
  -untilprefix [untilPrefix]

all case sensitive except for flags
`)
		return
	}

	afterTermIsSubstring := true
	requiredPos := 0
	untilSubstr := ``
	if S.ToLower(args[requiredPos+1]) == `-afterprefix` {
		afterTermIsSubstring = false
		requiredPos += 1
	}
	if len(args) < requiredPos+5 {
		fmt.Println(`error: missing required argument, remains: `, args[requiredPos+1:])
		return
	}
	if S.ToLower(args[requiredPos+1]) == `-untilsubstr` {
		untilSubstr = args[requiredPos+2]
		requiredPos += 2
	}
	if len(args) < requiredPos+5 {
		fmt.Println(`error: missing required argument, remains: `, args[requiredPos+1:])
		return
	}
	untilPrefix := ``
	if S.ToLower(args[requiredPos+1]) == `-untilprefix` {
		untilPrefix = args[requiredPos+2]
		requiredPos += 2
	}
	if len(args) < requiredPos+5 {
		fmt.Println(`error: missing required argument, remains: `, args[requiredPos+1:])
		return
	}

	searchTerm := args[requiredPos+1]
	replaceTerm := args[requiredPos+2]
	afterTerm := args[requiredPos+3]
	fileName := args[requiredPos+4]
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
		stopReplace := false
		for scanner.Scan() {
			line := scanner.Text()
			if startReplace {
				if !stopReplace {
					newLine := S.Replace(line, searchTerm, replaceTerm)
					if newLine != line {
						replaceCount++
						line = newLine
					}
					if untilSubstr != `` && S.Contains(line, untilSubstr) {
						stopReplace = true
					}
					if untilPrefix != `` && S.StartsWith(line, untilPrefix) {
						stopReplace = true
					}
				}
			} else {
				if afterTermIsSubstring {
					if S.Contains(line, afterTerm) {
						startReplace = true
					}
				} else {
					if S.StartsWith(line, afterTerm) {
						startReplace = true
					}
				}
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
