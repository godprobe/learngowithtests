package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(query string) (string, error) {
	definition, ok := d[query]
	if !ok {
		return "", errors.New("Not found")
	}
	return definition, nil
}
