package userfile

const (
	// OptionNotLockUserTable lock table
	OptionNotLockUserTable = 1
	// OptionOverrideFile override a file when exists
	OptionOverrideFile = 2
	// OptionInternalForceDirectory INTERNAL USE, determin if need internal call
	OptionInternalForceDirectory = 4

	// OptionForceReadMaster INTERNAL USE, determin if need internal call
	OptionForceReadMaster = 8

	// OptionDisplayDeleteFile INTERNAL USE, determin if need internal call
	OptionDisplayDeleteFile = 16
)

// NeedLockTable if the file need
func NeedLockTable(op int32) bool {
	return op&OptionNotLockUserTable != OptionNotLockUserTable
}

// NeedOverrideFile determin the file ov
func NeedOverrideFile(op int32) bool {
	return op&OptionOverrideFile == OptionOverrideFile
}

// NeedForceOperateDirectory determin the file is a directory, continue process
func NeedForceOperateDirectory(op int32) bool {
	return op&OptionInternalForceDirectory == OptionInternalForceDirectory
}

// NeedForceForceReadMaster force read master
func NeedForceForceReadMaster(op int32) bool {
	return op&OptionForceReadMaster == OptionForceReadMaster
}

// NeedDisplayDeleteFile Delete File
func NeedDisplayDeleteFile(op int32) bool {
	return op&OptionDisplayDeleteFile == OptionDisplayDeleteFile
}
