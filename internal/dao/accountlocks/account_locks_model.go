package accountlocks

import (
	"context"

	"github.com/wlevene/mgm/v3"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ AccountLocksModel = (*customAccountLocksModel)(nil)

type (
	// AccountLocksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountLocksModel.
	AccountLocksModel interface {
		accountLocksModel
		DropTable()
	}

	customAccountLocksModel struct {
		*defaultAccountLocksModel
	}
)

// NewAccountLocksModel returns a model for the mongo.
func NewAccountLocksModel(url, db, collection string, c cache.CacheConf) AccountLocksModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customAccountLocksModel{
		defaultAccountLocksModel: newDefaultAccountLocksModel(conn),
	}
}

func NewAccountLocksModelV2() AccountLocksModel {
	return &customAccountLocksModel{
		defaultAccountLocksModel: newDefaultAccountLocksModel(nil),
	}
}

func (model *customAccountLocksModel) DropTable() {
	mgm.Coll(&AccountLocks{}).Drop(context.Background())
}
