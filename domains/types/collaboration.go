package types

type StatusCollaboration string

const (
	Accepted StatusCollaboration = "accepted"
	Rejected StatusCollaboration = "rejected"
	Waiting  StatusCollaboration = "waiting"
	Running  StatusCollaboration = "running"
)