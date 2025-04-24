package usecase

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrInvalidTaskName = errors.New("invalid task name")
