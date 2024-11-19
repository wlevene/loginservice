package loginattempts

import (
	"context"
	"errors"
	"time"

	"github.com/wlevene/mgm/v3"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	DefaultTimeWindowsSize = 10 * time.Minute
)

var _ LoginAttemptsModel = (*customLoginAttemptsModel)(nil)

type (
	// LoginAttemptsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoginAttemptsModel.
	LoginAttemptsModel interface {
		loginAttemptsModel
		DropTable()

		FindFailAttemptsByTimeWindowSize(time_window_size time.Duration, user_id string) ([]*LoginAttempts, error)
		FindByTimeWindowSize(time_window_size time.Duration, user_id string) ([]*LoginAttempts, error)
		FindByDefaultTimeWindowSize(user_id string) ([]*LoginAttempts, error)
		FindByUserId(user_id string) (*LoginAttempts, error)
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

func (*customLoginAttemptsModel) FindByUserId(user_id string) (*LoginAttempts, error) {

	if user_id == "" {
		return nil, errors.New("user_id is nil")
	}

	filter := bson.M{
		"user_id": user_id,
	}

	result := &LoginAttempts{}

	err := mgm.Coll(result).First(filter, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (*customLoginAttemptsModel) FindByTimeWindowSize(time_window_size time.Duration,
	user_id string) ([]*LoginAttempts, error) {

	if user_id == "" {
		return nil, errors.New("user_id is nil")
	}

	// Calculate the start time of the window based on the current time and the duration
	startTime := time.Now().Add(-time_window_size)

	filter := bson.M{
		"user_id": user_id,
		"attempt_time": bson.M{
			"$gte": startTime, // Greater than or equal to start time
		},
	}

	var results []*LoginAttempts

	// Assuming 'mgm.Coll(&LoginAttempts{})' properly initializes your collection
	err := mgm.Coll(&LoginAttempts{}).SimpleFind(&results, filter)
	if err != nil {
		return nil, err
	}

	return results, nil

}

func (model *customLoginAttemptsModel) FindByDefaultTimeWindowSize(
	user_id string) ([]*LoginAttempts, error) {

	return model.FindByTimeWindowSize(DefaultTimeWindowsSize, user_id)

}

func (*customLoginAttemptsModel) FindFailAttemptsByTimeWindowSize(time_window_size time.Duration, user_id string) ([]*LoginAttempts, error) {

	if user_id == "" {
		return nil, errors.New("user_id is nil")
	}

	// Calculate the start time of the window based on the current time and the duration
	startTime := time.Now().Add(-time_window_size)

	filter := bson.M{
		"user_id": user_id,
		"attempt_time": bson.M{
			"$gte": startTime, // Greater than or equal to start time
		},
		"success": false,
	}

	var results []*LoginAttempts

	// Assuming 'mgm.Coll(&LoginAttempts{})' properly initializes your collection
	err := mgm.Coll(&LoginAttempts{}).SimpleFind(&results, filter)
	if err != nil {
		return nil, err
	}

	return results, nil

}
