package domain

import "errors"

var ErrorNotFound = errors.New("resource not found")
var ErrorInvalidPublisher = errors.New("Invalid publisher: does not exist")
var ErrorInvalidDeveloper = errors.New("Invalid developer: does not exist")
