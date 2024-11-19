package loginattempts

import (
	"github.com/wlevene/mgm/v3"
)

type LoginAttempts struct {
	mgm.DefaultModel `bson:",inline"`

	UseId       string `bson:"user_id"`
	AttemptTime int64  `bson:"attempt_time"`
	Success     bool   `bson:"success"`
	IPAddress   string `bson:"ip_address"`
}
