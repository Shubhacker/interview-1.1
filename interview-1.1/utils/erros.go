package utils

import (
	"log"
	"net/http"
)

type ErrorsOut struct {
	ErrorMsg  string
	ErrorCode int
}

type ErrorIntup struct {
	FailFunction string
	ErrorMsg string
}

type errorCode interface {
	Error() ErrorsOut
}

func (e ErrorIntup) Error() ErrorsOut {
	var msg string
	var errors ErrorsOut
	log.Println("Test interface >", e.ErrorMsg, " < - > ", e.FailFunction)
	errors.ErrorCode = http.StatusAccepted
	errors.ErrorMsg = msg
	return errors
}