package filetype

// github.com/zzzhr1990/go-common-entity/types/filetype

const (
	// UnknownType unknown
	UnknownType int32 = 0
	// DirectoryType for UNIX directory
	DirectoryType int32 = 10
	// ImageType for normal images
	ImageType int32 = 20
	// VideoType for video
	VideoType int32 = 30
	// DocumentType word, excel,powerpoint...
	DocumentType int32 = 40
	// AudioType for audio files
	AudioType int32 = 50

	// ArchiveType file
	ArchiveType int32 = 60

	// TorrentType BT
	TorrentType int32 = 70
)
