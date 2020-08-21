package file

// github.com/zzzhr1990/go-common-entity/tasks/file

const (
	//FileMimeDetect detect file mime
	FileMimeDetect int32 = 1001

	//FileNotifyUri detect file mime
	//FileNotifyUri int32 = 1002

	//FileDupDelete detect file mime
	FileDupDelete int32 = 1005

	// FilePornDetected detect pron
	FilePornDetected int32 = 1010

	// UserFileAdd detect pron
	UserFileAdd int32 = 1100

	// UserFileDelete detect pron
	UserFileDelete int32 = 1110
	// UserFileReCount detect pron
	UserFileReCount int32 = 1150
	//DetectIllegal illegal
	DetectIllegal int32 = 1210

	// TrashFile trash file to dust
	TrashFile int32 = 1310
	// RecoverFile recover file from trash
	RecoverFile int32 = 1320

	// RenameFile recover file from trash
	RenameFile int32 = 1330

	// MoveFile recover file from trash
	MoveFile int32 = 1340

	// CopyFile copy file to dest
	CopyFile int32 = 1350

	// MoveToCold move it to cold
	MoveToCold int32 = 1410
	// ClearHot clear hot
	ClearHot int32 = 1420
)
