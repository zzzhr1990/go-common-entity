package userfile

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// FormattedPath formated
type FormattedPath struct {
	Path     string
	Identity string
}

// GetFileName get file name
func (p *FormattedPath) GetFileName() string {
	path := p.Path
	sd := strings.Split(path, "/")
	le := len(sd)
	if le < 1 {
		return ""
	}
	return sd[le-1]
}

// GetPath get file name
func (p *FormattedPath) GetPath() string {
	return p.Path
}

// GetIdentity get file name
func (p *FormattedPath) GetIdentity() string {
	return p.Identity
}

// NewFormattedPath new instance
func NewFormattedPath(path string) *FormattedPath {
	cleanStr := strings.Replace(strings.Trim(path, " \r\n\t"), "\\", "/", -1)

	// split every
	s := strings.Split(cleanStr, "/")
	var sb strings.Builder

	for _, ss := range s {
		// sb.WriteString("/")
		clear := strings.Trim(ss, " \r\n\t")
		if clear != ".." && clear != "." && len(clear) > 0 {
			sb.WriteString("/")
			sb.WriteString(clear)
		}
	}

	cleanStr = sb.String()

	/*
		if strings.HasSuffix(cleanStr, "/") {
			cleanStr = cleanStr[:len(cleanStr)-1]
		}
	*/

	if !strings.HasPrefix(cleanStr, "/") {
		cleanStr = "/" + cleanStr
	}

	h := md5.New()
	h.Write([]byte(cleanStr))
	identity := hex.EncodeToString(h.Sum(nil))
	return &FormattedPath{
		Identity: identity,
		Path:     cleanStr,
	}
}
