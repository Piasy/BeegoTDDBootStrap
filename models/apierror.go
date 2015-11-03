package models

type ApiError struct {
	Code    int
	Message string
	Request string
}
