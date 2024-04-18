package model

type RequestStatus int64

const (
	RequestStatusUnknown RequestStatus = iota
	RequestStatusOpen
	RequestStatusClosed
	RequestStatusDeclined
)
