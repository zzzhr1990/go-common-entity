package userfile

const (
	// OptionNotLockUserTable lock table
	OptionNotLockUserTable = 1
	// OptionOverrideFile override a file when exists
	OptionOverrideFile = 2
	// OptionInternalForceDirectory INTERNAL USE, determin if need internal call
	OptionInternalForceDirectory = 4
)

// NeedLockTable if the file need
func NeedLockTable(op int32) bool {
	return op&OptionNotLockUserTable != OptionNotLockUserTable
}

// NeedOverrideFile determin the file ov
func NeedOverrideFile(op int32) bool {
	return op&OptionOverrideFile == OptionOverrideFile
}
