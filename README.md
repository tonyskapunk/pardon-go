# Pardon? (go)

"Pardon" is a NATO Phonetic Alphabet Translator concept, implemented in various languages. Original implementations:

- Python ([Pardon](https://github.com/komish/pardon))
- Kotlin ([Pardon-Kotlin](https://github.com/komish/pardon-kotlin))
- Alledgely there was a go version somewhere but now is "go-ne".

The concept is to take a string and print out the phonetic representation of the string characters. 


## Environment

At least, what is perceived to be potentially relevant

```
$ go version
go version go1.11 linux/amd64
```

## Purpose

This project is a personal learning project. This should not be considered a marker of proper golang syntax, grammar, or concepts.

## Sample Execution

```
$ go run pardon.go hello world hunter2
"hello"

	Hotel Echo Lima Lima Oscar

"world"

	Whiskey Oscar Romeo Lima Delta

"hunter2"

	Hotel Uniform November Tango Echo Romeo Two
```

## Compile

```
$ go build pardon.go
```

## Example

[![asciicast](https://asciinema.org/a/14.png)](https://asciinema.org/a/ujFi77qkMOV4kZ6NXojs3CR2L)
