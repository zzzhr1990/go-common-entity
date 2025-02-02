package model

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type CommomListToken struct {
	// Identity string `json:"identity" bson:"_id,omitempty"`
	UserIdentity string         `json:"user_identity" bson:"user_identity"`
	Skip         int64          `json:"skip" bson:"skip"`
	Limit        int64          `json:"limit" bson:"limit"`
	CreateTime   time.Time      `json:"create_time" bson:"create_time"`
	OrderBy      []*OrderByInfo `json:"order_by" bson:"order_by"`
}

type OrderByInfo struct {
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Asc   bool   `protobuf:"varint,2,opt,name=asc,proto3" json:"asc,omitempty"`
}

func (t *CommomListToken) ToTokenString() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

func CommomListTokenFromTokenString(token string) (*CommomListToken, error) {
	if len(token) == 0 {
		return &CommomListToken{
			OrderBy: []*OrderByInfo{},
		}, nil
	}
	decodedToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, ErrInvalidToken
	}
	data := &CommomListToken{}
	err = json.Unmarshal(decodedToken, data)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if data.Skip < 0 {
		data.Skip = 0
	}
	if data.Limit < 0 {
		data.Limit = 0
	}
	return data, nil
}
