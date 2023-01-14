package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("not found")

func (d Dictionary) Search(query string) (string, error) {
	definition, ok := d[query]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(entry string, definition string) {
	d[entry] = definition
}
