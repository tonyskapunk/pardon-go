package main

import (
    "fmt"
    "os"
    "strings"
)

// Map of letters to the translated (NATO) phonetic
var letterMap = map[string]string{
    "a": "Alpha",
    "b": "Bravo",
    "c": "Charlie",
    "d": "Delta",
    "e": "Echo",
    "f": "Foxtrot",
    "g": "Golf",
    "h": "Hotel",
    "i": "India",
    "j": "Juliet",
    "k": "Kilo",
    "l": "Lima",
    "m": "Mike",
    "n": "November",
    "o": "Oscar",
    "p": "Papa",
    "q": "Quebec",
    "r": "Romeo",
    "s": "Sierra",
    "t": "Tango",
    "u": "Uniform",
    "v": "Victor",
    "w": "Whiskey",
    "x": "X-Ray",
    "y": "Yankee",
    "z": "Zulu",
    "0": "Zero",
    "1": "One",
    "2": "Two",
    "3": "Three",
    "4": "Four",
    "5": "Five",
    "6": "Six",
    "7": "Seven",
    "8": "Eight",
    "9": "Nine",
}

func translateWord(s string) string {
    var translatedword string
    s = strings.ToLower(s)
    for _, r := range s {
      // Concatenate the mapped letter to the translated word
      translatedword += letterMap[string(r)] + " "
    }
    return translatedword
}

func main() {
    // Capture the arguments in an array of strings
    var words []string = os.Args[1:]
    for _,w := range words {
        // Translate the word into NATO Phonetic
        var s string = translateWord(w)
        // Print the provided word and the translated one
        fmt.Printf("\"%s\"\n\n\t%s\n\n", w, s)
    }
}
