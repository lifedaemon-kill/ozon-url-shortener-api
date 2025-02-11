package ierrors

import "errors"

// NoSuchValue used in Storage error if no url there
var NoSuchValue = errors.New("no such url")

// UnknownStorageType used in Main args parser
var UnknownStorageType = errors.New("unknown storage type")

var ValueAlreadyExist = errors.New("value already exist")
