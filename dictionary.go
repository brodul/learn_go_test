package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("Not found")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}
