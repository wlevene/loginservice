package loginattempts

import (
	"context"

	"github.com/wlevene/mgm/v3"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ LoginAttemptsModel = (*customLoginAttemptsModel)(nil)

type (
	// LoginAttemptsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoginAttemptsModel.
	LoginAttemptsModel interface {
		loginAttemptsModel
		DropTable()
	}

	customLoginAttemptsModel struct {
		*defaultLoginAttemptsModel
	}
)

// NewLoginAttemptsModel returns a model for the mongo.
func NewLoginAttemptsModel(url, db, collection string, c cache.CacheConf) LoginAttemptsModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customLoginAttemptsModel{
		defaultLoginAttemptsModel: newDefaultLoginAttemptsModel(conn),
	}
}

func NewLoginAttemptsModelV2() LoginAttemptsModel {
	return &customLoginAttemptsModel{
		defaultLoginAttemptsModel: newDefaultLoginAttemptsModel(nil),
	}
}

func (model *customLoginAttemptsModel) DropTable() {
	mgm.Coll(&LoginAttempts{}).Drop(context.Background())
}
