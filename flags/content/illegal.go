package content

// github.com/zzzhr1990/go-common-entity/flags/content

const (
	// FlagPornContent      BIN 001
	FlagPornContent int32 = 1
	// FlagPoliticalContent BIN 010
	FlagPoliticalContent int32 = 2
	// FlagTerrorContent    BIN 100
	FlagTerrorContent int32 = 4

	// CheckFlag            BIN 111
	CheckFlag int32 = 8
)

// IsPorn porn
func IsPorn(flag int32) bool {
	return flag&FlagPornContent > 0
}

// IsPolitical political
func IsPolitical(flag int32) bool {
	return flag&FlagPoliticalContent > 0
}

// IsFlagTerror Terror
func IsFlagTerror(flag int32) bool {
	return flag&FlagTerrorContent > 0
}

// IsIllegal Terror
func IsIllegal(flag int32) bool {
	return flag&CheckFlag > 0
}
