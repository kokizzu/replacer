# Replacer

A simple tool to find and replace substring per line (that are should be way slower than `sed`), but way simpler to use for simple non-regex use cases. 

## Example
```
$ echo 'a b c
c d e
afterThisLine
a b c
c d e' > a.txt

$ replacer c ayaya afterThisLine a.txt
Done 2 replacement

$ cat a.txt
a b c
c d e
afterThisLine
a b ayaya
ayaya d e
```
