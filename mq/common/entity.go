package common

//  github.com/zzzhr1990/go-common-entity/mq/common

import (
	"encoding/json"
)

// Entity for MQ
type Entity struct {
	Type     int32
	Identity string
	Data     json.RawMessage
}
