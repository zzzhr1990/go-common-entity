package common

//  github.com/zzzhr1990/go-server-common/mq/common

import (
	"encoding/json"
)

// Entity for MQ
type Entity struct {
	Type     int32
	Identity string
	Data     json.RawMessage
}
