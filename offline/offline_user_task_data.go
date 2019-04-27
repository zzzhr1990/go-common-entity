package offline

// TaskData detail for store
// github.com/zzzhr1990/go-common-entity
type TaskData struct {
	Hash string `json:"hash,omitempty"`
	File string `json:"file,omitempty"`
	User int64  `json:"user"`
	Path string `json:"path,omitempty"`
}
