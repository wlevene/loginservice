package accountlocks

import (
	"github.com/wlevene/mgm/v3"
)

type AccountLocks struct {
	mgm.DefaultModel `bson:",inline"`

	UseId      string `bson:"user_id"`
	LockTime   int64  `bson:"lock_time"`
	UnlockTime int64  `bson:"unlock_time"`
	LockReason string `bson:"lock_reason"`
}
