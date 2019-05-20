package tasks

// github.com/zzzhr1990/go-common-entity/tasks

const (
	//FileMimeDetect detect file mime
	FileMimeDetect int32 = 1001

	//FileDupDelete detect file mime
	FileDupDelete int32 = 1005
	// OfflineDownloadDetect detect
	// Detect Unkown Download Task
	OfflineDownloadDetect int32 = 5001

	// OfflineDownloadURL URL Only
	// URL type, http/https prifix
	OfflineDownloadURL int32 = 5020

	// OfflineDownloadTorrent Torrent Only
	// URL type, http/https prifix
	OfflineDownloadTorrent int32 = 5030

	// OfflineDownloadThunder thunder
	OfflineDownloadThunder int32 = 5040

	// OfflineDownloadMagnet magnet
	OfflineDownloadMagnet int32 = 5050
	// OfflineDownloadEd2k ED2K
	OfflineDownloadEd2k int32 = 5060
	// OfflineDownloadBaidu baidu offline
	OfflineDownloadBaidu int32 = 5070
	// OfflineDownloadBilibili bili
	OfflineDownloadBilibili int32 = 5080
	// OfflineCopyDirect if file completed, just copy user file
	OfflineCopyDirect int32 = 6010

	///////////////////////////

	// VideoPreviewGenerate generate video
	VideoPreviewGenerate int32 = 7010
	// AudioPreviewGenerate generate audio
	AudioPreviewGenerate int32 = 7020
	//////////////////////////
)