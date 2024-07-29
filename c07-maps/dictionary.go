package c07

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you're looking for")
	}
	return definition, nil
}

// map is not a reference, it is a pointer to a runtime.hmap structure
// by using below, the pointer is copied but the mutation could happen
func (d Dictionary) Add(word string, definition string) error {
	if _, ok := d[word]; ok {
		return errors.New("cannot add word because it already exists")
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	if _, ok := d[word]; !ok {
		return errors.New("cannot update word because it already exists")
	}
	d[word] = definition
	return nil
}
