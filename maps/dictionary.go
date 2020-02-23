package main

import "errors"

//Dictionary is a map with string values
type Dictionary map[string]string

//ErrNotFound is the error msg for unknown words
var ErrNotFound = errors.New("could not find the word you were looking for")

//Search is a function that iterates over a map to find a specific key value pair
//it is a method of Dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

//Add will add a word to the dictionary
func (d Dictionary) Add(word, definition string) {
	d[word] = definition

}

func main() {

}
