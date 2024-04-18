package model

type RequestPriority int64

const (
	RequestPriorityUnknown RequestPriority = iota
	RequestPriorityDefault
	RequestPriorityLow
	RequestPriorityHigh
)
