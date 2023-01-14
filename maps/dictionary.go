package main

type Dictionary map[string]string

const (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrDoesNotExist  = DictionaryErr("cannot update word because it is not in the dictionary")
	ErrAlreadyExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(entry, definition string) error {
	_, err := d.Search(entry)

	switch err {
	case ErrNotFound:
		return ErrDoesNotExist
	case nil:
		d[entry] = definition
	default:
		return err
	}

	return nil
}
