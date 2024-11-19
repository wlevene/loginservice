package dao

import (
	"log"
	"time"

	"github.com/wlevene/loginservice/internal/dao/accountlocks"
	"github.com/wlevene/loginservice/internal/dao/loginattempts"
	"github.com/wlevene/loginservice/internal/dao/user"
	"github.com/wlevene/mgm/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dao struct {
	UserModel          user.UserModel
	LoginAttemptsModel loginattempts.LoginAttemptsModel
	AccountLocksModel  accountlocks.AccountLocksModel
}

func NewDao(mongo_address string, database_name string) *Dao {

	if mongo_address == "" {
		log.Fatal("### init insou DB mongo_address")
		return nil
	}

	if database_name == "" {
		log.Fatal("### init insou DB database_name is nil")
		return nil
	}

	err := mgm.SetDefaultConfig(
		&mgm.Config{CtxTimeout: 10 * time.Second},
		database_name,
		options.Client().ApplyURI(mongo_address))

	if err != nil {
		log.Fatal("### init grapes DB has error:", err)
	}

	logx.Info("init db success :", mongo_address)

	return &Dao{

		UserModel:          user.NewUserModelV2(),
		LoginAttemptsModel: loginattempts.NewLoginAttemptsModelV2(),
		AccountLocksModel:  accountlocks.NewAccountLocksModelV2(),
	}
}
