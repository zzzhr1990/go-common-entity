package offline

// TaskData detail for store
// github.com/zzzhr1990/go-common-entity/offline
type TaskData struct {
	URL      string `json:"url,omitempty"`
	Identity string `json:"identity,omitempty"`
	File     string `json:"file,omitempty"`
	User     int64  `json:"user"`
	Path     string `json:"path,omitempty"`
}

// ServerData data for server
type ServerData struct {
	URL           string `json:"url,omitempty"`
	ShareIdentity string `json:"shareIdentity,omitempty"`
	OwnerIdentity string `json:"ownerIdentity,omitempty"`
	CreateTime    int64  `json:"createTime,omitempty"`
	Path          string `json:"path,omitempty"`
	Charset       string `json:"charset,omitempty"`
	Password      string `json:"password,omitempty"`
}
