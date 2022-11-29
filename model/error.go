package model

import "errors"

var ErrCompanyAlreadyExists = errors.New("company with given id already exists")
var ErrCompanyNotFound = errors.New("company with given id not found")
