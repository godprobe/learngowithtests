package main

type Dictionary map[string]string

func (d Dictionary) Search(query string) string {
	return d[query]
}
