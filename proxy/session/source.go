package session

type Source interface {
	Type() SourceType

	TranslateEid(now uint64) uint64
	TranslateUid(now int64) int64

	OldRuntimeId() uint64
	NewRuntimeId() uint64

	OldUniqueId() int64
	NewUniqueId() int64
}

type SourceType int

const (
	SourceTypePlayer = iota
	SourceTypeSession
)
