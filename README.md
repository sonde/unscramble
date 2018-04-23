# unscramble
`unscramble.go` reads a dictionary from a dict file, register all words
in a length indexed map with the sorted version of each word as
key and the word as value. `func unscramble` is called with the input
sentence and dictionary reference as parameters, returning an
unscrambled sentence and a true/false (false meaning we failed to
unscrable, and true meaning we managed to unscramble to a legal output - 
not necessarily the intended sentence as a large dictionary result in multiple matches)

Example: `elhloothtedrowl => hello to the world`

# Examples:
```
$ go build
$ unscramble imaatactonagod
imaatactonagod => i am a cat not a dog
$ unscramble imaatactonagod dict.large 
imaatactonagod => matai octan goad
```

## Testing
Tests are implemented in unscramble_test.go:
```
$ go test
PASS
ok  	unscramble	2.727s
```
