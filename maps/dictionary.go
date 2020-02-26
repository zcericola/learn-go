package main

//Dictionary is a map with string values
type Dictionary map[string]string

//DictionaryErr is the type for an error
type DictionaryErr string

//Error Messages
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it doesn't exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

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
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

//Update will update an existing word in the dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

//Delete will remove a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func main() {

}
