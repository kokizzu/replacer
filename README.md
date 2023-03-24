# Replacer

A simple tool to find and replace substring per line (that are should be way slower than `sed`), but way simpler to use for simple non-regex use cases. 

## Install

requires [Golang](//golang.org/)
```
go install github.com/kokizzu/replacer@latest
```

## Changelog

- `2021-08-22` first version
- `2023-03-25` add 2 more flag: `-untilsbustr` and `-untilprefix` to replace only until specific substring or prefix
- `2023-03-25` 3rd parameter became substring search now, previously was prefix search, added 1 new flag: `-afterprefix`

## Usage

```
./replacer [-afterprefix] [-untilsubstr UNTIL_TERM] [-untilprefix UNTIL_TERM] FROM_SUBSTR TO_SUBSTR AFTER_TERM file
```

Must be in order

  - `-afterprefix` AFTER_TERM will become prefix search instead of default substring search
  - `-untilsubstr UNTIL_TERM` will replace only until UNTIL_TERM substring found
  - `-untilprefix UNTIL_TERM` will replace only until first UNTIL_TERM prefix found

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
