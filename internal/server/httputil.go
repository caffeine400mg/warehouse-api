package server

import (
	"errors"
	"strings"
)

var ErrInvalidIndex = errors.New("INVALID INDEX")
var ErrInvalidPath = errors.New("INVALID PATH")
var ErrNotFound = errors.New("NOT FOUND")
var ErrMethodNotAllowed = errors.New("METHOD NOT ALLOWED")

func PathIndexValue(path string, index int) (string, error) {
	if index < 0 {
		return "", ErrInvalidIndex
	}
	stringArr := strings.Split(strings.Trim(path, "/"), "/")
	if len(stringArr) <= index {
		return "", ErrInvalidPath
	}
	return stringArr[index], nil
}
