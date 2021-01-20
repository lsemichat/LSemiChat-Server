package constants

type SearchThreadTarget string

const (
	TargetUser       SearchThreadTarget = "user"
	TargetThreadName SearchThreadTarget = "thread"
	TargetTag        SearchThreadTarget = "tag"
)
