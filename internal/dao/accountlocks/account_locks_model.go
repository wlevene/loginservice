package accountlocks

import (
	"context"
	"errors"

	"github.com/wlevene/mgm/v3"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"

	"go.mongodb.org/mongo-driver/bson"
)

var _ AccountLocksModel = (*customAccountLocksModel)(nil)

type (
	// AccountLocksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountLocksModel.
	AccountLocksModel interface {
		accountLocksModel
		DropTable()

		FindByUserId(user_id string) (*AccountLocks, error)
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

func (model *customAccountLocksModel) FindByUserId(user_id string) (*AccountLocks, error) {

	if user_id == "" {
		return nil, errors.New("user_id is nil")
	}

	filter := bson.M{
		"user_id": user_id,
	}

	result := &AccountLocks{}

	err := mgm.Coll(result).First(filter, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
