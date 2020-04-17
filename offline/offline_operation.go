package offline

const (
	// OptionForceAddNew add new force
	OptionForceAddNew = 1
)

// NeedForceAddNew if the task exists, add it anyway
func NeedForceAddNew(op int32) bool {
	return op&OptionForceAddNew == OptionForceAddNew
}
