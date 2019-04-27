package offline

// TaskData detail for store
// github.com/zzzhr1990/go-common-entity
type TaskData struct {
	URL      string `json:"url,omitempty"`
	Identity string `json:"identity,omitempty"`
	File     string `json:"file,omitempty"`
	User     int64  `json:"user"`
	Path     string `json:"path,omitempty"`
}
