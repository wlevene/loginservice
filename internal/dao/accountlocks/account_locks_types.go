package accountlocks

import (
	"fmt"

	"github.com/wlevene/mgm/v3"
)

type AccountLocks struct {
	mgm.DefaultModel `bson:",inline"`

	UseId      string `bson:"user_id"`
	LockTime   int64  `bson:"lock_time"`
	UnlockTime int64  `bson:"unlock_time"`
	LockReason string `bson:"lock_reason"`
}

func (model *AccountLocks) ToString() string {
	return fmt.Sprintf("UserId:%s,LockTime:%d,UnlockTime:%d,LockReason:%s", model.UseId, model.LockTime, model.UnlockTime, model.LockReason)
}
