package ierrors

import "errors"

// NoSuchValue used in Storage error if no url there
var NoSuchValue = errors.New("no such url")

// UnknownStorageType used in Main args parser
var UnknownStorageType = errors.New("unknown storage type")

var SourceAlreadyExist = errors.New("source already exist")

var InvalidURL = errors.New("invalid url")
