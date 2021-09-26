package service

type ErrorType int64

const (
	DatabaseError ErrorType = iota
	JsonError
	BadRequest
	NoError
)
