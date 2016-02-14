package characteristic

type ObstructionDetected struct {
	*Bool
}

func NewObstructionDetected(value bool) *ObstructionDetected {
	b := NewBool(value, PermsReadOnly())
	b.Type = TypeObstructionDetected

	return &ObstructionDetected{b}
}
