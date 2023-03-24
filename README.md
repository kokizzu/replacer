# Replacer

A simple tool to find and replace substring per line (that are should be way slower than `sed`), but way simpler to use for simple non-regex use cases. 

## Install

requires [Golang](//golang.org/)
```
go install github.com/kokizzu/replacer@latest
```

## Usage

```
./replacer [-afterprefix] [-untilsubstr untilSubstring] [-untilprefix untilPrefix] from to afterTerm file
```

Must be in order

  - `-afterprefix` afterTerm will become prefix search instead of default substrsearch
  - `-untilsubstr` untilSubstring 
  - `-untilprefix` untilPrefix

## Example using CLI
```
$ echo 'a b c
c d e
afterThisLine
a b c
c d e' > a.txt

$ replacer c ayaya afterThisTerm a.txt
Done 2 replacement

$ cat a.txt
a b c
c d e
afterThisLine
a b ayaya
ayaya d e
```

## Example using go:generate
```
//go:generate replacer 'Id" form' 'Id,string" form' type bla.go
//go:generate replacer 'json:"id"' 'json:id,string" form' type bla.go
//go:generate replacer 'By" form' 'By,string" form' type bla.go
```
