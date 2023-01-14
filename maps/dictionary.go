package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("can't add; word already exists in dictionary")

func (d Dictionary) Search(query string) (string, error) {
	definition, ok := d[query]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(entry string, definition string) error {
	_, err := d.Search(entry)

	switch err {
	case ErrNotFound:
		d[entry] = definition
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}
